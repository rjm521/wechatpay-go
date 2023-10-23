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

type OrdersApiService services.Service

// CreateOrder 请求分账API
//
// 微信订单支付成功后，商户发起分账请求，将结算后的资金分到分账接收方
// 注意：
// - 对同一笔订单最多能发起20次分账请求，每次请求最多分给50个接收方
// - 此接口采用异步处理模式，即在接收到商户请求后，会先受理请求再异步处理，最终的分账结果可以通过查询分账接口获取
//
// 错误码列表
// |名称|状态码|描述示例|原因|解决方案
// |-|-|-|-|-|
// |SYSTEM_ERROR|500|系统错误|系统超时|系统异常，请使用相同参数稍后重新调用|
// |PARAM_ERROR|400|订单号格式不正确|请求参数不符合参数格式|请使用正确的参数重新调用|
// |INVALID_REQUEST|400|非分账订单不支持分账|请求参数符合参数格式，但不符合业务规则|请根据返回的错误信息确认违反的业务规则|
// |RATELIMIT_EXCEED|429|对同笔订单分账频率过高|接口有频率限制|请降低频率后重试|
// |NOT_ENOUGH|403|分账金额不足|传入的分账金额超过了当前订单的剩余可分金额|调整分账金额|
// |NO_AUTH|403|商户无权限|未开通分账权限|请开通商户号分账权限|
func (a *OrdersApiService) CreateOrder(ctx context.Context, req CreateOrderRequest) (resp *OrdersEntity, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// 对请求中敏感字段进行加密
	encReq := req.Clone()
	encryptCertificate, err := a.Client.EncryptRequest(ctx, encReq)
	if err != nil {
		return nil, nil, fmt.Errorf("encrypt request failed: %v", err)
	}

	if encryptCertificate != "" {
		localVarHeaderParams.Set(consts.WechatPaySerial, encryptCertificate)
	}
	req = *encReq

	localVarPath := consts.WechatPayAPIServer + "/v3/profitsharing/orders"
	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract OrdersEntity from Http Response
	resp = new(OrdersEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryOrder 查询分账结果API
//
// 发起分账请求后，可调用此接口查询分账结果
// 注意：
// - 发起解冻剩余资金请求后，可调用此接口查询解冻剩余资金的结果
//
// 错误码列表
// |名称|状态码|描述示例|原因|解决方案
// |-|-|-|-|-|
// |SYSTEM_ERROR|500|系统错误|系统超时|系统异常，请使用相同参数稍后重新调用|
// |PARAM_ERROR|400|商户号未设置|请求参数不符合参数格式|请使用正确的参数重新调用|
// |RATELIMIT_EXCEED|429|商户发起分账查询的频率过高|接口有频率限制|请降低频率后重试|
// |RESOURCE_NOT_EXISTS|404|记录不存在|分账单不存在|请检查请求的单号是否正确|
func (a *OrdersApiService) QueryOrder(ctx context.Context, req QueryOrderRequest) (resp *OrdersEntity, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.OutOrderNo == nil {
		return nil, nil, fmt.Errorf("field `OutOrderNo` is required and must be specified in QueryOrderRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/profitsharing/orders/{out_order_no}"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"out_order_no"+"}", neturl.PathEscape(core.ParameterToString(*req.OutOrderNo, "")), -1)

	// Make sure All Required Params are properly set
	if req.TransactionId == nil {
		return nil, nil, fmt.Errorf("field `TransactionId` is required and must be specified in QueryOrderRequest")
	}

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	if req.SubMchid != nil {
		localVarQueryParams.Add("sub_mchid", core.ParameterToString(*req.SubMchid, ""))
	}
	localVarQueryParams.Add("transaction_id", core.ParameterToString(*req.TransactionId, ""))

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract OrdersEntity from Http Response
	resp = new(OrdersEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// UnfreezeOrder 解冻剩余资金API
//
// 不需要进行分账的订单，可直接调用本接口将订单的金额全部解冻给特约商户
// 注意：
// - 调用分账接口后，需要解冻剩余资金时，调用本接口将剩余的分账金额全部解冻给特约商户
// - 此接口采用异步处理模式，即在接收到商户请求后，会先受理请求再异步处理，最终的分账结果可以通过查询分账接口获取
//
// 错误码列表
// |名称|状态码|描述示例|原因|解决方案
// |-|-|-|-|-|
// |SYSTEM_ERROR|500|系统错误|系统超时|系统异常，请使用相同参数稍后重新调用|
// |PARAM_ERROR|400|订单号格式不正确|请求参数不符合参数格式|请使用正确的参数重新调用|
// |INVALID_REQUEST|400|非分账订单不支持解冻剩余资金|请求参数符合参数格式，但不符合业务规则|请根据返回的错误信息确认违反的业务规则|
// |RATELIMIT_EXCEED|429|对同笔订单分账频率过高|接口有频率限制|请降低频率后重试|
// |NOT_ENOUGH|403|分账金额为0|分账已完成|分账已完成，无需再请求解冻剩余资金|
// |NO_AUTH|403|商户无权限|未开通分账权限|请开通商户号分账权限|
func (a *OrdersApiService) UnfreezeOrder(ctx context.Context, req UnfreezeOrderRequest) (resp *OrdersEntity, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/profitsharing/orders/unfreeze"
	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract OrdersEntity from Http Response
	resp = new(OrdersEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
