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

	"github.com/rjm521/wechatpay-go/core"
	"github.com/rjm521/wechatpay-go/core/consts"
	"github.com/rjm521/wechatpay-go/services"
)

type BillShipmentApiService services.Service

// SplitBill 获取分账账单文件下载地址
//
// 下载接口说明：
// 微信支付按天提供分账账单文件，商户可以通过该接口获取账单文件的下载地址。文件内包含分账相关的金额、时间等信息，供商户核对到账等情况。
// 注意：
// - 微信侧未成功的分账单不会出现在对账单中。
// - 对账单中涉及金额的字段单位为“元”；
// - 对账单接口只能下载三个月以内的账单。
//
// 文件格式说明：
// - 账单文件包括明细数据和汇总数据两部分，每一部分都包含一行表头和若干行具体数据。
// - 明细数据每一行对应一笔分账或一笔回退，同时每一个数据前加入了字符`，以避免数据被Excel按科学计数法处理。如需汇总金额等数据，可以批量替换掉该字符。
//
// 错误码列表
// |名称|状态码|描述示例|原因|解决方案
// |-|-|-|-|-|
// |SYSTEM_ERROR|500|系统错误|系统超时|系统异常，请使用相同参数稍后重新调用|
// |PARAM_ERROR|400|日期格式不正确|请求参数不符合参数格式|请使用正确的参数重新调用|
// |RATELIMIT_EXCEED|429|商户发起获取账单的频率过高|接口有频率限制|请降低频率后重试|
// |INVALID_REQUEST|400|服务商和分账方无受理关系|服务商和分账方无受理关系|请检查子商户号是否正确
// |INVALID_REQUEST|400|请求的账单日期已过期|请求的账单日期已过期|请调整账单日期
// |INVALID_REQUEST|400|请求的账单正在生成中|账单尚未生成|请确认账单时间是否正确，当日的账单会在次日上午10点以后可供下载
// |NO_STATEMENT_EXIST|400|账单文件不存在|账单尚未生成|请确认账单时间是否正确，当日的账单会在次日上午10点以后可供下载
func (a *BillShipmentApiService) SplitBill(ctx context.Context, req SplitBillRequest) (resp *SplitBillResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/profitsharing/bills"
	// Make sure All Required Params are properly set
	if req.BillDate == nil {
		return nil, nil, fmt.Errorf("field `BillDate` is required and must be specified in SplitBillRequest")
	}

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	if req.SubMchid != nil {
		localVarQueryParams.Add("sub_mchid", core.ParameterToString(*req.SubMchid, ""))
	}
	localVarQueryParams.Add("bill_date", core.ParameterToString(*req.BillDate, ""))
	if req.TarType != nil {
		localVarQueryParams.Add("tar_type", core.ParameterToString(*req.TarType, ""))
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract SplitBillResponse from Http Response
	resp = new(SplitBillResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
