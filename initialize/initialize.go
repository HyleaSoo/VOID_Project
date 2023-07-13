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
	echo()
}

func echo() {
	fmt.Println(`
    ┌───────────────────────────────────────────────────────────────────────────────────────────┐
    │                                       void-project
    ├───────────────────────────────────────────────────────────────────────────────────────────┤
    │ HyleaSoo's void-project is a web application architecture developed in Go.
    ├───────────────────────────────────────────────────────────────────────────────────────────┤
    │ Project repository link: https://github.com/HyleaSoo/void-project
    ├───────────────────────────────────────────────────────────────────────────────────────────┤
    │ 银河系 🌌⚛️🧬🧊🔮🗡️✡️🏞️🌈🎮🪞🫧 Requests.                                       2023
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
