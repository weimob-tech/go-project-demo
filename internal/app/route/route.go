package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/weimob-tech/go-project-demo/internal/pkg/controller"
)

func Route(h *server.Hertz) {
	h.GET("demo/index/get/:user", controller.DemoGetController())
	h.POST("demo/index/post", controller.DemoPostController())

	h.GET("demo/httpClient/getInvoke", controller.HttpClientGetInvokeController())
	h.GET("demo/httpClient/postInvoke", controller.HttpClientPostInvokeController())

	h.GET("demo/redis/get/:key", controller.RedisGetController())
	h.POST("demo/redis/set", controller.RedisSetController())

	h.GET("demo/mysql/insert", controller.MysqlInsertController())
	h.GET("demo/mysql/query", controller.MysqlQueryController())

	h.GET("demo/log/info", controller.LogInfoController())
	h.GET("demo/log/debug", controller.LogDebugController())
	h.GET("demo/log/warn", controller.LogWarnController())
	h.GET("demo/log/error", controller.LogErrorController())
	h.GET("demo/log/panic", controller.LogPanicController())
	h.GET("demo/log/fatal", controller.LogFatalController())

	h.GET("demo/api/wosInvoke", controller.WeimobApiWosController())
	h.GET("demo/api/xinyunInvoke", controller.WeimobApiXinyunController())

	h.GET("demo/apollo/getConfig", controller.ApolloController())

	h.GET("demo/oauth/accesstoken", controller.AccessTokenController())

	h.GET("demo/encrypt", controller.EncryptController())
}
