package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	//init mysql driver
	"github.com/sjw/user/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRedisDB, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db  *gorm.DB
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

func NewDB(conf *conf.Data) *gorm.DB {

	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}

func NewRedisDB(conf *conf.Data) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr: conf.Redis.Addr,
		DB:   0,
	})

	return rdb
}
