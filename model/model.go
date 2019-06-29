package model

import (
	"github.com/chenzr/gin-web/cache"
	"github.com/chenzr/gin-web/config"
	"github.com/chenzr/gin-web/db"
)

var (
	MgConn     *db.DBConnection = nil
	MgDB       string           = ""
	RedisCache *cache.Cache     = nil
)

func Load() {
	MgConn = db.NewConnection(config.GetConfig().MongoURL)
	MgDB = config.GetConfig().MongoDB
	RedisCache = cache.NewCache(config.GetConfig().RedisURL)
}
