package common

type Code int

const (
	// CodeInvalidInputParameter 输入参数无效
	CodeInvalidInputParameter Code = 10002
	// CodeVerificationCodeVerificationFailed 验证码验证失败
	CodeVerificationCodeVerificationFailed Code = 10004
	// CodeNoAccess 无访问权限
	CodeNoAccess Code = 10005
	// CodeNeedLogin 需要登录
	CodeNeedLogin Code = 30002
	// CodeConditionsNotMet 未达条件
	CodeConditionsNotMet Code = 30011
	// CodeCorrespondingUserCannotBeFound 无法找到对应的用户
	CodeCorrespondingUserCannotBeFound Code = 30013
	// CodeOutOfStock 缺货
	CodeOutOfStock Code = 30023
	// CodeApikeyError 密钥认证错误或已失效
	CodeApikeyError Code = 30039
	// CodeRequiresSecondaryVerification 需要二次验证
	CodeRequiresSecondaryVerification Code = 30043
	// CodeInvalidVerification 验证无效(验证码错误)
	CodeInvalidVerification Code = 30047
	// CodePortIsAlreadyInUse 端口已被使用
	CodePortIsAlreadyInUse Code = 70020
	// CodeProductHasExpired 产品已过期
	CodeProductHasExpired Code = 70021
	// CodeTheCurrentStateOfTheProductCannotPerformThisOperation 产品当前状态无法执行此操作
	CodeTheCurrentStateOfTheProductCannotPerformThisOperation Code = 70026
	// CodeAgreementRequired 需提供协议
	CodeAgreementRequired Code = 70057
	// CodeDnsVerificationFailed DNS验证失败
	CodeDnsVerificationFailed Code = 110006
)

// 基础响应
type BaseResponse struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

// 基础操作响应
type BasicOperationResponse struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

// 通用模式切换请求
type SwitchModeRequest struct {
	Mode bool `json:"mode"`
}

// VNC连接信息
type VncConnectionInfo struct {
	Code int `json:"code"`
	Data struct {
		RequestURL  string `json:"RequestURL"`  // 空
		RedirectURL string `json:"RedirectURL"` // 空
		PVEAuth     string `json:"PVEAuth"`     // 空
		VNCProxyURL string `json:"VNCProxyURL"` // PVE代理地址(连接地址)，里面会有一个unicode码(\u0026)(&),转码后就是可以直接在浏览器上使用的VNC连接地址
	} `json:"data"`
}

const (
	TBPass   = "关注成功"
	BiliPass = "雨云爱你"
	QQPass   = "我爱雨云"
	ZZY      = "关注雨云谢谢喵"
)
