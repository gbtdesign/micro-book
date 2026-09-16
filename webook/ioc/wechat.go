package ioc

import (
	"golang/webook/internal/service/oauth2/wechat"
	logger2 "golang/webook/pkg/logger"
	"os"

	"github.com/joho/godotenv"
)

func InitWechatService(l logger2.LoggerV1) wechat.Service {
	_ = godotenv.Load() // 加载 .env 文件，忽略错误（如果文件不存在）
	appId, ok := os.LookupEnv("WECHAT_APP_ID")
	if !ok {
		panic("没有找到环境变量 WECHAT_APP_ID ")
	}
	appKey, ok := os.LookupEnv("WECHAT_APP_SECRET")
	if !ok {
		panic("没有找到环境变量 WECHAT_APP_SECRET")
	}
	// 692jdHsogrsYqxaUK9fgxw
	return wechat.NewService(appId, appKey, l)
}

//func NewWechatHandlerConfig() web.WechatHandlerConfig {
//	return web.WechatHandlerConfig{
//		Secure: false,
//	}
//}
