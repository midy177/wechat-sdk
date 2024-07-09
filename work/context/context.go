package context

import (
	"github.com/midy177/wechat-sdk/credential"
	"github.com/midy177/wechat-sdk/work/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
