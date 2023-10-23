// Copyright 2021 Tencent Inc. All rights reserved.
//
// 营销商家券对外API
//
// No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
//
// API version: 0.0.11

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package merchantexclusivecoupon

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

type BusiFavorApiService services.Service

// CouponCodeInfo 查询预存code详情
//
// 查询预存code详情
func (a *BusiFavorApiService) CouponCodeInfo(ctx context.Context, req CouponCodeInfoRequest) (resp *CouponCodeInfoResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.StockId == nil {
		return nil, nil, fmt.Errorf("field `StockId` is required and must be specified in CouponCodeInfoRequest")
	}
	if req.CouponCode == nil {
		return nil, nil, fmt.Errorf("field `CouponCode` is required and must be specified in CouponCodeInfoRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/busifavor/stocks/{stock_id}/couponcodes/{coupon_code}"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"stock_id"+"}", neturl.PathEscape(core.ParameterToString(*req.StockId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"coupon_code"+"}", neturl.PathEscape(core.ParameterToString(*req.CouponCode, "")), -1)

	// Make sure All Required Params are properly set

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	if req.Appid != nil {
		localVarQueryParams.Add("appid", core.ParameterToString(*req.Appid, ""))
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

	// Extract CouponCodeInfoResponse from Http Response
	resp = new(CouponCodeInfoResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// CreateBusifavorStock 创建商家券
//
// 商户可以通过该接口创建商家券。微信支付生成商家券批次后并返回商家券批次号给到商户，商户可调用发券接口[【小程序发券】](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_3_1.shtml)、[【H5发券】](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_4_1.shtml)发放该批次商家券。
// 频率限制：接口级限制1000/min
func (a *BusiFavorApiService) CreateBusifavorStock(ctx context.Context, req CreateBusifavorStockRequest) (resp *CreateBusiFavorStockResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/busifavor/stocks"
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

	// Extract CreateBusiFavorStockResponse from Http Response
	resp = new(CreateBusiFavorStockResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// DeleteCouponCode 删除预存code
//
// 删除预存code
func (a *BusiFavorApiService) DeleteCouponCode(ctx context.Context, req DeleteCouponCodeRequest) (resp *DeleteCouponCodeResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.StockId == nil {
		return nil, nil, fmt.Errorf("field `StockId` is required and must be specified in DeleteCouponCodeRequest")
	}
	if req.CouponCode == nil {
		return nil, nil, fmt.Errorf("field `CouponCode` is required and must be specified in DeleteCouponCodeRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/busifavor/stocks/{stock_id}/couponcodes/{coupon_code}"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"stock_id"+"}", neturl.PathEscape(core.ParameterToString(*req.StockId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"coupon_code"+"}", neturl.PathEscape(core.ParameterToString(*req.CouponCode, "")), -1)

	// Make sure All Required Params are properly set
	if req.DeleteRequestNo == nil {
		return nil, nil, fmt.Errorf("field `DeleteRequestNo` is required and must be specified in DeleteCouponCodeRequest")
	}

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("delete_request_no", core.ParameterToString(*req.DeleteRequestNo, ""))

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract DeleteCouponCodeResponse from Http Response
	resp = new(DeleteCouponCodeResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// ModifyBudget 修改批次预算
//
// 商户可以通过该接口修改批次单天发放上限数量或者批次最大发放数量
func (a *BusiFavorApiService) ModifyBudget(ctx context.Context, req ModifyBudgetRequest) (resp *ModifyBudgetResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.StockId == nil {
		return nil, nil, fmt.Errorf("field `StockId` is required and must be specified in ModifyBudgetRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/busifavor/stocks/{stock_id}/budget"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"stock_id"+"}", neturl.PathEscape(core.ParameterToString(*req.StockId, "")), -1)

	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = &ModifyBudgetBody{
		TargetMaxCoupons:       req.TargetMaxCoupons,
		CurrentMaxCoupons:      req.CurrentMaxCoupons,
		TargetMaxCouponsByDay:  req.TargetMaxCouponsByDay,
		CurrentMaxCouponsByDay: req.CurrentMaxCouponsByDay,
		ModifyBudgetRequestNo:  req.ModifyBudgetRequestNo,
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract ModifyBudgetResponse from Http Response
	resp = new(ModifyBudgetResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// ModifyStockInfo 修改商家券基本信息
//
// 商户可以通过该接口修改商家券基本信息
// 前置条件： 已创建商家券批次，且修改时间位于有效期结束时间前
func (a *BusiFavorApiService) ModifyStockInfo(ctx context.Context, req ModifyStockInfoRequest) (result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.StockId == nil {
		return nil, fmt.Errorf("field `StockId` is required and must be specified in ModifyStockInfoRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/busifavor/stocks/{stock_id}"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"stock_id"+"}", neturl.PathEscape(core.ParameterToString(*req.StockId, "")), -1)

	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = &ModifyStockInfoBody{
		CustomEntrance:     req.CustomEntrance,
		StockName:          req.StockName,
		Comment:            req.Comment,
		GoodsName:          req.GoodsName,
		OutRequestNo:       req.OutRequestNo,
		DisplayPatternInfo: req.DisplayPatternInfo,
		CouponUseRule:      req.CouponUseRule,
		StockSendRule:      req.StockSendRule,
		NotifyConfig:       req.NotifyConfig,
		Subsidy:            req.Subsidy,
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return result, err
	}

	return result, nil
}

// QueryCouponCodeList 查询预存code列表
//
// 查询预存code列表
func (a *BusiFavorApiService) QueryCouponCodeList(ctx context.Context, req QueryCouponCodeListRequest) (resp *CouponCodeListResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.StockId == nil {
		return nil, nil, fmt.Errorf("field `StockId` is required and must be specified in QueryCouponCodeListRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/busifavor/stocks/{stock_id}/couponcodes"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"stock_id"+"}", neturl.PathEscape(core.ParameterToString(*req.StockId, "")), -1)

	// Make sure All Required Params are properly set

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	if req.Limit != nil {
		localVarQueryParams.Add("limit", core.ParameterToString(*req.Limit, ""))
	}
	if req.Offset != nil {
		localVarQueryParams.Add("offset", core.ParameterToString(*req.Offset, ""))
	}
	if req.Appid != nil {
		localVarQueryParams.Add("appid", core.ParameterToString(*req.Appid, ""))
	}
	if req.Status != nil {
		localVarQueryParams.Add("status", core.ParameterToString(*req.Status, ""))
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

	// Extract CouponCodeListResponse from Http Response
	resp = new(CouponCodeListResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryStock 查询商家券批次详情
//
// 商户可通过该接口查询已创建的商家券批次详情信息。
func (a *BusiFavorApiService) QueryStock(ctx context.Context, req QueryStockRequest) (resp *StockGetResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.StockId == nil {
		return nil, nil, fmt.Errorf("field `StockId` is required and must be specified in QueryStockRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/busifavor/stocks/{stock_id}"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"stock_id"+"}", neturl.PathEscape(core.ParameterToString(*req.StockId, "")), -1)

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

	// Extract StockGetResponse from Http Response
	resp = new(StockGetResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// UploadCouponCode 上传预存code
//
// 商家券的code码可由微信后台随机分配，同时支持商户自定义。如商家已有自己的优惠券系统，可直接使用自定义模式。即商家预先向微信支付上传券code，当券在发放时，微信支付自动从已导入的code中随机取值（不能指定），派发给用户。
func (a *BusiFavorApiService) UploadCouponCode(ctx context.Context, req UploadCouponCodeRequest) (resp *UploadCouponCodeResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.StockId == nil {
		return nil, nil, fmt.Errorf("field `StockId` is required and must be specified in UploadCouponCodeRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/busifavor/stocks/{stock_id}/couponcodes"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"stock_id"+"}", neturl.PathEscape(core.ParameterToString(*req.StockId, "")), -1)

	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = &UploadCouponCodeBody{
		CouponCodeList:  req.CouponCodeList,
		UploadRequestNo: req.UploadRequestNo,
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract UploadCouponCodeResponse from Http Response
	resp = new(UploadCouponCodeResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
