/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type APIKeySourceType string

const (
	APIKeySourceType_HEADER     APIKeySourceType = "HEADER"
	APIKeySourceType_AUTHORIZER APIKeySourceType = "AUTHORIZER"
)

type APIKeysFormat string

const (
	APIKeysFormat_csv APIKeysFormat = "csv"
)

type AuthorizerType string

const (
	AuthorizerType_TOKEN              AuthorizerType = "TOKEN"
	AuthorizerType_REQUEST            AuthorizerType = "REQUEST"
	AuthorizerType_COGNITO_USER_POOLS AuthorizerType = "COGNITO_USER_POOLS"
)

type CacheClusterSize string

const (
	CacheClusterSize_0_5  CacheClusterSize = "0.5"
	CacheClusterSize_1_6  CacheClusterSize = "1.6"
	CacheClusterSize_6_1  CacheClusterSize = "6.1"
	CacheClusterSize_13_5 CacheClusterSize = "13.5"
	CacheClusterSize_28_4 CacheClusterSize = "28.4"
	CacheClusterSize_58_2 CacheClusterSize = "58.2"
	CacheClusterSize_118  CacheClusterSize = "118"
	CacheClusterSize_237  CacheClusterSize = "237"
)

type CacheClusterStatus string

const (
	CacheClusterStatus_CREATE_IN_PROGRESS CacheClusterStatus = "CREATE_IN_PROGRESS"
	CacheClusterStatus_AVAILABLE          CacheClusterStatus = "AVAILABLE"
	CacheClusterStatus_DELETE_IN_PROGRESS CacheClusterStatus = "DELETE_IN_PROGRESS"
	CacheClusterStatus_NOT_AVAILABLE      CacheClusterStatus = "NOT_AVAILABLE"
	CacheClusterStatus_FLUSH_IN_PROGRESS  CacheClusterStatus = "FLUSH_IN_PROGRESS"
)

type ConnectionType string

const (
	ConnectionType_INTERNET ConnectionType = "INTERNET"
	ConnectionType_VPC_LINK ConnectionType = "VPC_LINK"
)

type ContentHandlingStrategy string

const (
	ContentHandlingStrategy_CONVERT_TO_BINARY ContentHandlingStrategy = "CONVERT_TO_BINARY"
	ContentHandlingStrategy_CONVERT_TO_TEXT   ContentHandlingStrategy = "CONVERT_TO_TEXT"
)

type DocumentationPartType string

const (
	DocumentationPartType_API             DocumentationPartType = "API"
	DocumentationPartType_AUTHORIZER      DocumentationPartType = "AUTHORIZER"
	DocumentationPartType_MODEL           DocumentationPartType = "MODEL"
	DocumentationPartType_RESOURCE        DocumentationPartType = "RESOURCE"
	DocumentationPartType_METHOD          DocumentationPartType = "METHOD"
	DocumentationPartType_PATH_PARAMETER  DocumentationPartType = "PATH_PARAMETER"
	DocumentationPartType_QUERY_PARAMETER DocumentationPartType = "QUERY_PARAMETER"
	DocumentationPartType_REQUEST_HEADER  DocumentationPartType = "REQUEST_HEADER"
	DocumentationPartType_REQUEST_BODY    DocumentationPartType = "REQUEST_BODY"
	DocumentationPartType_RESPONSE        DocumentationPartType = "RESPONSE"
	DocumentationPartType_RESPONSE_HEADER DocumentationPartType = "RESPONSE_HEADER"
	DocumentationPartType_RESPONSE_BODY   DocumentationPartType = "RESPONSE_BODY"
)

type DomainNameStatus_SDK string

const (
	DomainNameStatus_SDK_AVAILABLE                      DomainNameStatus_SDK = "AVAILABLE"
	DomainNameStatus_SDK_UPDATING                       DomainNameStatus_SDK = "UPDATING"
	DomainNameStatus_SDK_PENDING                        DomainNameStatus_SDK = "PENDING"
	DomainNameStatus_SDK_PENDING_CERTIFICATE_REIMPORT   DomainNameStatus_SDK = "PENDING_CERTIFICATE_REIMPORT"
	DomainNameStatus_SDK_PENDING_OWNERSHIP_VERIFICATION DomainNameStatus_SDK = "PENDING_OWNERSHIP_VERIFICATION"
)

type EndpointType string

const (
	EndpointType_REGIONAL EndpointType = "REGIONAL"
	EndpointType_EDGE     EndpointType = "EDGE"
	EndpointType_PRIVATE  EndpointType = "PRIVATE"
)

type GatewayResponseType string

const (
	GatewayResponseType_DEFAULT_4XX                    GatewayResponseType = "DEFAULT_4XX"
	GatewayResponseType_DEFAULT_5XX                    GatewayResponseType = "DEFAULT_5XX"
	GatewayResponseType_RESOURCE_NOT_FOUND             GatewayResponseType = "RESOURCE_NOT_FOUND"
	GatewayResponseType_UNAUTHORIZED                   GatewayResponseType = "UNAUTHORIZED"
	GatewayResponseType_INVALID_API_KEY                GatewayResponseType = "INVALID_API_KEY"
	GatewayResponseType_ACCESS_DENIED                  GatewayResponseType = "ACCESS_DENIED"
	GatewayResponseType_AUTHORIZER_FAILURE             GatewayResponseType = "AUTHORIZER_FAILURE"
	GatewayResponseType_AUTHORIZER_CONFIGURATION_ERROR GatewayResponseType = "AUTHORIZER_CONFIGURATION_ERROR"
	GatewayResponseType_INVALID_SIGNATURE              GatewayResponseType = "INVALID_SIGNATURE"
	GatewayResponseType_EXPIRED_TOKEN                  GatewayResponseType = "EXPIRED_TOKEN"
	GatewayResponseType_MISSING_AUTHENTICATION_TOKEN   GatewayResponseType = "MISSING_AUTHENTICATION_TOKEN"
	GatewayResponseType_INTEGRATION_FAILURE            GatewayResponseType = "INTEGRATION_FAILURE"
	GatewayResponseType_INTEGRATION_TIMEOUT            GatewayResponseType = "INTEGRATION_TIMEOUT"
	GatewayResponseType_API_CONFIGURATION_ERROR        GatewayResponseType = "API_CONFIGURATION_ERROR"
	GatewayResponseType_UNSUPPORTED_MEDIA_TYPE         GatewayResponseType = "UNSUPPORTED_MEDIA_TYPE"
	GatewayResponseType_BAD_REQUEST_PARAMETERS         GatewayResponseType = "BAD_REQUEST_PARAMETERS"
	GatewayResponseType_BAD_REQUEST_BODY               GatewayResponseType = "BAD_REQUEST_BODY"
	GatewayResponseType_REQUEST_TOO_LARGE              GatewayResponseType = "REQUEST_TOO_LARGE"
	GatewayResponseType_THROTTLED                      GatewayResponseType = "THROTTLED"
	GatewayResponseType_QUOTA_EXCEEDED                 GatewayResponseType = "QUOTA_EXCEEDED"
	GatewayResponseType_WAF_FILTERED                   GatewayResponseType = "WAF_FILTERED"
)

type IntegrationType string

const (
	IntegrationType_HTTP       IntegrationType = "HTTP"
	IntegrationType_AWS        IntegrationType = "AWS"
	IntegrationType_MOCK       IntegrationType = "MOCK"
	IntegrationType_HTTP_PROXY IntegrationType = "HTTP_PROXY"
	IntegrationType_AWS_PROXY  IntegrationType = "AWS_PROXY"
)

type LocationStatusType string

const (
	LocationStatusType_DOCUMENTED   LocationStatusType = "DOCUMENTED"
	LocationStatusType_UNDOCUMENTED LocationStatusType = "UNDOCUMENTED"
)

type Op string

const (
	Op_add     Op = "add"
	Op_remove  Op = "remove"
	Op_replace Op = "replace"
	Op_move    Op = "move"
	Op_copy    Op = "copy"
	Op_test    Op = "test"
)

type PutMode string

const (
	PutMode_merge     PutMode = "merge"
	PutMode_overwrite PutMode = "overwrite"
)

type QuotaPeriodType string

const (
	QuotaPeriodType_DAY   QuotaPeriodType = "DAY"
	QuotaPeriodType_WEEK  QuotaPeriodType = "WEEK"
	QuotaPeriodType_MONTH QuotaPeriodType = "MONTH"
)

type SecurityPolicy string

const (
	SecurityPolicy_TLS_1_0 SecurityPolicy = "TLS_1_0"
	SecurityPolicy_TLS_1_2 SecurityPolicy = "TLS_1_2"
)

type UnauthorizedCacheControlHeaderStrategy string

const (
	UnauthorizedCacheControlHeaderStrategy_FAIL_WITH_403                   UnauthorizedCacheControlHeaderStrategy = "FAIL_WITH_403"
	UnauthorizedCacheControlHeaderStrategy_SUCCEED_WITH_RESPONSE_HEADER    UnauthorizedCacheControlHeaderStrategy = "SUCCEED_WITH_RESPONSE_HEADER"
	UnauthorizedCacheControlHeaderStrategy_SUCCEED_WITHOUT_RESPONSE_HEADER UnauthorizedCacheControlHeaderStrategy = "SUCCEED_WITHOUT_RESPONSE_HEADER"
)

type VPCLinkStatus_SDK string

const (
	VPCLinkStatus_SDK_AVAILABLE VPCLinkStatus_SDK = "AVAILABLE"
	VPCLinkStatus_SDK_PENDING   VPCLinkStatus_SDK = "PENDING"
	VPCLinkStatus_SDK_DELETING  VPCLinkStatus_SDK = "DELETING"
	VPCLinkStatus_SDK_FAILED    VPCLinkStatus_SDK = "FAILED"
)
