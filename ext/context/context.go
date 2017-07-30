package context

import (
	conf "github.com/hiromaily/go-goa/ext/configs"
	"github.com/hiromaily/golibs/db/gorm"
)

// Ctx is context object
type Ctx struct {
	Conf *conf.Config
	Db   *gorm.GR
}

func (c *Ctx) initDB() {
	dbInfo := conf.GetConf().MySQL
	gorm.New(dbInfo.Host, dbInfo.DbName, dbInfo.User, dbInfo.Pass, dbInfo.Port)
	c.Db = gorm.GetDB()
}

// SetupContext is to setup context
func SetupContext(c *conf.Config) *Ctx {
	ctx := &Ctx{Conf: c}
	ctx.initDB()

	return ctx
}
