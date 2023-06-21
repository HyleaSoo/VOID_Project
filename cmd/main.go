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
    │ Elysium, in the Blue Sky. ファンタジーアドベンチャー。 泡泡枪、七彩
    ├───────────────────────────────────────────────────────────────────────────────────────────┤
    │ 银河系 🌌⚛️🔮🗡️✡️🏞️🎮 Requests.                                                   2023
    ├───────────────────────────────────────────────────────────────────────────────────────────┤
    │                                                                       —————— Hylea Soo
    └───────────────────────────────────────────────────────────────────────────────────────────┘
    `)

	// 初始化数据库连接
	initialize.InitRepository()

	// 绑定路由
	r := gin.Default()
	router.SetApiRouter(r) // api router
	router.SetWebRouter(r) // view router (html templates)

	// 启动监听服务
	err := r.Run(":5555")
	if err != nil {
		panic(err)
	}
}
