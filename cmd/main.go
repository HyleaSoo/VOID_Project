package main

import (
	"fmt"
	"void-project/initialize"
	"void-project/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(`
    ┌───────────────────────────────────────────────────────────────────────────────────────────┐
    │ Sū Shēngxǜ's from past to present VOID CHAOS False Philosophy code.
    ├───────────────────────────────────────────────────────────────────────────────────────────┤
    │ Elysium, in the Blue Sky. ファンタジーアドベンチャー。 泡泡枪、七彩、环世界宇宙飞船
    ├───────────────────────────────────────────────────────────────────────────────────────────┤
    │ 银河系 🌌⚛️🔮🗡️✡️🏞️🎮 Requests.                                                   2023
    ├───────────────────────────────────────────────────────────────────────────────────────────┤
    │                                                                       —————— Hylea Soo
    └───────────────────────────────────────────────────────────────────────────────────────────┘
    `)
	// 初始化配置
	initialize.InitConfig()
	// 初始化数据库连接
	initialize.InitRepository()

	// 初始化Server日志
	logClose := initialize.InitServerLog()
	defer logClose()

	// Server模式 debug/release
	// gin.SetMode(gin.ReleaseMode)

	// Gin引擎实例
	r := gin.Default()

	// 绑定路由
	router.SetApiRouter(r) // api router
	router.SetWebRouter(r) // view router (html templates)

	// 启动监听服务
	err := r.Run(":5555")
	if err != nil {
		panic(err)
	}
}
