package context

import (
	"github.com/midy177/wechat-sdk/credential"
	"github.com/midy177/wechat-sdk/officialaccount/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
