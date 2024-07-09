package checkin

import (
	"github.com/midy177/wechat-sdk/work/context"
)

// Client 打卡接口实例
type Client struct {
	*context.Context
}

// NewClient 初始化实例
func NewClient(ctx *context.Context) *Client {
	return &Client{
		ctx,
	}
}
