package initialize

import (
	"fmt"
	"io"
	"os"
	"void-project/global"
	"void-project/internal/repository/driver"
	"void-project/pkg/logger"

	"github.com/gin-gonic/gin"
)

func init() {
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
}

// 初始化配置文件
func InitConfig() {
	global.InitConfig()
}

// 初始化数据库连接
func InitRepository() {
	driver.InitMySQL()
	driver.InitRedis()
	driver.InitSQLite()
}

// 初始化Server日志
func InitServerLog() func() {
	file, err := logger.OpenLogFile(logger.ServerLevel)
	if err != nil {
		panic(err)
	}
	// defer file.Close()
	gin.DefaultWriter = io.MultiWriter(os.Stdout, file)
	gin.ForceConsoleColor()
	return func() {
		file.Close()
	}
}
