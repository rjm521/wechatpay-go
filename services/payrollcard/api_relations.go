// Copyright 2021 Tencent Inc. All rights reserved.
//
// 微工卡接口文档
//
// 服务商通过本API文档提供的接口，查询商户和微工卡的授权关系、生成预授权的token口令、核身预下单、核身结果的查询等。
//
// API version: 1.5.2

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package payrollcard

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

type RelationsApiService services.Service

// GetRelation 查询微工卡授权关系
//
// 查询商户和微信支付用户的微工卡授权关系
func (a *RelationsApiService) GetRelation(ctx context.Context, req GetRelationRequest) (resp *RelationEntity, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.Openid == nil {
		return nil, nil, fmt.Errorf("field `Openid` is required and must be specified in GetRelationRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/payroll-card/relations/{openid}"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"openid"+"}", neturl.PathEscape(core.ParameterToString(*req.Openid, "")), -1)

	// Make sure All Required Params are properly set
	if req.SubMchid == nil {
		return nil, nil, fmt.Errorf("field `SubMchid` is required and must be specified in GetRelationRequest")
	}

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("sub_mchid", core.ParameterToString(*req.SubMchid, ""))
	if req.Appid != nil {
		localVarQueryParams.Add("appid", core.ParameterToString(*req.Appid, ""))
	}
	if req.SubAppid != nil {
		localVarQueryParams.Add("sub_appid", core.ParameterToString(*req.SubAppid, ""))
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

	// Extract RelationEntity from Http Response
	resp = new(RelationEntity)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
