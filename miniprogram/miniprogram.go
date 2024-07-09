package miniprogram

import (
	"github.com/midy177/wechat-sdk/credential"
	"github.com/midy177/wechat-sdk/internal/openapi"
	"github.com/midy177/wechat-sdk/miniprogram/analysis"
	"github.com/midy177/wechat-sdk/miniprogram/auth"
	"github.com/midy177/wechat-sdk/miniprogram/business"
	"github.com/midy177/wechat-sdk/miniprogram/config"
	"github.com/midy177/wechat-sdk/miniprogram/content"
	"github.com/midy177/wechat-sdk/miniprogram/context"
	"github.com/midy177/wechat-sdk/miniprogram/encryptor"
	"github.com/midy177/wechat-sdk/miniprogram/message"
	"github.com/midy177/wechat-sdk/miniprogram/minidrama"
	"github.com/midy177/wechat-sdk/miniprogram/order"
	"github.com/midy177/wechat-sdk/miniprogram/privacy"
	"github.com/midy177/wechat-sdk/miniprogram/qrcode"
	"github.com/midy177/wechat-sdk/miniprogram/redpacketcover"
	"github.com/midy177/wechat-sdk/miniprogram/riskcontrol"
	"github.com/midy177/wechat-sdk/miniprogram/security"
	"github.com/midy177/wechat-sdk/miniprogram/shortlink"
	"github.com/midy177/wechat-sdk/miniprogram/subscribe"
	"github.com/midy177/wechat-sdk/miniprogram/tcb"
	"github.com/midy177/wechat-sdk/miniprogram/urllink"
	"github.com/midy177/wechat-sdk/miniprogram/urlscheme"
	"github.com/midy177/wechat-sdk/miniprogram/virtualpayment"
	"github.com/midy177/wechat-sdk/miniprogram/werun"
)

// MiniProgram 微信小程序相关 API
type MiniProgram struct {
	ctx *context.Context
}

// NewMiniProgram 实例化小程序 API
func NewMiniProgram(cfg *config.Config) *MiniProgram {
	defaultAkHandle := credential.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret, credential.CacheKeyMiniProgramPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &MiniProgram{ctx}
}

// SetAccessTokenHandle 自定义 access_token 获取方式
func (miniProgram *MiniProgram) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle) {
	miniProgram.ctx.AccessTokenHandle = accessTokenHandle
}

// GetContext get Context
func (miniProgram *MiniProgram) GetContext() *context.Context {
	return miniProgram.ctx
}

// GetEncryptor  小程序加解密
func (miniProgram *MiniProgram) GetEncryptor() *encryptor.Encryptor {
	return encryptor.NewEncryptor(miniProgram.ctx)
}

// GetAuth 登录/用户信息相关接口
func (miniProgram *MiniProgram) GetAuth() *auth.Auth {
	return auth.NewAuth(miniProgram.ctx)
}

// GetAnalysis 数据分析
func (miniProgram *MiniProgram) GetAnalysis() *analysis.Analysis {
	return analysis.NewAnalysis(miniProgram.ctx)
}

// GetBusiness 业务接口
func (miniProgram *MiniProgram) GetBusiness() *business.Business {
	return business.NewBusiness(miniProgram.ctx)
}

// GetPrivacy 小程序隐私协议相关 API
func (miniProgram *MiniProgram) GetPrivacy() *privacy.Privacy {
	return privacy.NewPrivacy(miniProgram.ctx)
}

// GetQRCode 小程序码相关 API
func (miniProgram *MiniProgram) GetQRCode() *qrcode.QRCode {
	return qrcode.NewQRCode(miniProgram.ctx)
}

// GetTcb 小程序云开发 API
func (miniProgram *MiniProgram) GetTcb() *tcb.Tcb {
	return tcb.NewTcb(miniProgram.ctx)
}

// GetSubscribe 小程序订阅消息
func (miniProgram *MiniProgram) GetSubscribe() *subscribe.Subscribe {
	return subscribe.NewSubscribe(miniProgram.ctx)
}

// GetCustomerMessage 客服消息接口
func (miniProgram *MiniProgram) GetCustomerMessage() *message.Manager {
	return message.NewCustomerMessageManager(miniProgram.ctx)
}

// GetWeRun 微信运动接口
func (miniProgram *MiniProgram) GetWeRun() *werun.WeRun {
	return werun.NewWeRun(miniProgram.ctx)
}

// GetContentSecurity 内容安全接口
func (miniProgram *MiniProgram) GetContentSecurity() *content.Content {
	return content.NewContent(miniProgram.ctx)
}

// GetURLLink 小程序 URL Link 接口
func (miniProgram *MiniProgram) GetURLLink() *urllink.URLLink {
	return urllink.NewURLLink(miniProgram.ctx)
}

// GetRiskControl 安全风控接口
func (miniProgram *MiniProgram) GetRiskControl() *riskcontrol.RiskControl {
	return riskcontrol.NewRiskControl(miniProgram.ctx)
}

// GetSecurity 内容安全接口
func (miniProgram *MiniProgram) GetSecurity() *security.Security {
	return security.NewSecurity(miniProgram.ctx)
}

// GetShortLink 小程序短链接口
func (miniProgram *MiniProgram) GetShortLink() *shortlink.ShortLink {
	return shortlink.NewShortLink(miniProgram.ctx)
}

// GetSURLScheme 小程序 URL Scheme 接口
func (miniProgram *MiniProgram) GetSURLScheme() *urlscheme.URLScheme {
	return urlscheme.NewURLScheme(miniProgram.ctx)
}

// GetOpenAPI openApi 管理接口
func (miniProgram *MiniProgram) GetOpenAPI() *openapi.OpenAPI {
	return openapi.NewOpenAPI(miniProgram.ctx)
}

// GetVirtualPayment 小程序虚拟支付
func (miniProgram *MiniProgram) GetVirtualPayment() *virtualpayment.VirtualPayment {
	return virtualpayment.NewVirtualPayment(miniProgram.ctx)
}

// GetMessageReceiver 获取消息推送接收器
func (miniProgram *MiniProgram) GetMessageReceiver() *message.PushReceiver {
	return message.NewPushReceiver(miniProgram.ctx)
}

// GetShipping 小程序发货信息管理服务
func (miniProgram *MiniProgram) GetShipping() *order.Shipping {
	return order.NewShipping(miniProgram.ctx)
}

// GetMiniDrama 小程序娱乐微短剧
func (miniProgram *MiniProgram) GetMiniDrama() *minidrama.MiniDrama {
	return minidrama.NewMiniDrama(miniProgram.ctx)
}

// GetRedPacketCover 小程序微信红包封面 API
func (miniProgram *MiniProgram) GetRedPacketCover() *redpacketcover.RedPacketCover {
	return redpacketcover.NewRedPacketCover(miniProgram.ctx)
}

// GetUpdatableMessage 小程序动态消息
func (miniProgram *MiniProgram) GetUpdatableMessage() *message.UpdatableMessage {
	return message.NewUpdatableMessage(miniProgram.ctx)
}
