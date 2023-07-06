package driver

import (
	"fmt"
	"time"
	"void-project/global"
	"void-project/pkg"
	log "void-project/pkg/logger"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	MySQL     *gorm.DB
	SQLite    *gorm.DB
	SQLServer *gorm.DB
	Redis     *redis.Client
)

// 初始化MySQL数据库连接
func InitMySQL() {
	// 读取配置文件
	op := global.Config.DB.MySQL

	var err error
	MySQL, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", op.User, op.Password, op.Host, op.Port, op.DBName)), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		// SQL语句记录日志
		Logger: logger.New(
			log.NewSQLLogger(),
			logger.Config{
				SlowThreshold:             time.Second, // 慢SQL阈值
				LogLevel:                  logger.Info, // 日志级别
				IgnoreRecordNotFoundError: true,        // 忽略‘记录未找到’错误
				Colorful:                  true,        // 彩色打印
			},
		),
	})
	if err != nil {
		panic(err)
	}
}

// 初始化SQLite连接
func InitSQLite() {
	//读取配置文件
	op := global.Config.DB.SQLite

	var err error
	SQLite, err = gorm.Open(sqlite.Open(pkg.GetRootPath()+op.Path), &gorm.Config{
		// SQL语句记录
		Logger: logger.New(
			log.NewSQLLogger(),
			logger.Config{
				SlowThreshold:             time.Second, // 慢SQL阈值
				LogLevel:                  logger.Info, // 日志级别
				IgnoreRecordNotFoundError: true,        // 忽略‘记录未找到’错误
				Colorful:                  true,        // 彩色打印
			},
		),
	})
	if err != nil {
		panic(err)
	}
}

// 初始化Redis连接
func InitRedis() {
	op := global.Config.Cache.Redis
	Redis = redis.NewClient(&op)
}
