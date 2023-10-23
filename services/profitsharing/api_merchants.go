// Copyright 2021 Tencent Inc. All rights reserved.
//
// 微信支付分账API
//
// 微信支付分账API
//
// API version: 0.0.3

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package profitsharing

import (
	"context"
	"fmt"
	nethttp "net/http"
	neturl "net/url"
	"strings"

	"github.com/rjm521/wechatpay-go/core"
	"github.com/rjm521/wechatpay-go/core/consts"
	"github.com/rjm521/wechatpay-go/services"
)

type MerchantsApiService services.Service

// QueryMerchantRatio 查询最大分账比例API
//
// 可调用此接口查询特约商户设置的允许服务商分账的最大比例
//
//
// 错误码列表
// |名称|状态码|描述示例|原因|解决方案
// |-|-|-|-|-|
// |SYSTEM_ERROR|500|系统错误|系统超时|系统异常，请使用相同参数稍后重新调用|
// |PARAM_ERROR|400|商户号未设置|请求参数不符合参数格式|请使用正确的参数重新调用|
// |RATELIMIT_EXCEED|429|商户发起查询的频率过高|接口有频率限制|请降低频率后重试|
// |NO_AUTH|403|服务商未开通分账功能|服务商未开通分账|请先开通分账功能|
// |NO_AUTH|403|非服务商下的签约子商户|服务商未与子商户签约|请先签约子商户|
// |NO_AUTH|403|子商户未开通分账功能|子商户未开通分账功能|请先开通子商户的分账功能|
func (a *MerchantsApiService) QueryMerchantRatio(ctx context.Context, req QueryMerchantRatioRequest) (resp *QueryMerchantRatioResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.SubMchid == nil {
		return nil, nil, fmt.Errorf("field `SubMchid` is required and must be specified in QueryMerchantRatioRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/profitsharing/merchant-configs/{sub_mchid}"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"sub_mchid"+"}", neturl.PathEscape(core.ParameterToString(*req.SubMchid, "")), -1)

	// Make sure All Required Params are properly set

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract QueryMerchantRatioResponse from Http Response
	resp = new(QueryMerchantRatioResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
