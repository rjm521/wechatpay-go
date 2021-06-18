package cert

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"sync"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/signers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/validators"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// isSameCertificateMap Check if two CertificateMaps stores same certificates.
// Normally, checking serial number set is enough.
func isSameCertificateMap(l, r map[string]*x509.Certificate) bool {
	if l == nil && r == nil {
		return true
	}

	if len(l) != len(r) {
		return false
	}

	for serialNo := range l {
		if _, ok := r[serialNo]; !ok {
			return false
		}
	}

	return true
}

// CertificateDownloader 平台证书下载器，下载完成后可直接获取 x509.Certificate 对象或导出证书内容
type CertificateDownloader struct {
	certContents map[string]string   // 证书文本内容，用于导出
	certificates core.CertificateMap // 证书实例
	client       *core.Client        // 微信支付 API v3 Go SDK HTTPClient
	mchAPIv3Key  string              // 商户APIv3密钥
	lock         sync.RWMutex
}

// Get 获取证书序列号对应的平台证书
func (o *CertificateDownloader) Get(ctx context.Context, serialNo string) (*x509.Certificate, bool) {
	o.lock.RLock()
	defer o.lock.RUnlock()

	return o.certificates.Get(ctx, serialNo)
}

// GetAll 获取平台证书Map
func (o *CertificateDownloader) GetAll(ctx context.Context) map[string]*x509.Certificate {
	o.lock.RLock()
	defer o.lock.RUnlock()

	return o.certificates.GetAll(ctx)
}

// GetNewestSerial 获取最新的平台证书的证书序列号
func (o *CertificateDownloader) GetNewestSerial(ctx context.Context) string {
	o.lock.RLock()
	defer o.lock.RUnlock()

	return o.certificates.GetNewestSerial(ctx)
}

// Export 获取证书序列号对应的平台证书内容
func (o *CertificateDownloader) Export(_ context.Context, serialNo string) (string, bool) {
	o.lock.RLock()
	defer o.lock.RUnlock()

	content, ok := o.certContents[serialNo]
	return content, ok
}

// ExportAll 获取平台证书内容Map
func (o *CertificateDownloader) ExportAll(_ context.Context) map[string]string {
	o.lock.RLock()
	defer o.lock.RUnlock()

	ret := make(map[string]string)
	for serialNo, content := range o.certContents {
		ret[serialNo] = content
	}

	return ret
}

func (o *CertificateDownloader) decryptCertificate(
	_ context.Context, encryptCertificate *encryptCertificate,
) (string, error) {
	plaintext, err := utils.DecryptAES256GCM(
		o.mchAPIv3Key, *encryptCertificate.AssociatedData,
		*encryptCertificate.Nonce, *encryptCertificate.Ciphertext,
	)
	if err != nil {
		return "", fmt.Errorf("decrypt downloaded certificate failed: %v", err)
	}

	return plaintext, nil
}

func (o *CertificateDownloader) updateCertificates(
	ctx context.Context, certContents map[string]string, certificates map[string]*x509.Certificate,
) {
	o.lock.Lock()
	defer o.lock.Unlock()
	if isSameCertificateMap(o.certificates.GetAll(ctx), certificates) {
		return
	}

	o.certContents = certContents
	o.certificates.Reset(certificates)
	o.client = core.NewClientWithValidator(
		o.client,
		&validators.WechatPayValidator{
			Verifier: verifiers.NewSHA256WithRSAVerifier(o),
		},
	)
}

func (o *CertificateDownloader) performDownloading(ctx context.Context) (*downloadCertificatesResponse, error) {
	result, err := o.client.Get(ctx, consts.WechatPayAPIServer+"/v3/certificates")
	if err != nil {
		return nil, err
	}

	resp := new(downloadCertificatesResponse)
	if err = core.UnMarshalResponse(result.Response, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// DownloadCertificates 立即下载平台证书列表
func (o *CertificateDownloader) DownloadCertificates(ctx context.Context) error {
	resp, err := o.performDownloading(ctx)
	if err != nil {
		return err
	}

	rawCertContentMap := make(map[string]string)
	certificateMap := make(map[string]*x509.Certificate)
	for _, rawCertificate := range resp.Data {
		certContent, err := o.decryptCertificate(ctx, rawCertificate.EncryptCertificate)
		if err != nil {
			return err
		}

		certificate, err := utils.LoadCertificate(certContent)
		if err != nil {
			return fmt.Errorf("parse downlaoded certificate failed: %v, certcontent:%v", err, certContent)
		}

		serialNo := utils.GetCertificateSerialNumber(*certificate)

		rawCertContentMap[serialNo] = certContent
		certificateMap[serialNo] = certificate
	}

	if len(certificateMap) == 0 {
		return fmt.Errorf("no certificate downloaded")
	}

	o.updateCertificates(ctx, rawCertContentMap, certificateMap)
	return nil
}

// NewCertificateDownloader 使用商户号/商户私钥等信息初始化商户的平台证书下载器 CertificateDownloader
// 初始化完成后会立即发起一次下载，确保下载器被正确初始化。
func NewCertificateDownloader(
	ctx context.Context, mchID string, privateKey *rsa.PrivateKey, certificateSerialNo string, mchAPIv3Key string,
) (*CertificateDownloader, error) {
	settings := core.DialSettings{
		Signer: &signers.SHA256WithRSASigner{
			MchID:               mchID,
			PrivateKey:          privateKey,
			CertificateSerialNo: certificateSerialNo,
		},
		Validator: &validators.NullValidator{},
	}

	client, err := core.NewClientWithDialSettings(ctx, &settings)
	if err != nil {
		return nil, fmt.Errorf("create downloader failed, create client err:%v", err)
	}

	return NewCertificateDownloaderWithClient(ctx, client, mchAPIv3Key)
}

// NewCertificateDownloaderWithClient 使用 core.Client 初始化商户的平台证书下载器 CertificateDownloader
// 初始化完成后会立即发起一次下载，确保下载器被正确初始化。
func NewCertificateDownloaderWithClient(
	ctx context.Context, client *core.Client, mchAPIv3Key string,
) (*CertificateDownloader, error) {
	downloader := CertificateDownloader{
		client:      client,
		mchAPIv3Key: mchAPIv3Key,
	}

	if err := downloader.DownloadCertificates(ctx); err != nil {
		return nil, err
	}

	return &downloader, nil
}
