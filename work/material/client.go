package material

import (
	"github.com/midy177/wechat-sdk/work/context"
)

// Client 素材管理接口实例
type Client struct {
	*context.Context
}

// NewClient 初始化实例
func NewClient(ctx *context.Context) *Client {
	return &Client{
		ctx,
	}
}
