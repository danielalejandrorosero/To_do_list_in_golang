package conn

import (
	"github.com/danielalejandrorosero/to_do_list/app/common"
	"github.com/danielalejandrorosero/to_do_list/app/utils/consts"
	"github.com/danielalejandrorosero/to_do_list/infra/logger"
	"gopkg.in/mgo.v2"
)

func ConnectDb() *mgo.Database {
	sess, err := mgo.Dial(consts.HostName)
	common.CheckErr(err)
	sess.SetMode(mgo.Monotonic, true)
	db := sess.DB(consts.DbName)
	logger.Info("Databases connection successful...")
	return db

}
