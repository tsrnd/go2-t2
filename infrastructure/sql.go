package infrastructure

import (
	"github.com/jinzhu/gorm"
	// blank import.
	_ "github.com/lib/pq"
	"github.com/tsrnd/trainning/shared/gorm/database"
)

const (
	// DBMaster set master database string.
	DBMaster = "master"
	// DBRead set read replica database string.
	DBRead = "read"
)

// SQL struct.
type SQL struct {
	// Master connections master database.
	Master *gorm.DB
	// Read connections read replica database.
	Read *gorm.DB
}

type dbInfo struct {
	dbms    string
	host    string
	user    string
	pass    string
	name    string
	logmode bool
}

// NewSQL returns new SQL.
func NewSQL() *SQL {
	info := map[string]dbInfo{}
	info[DBMaster] = dbInfo{
		dbms:    GetConfigString("database_master.dbms"),
		host:    GetConfigString("database_master.host"),
		user:    GetConfigString("database_master.user"),
		pass:    GetConfigString("database_master.pass"),
		name:    GetConfigString("database_master.name"),
		logmode: GetConfigBool("database_master.logmode"),
	}
	info[DBRead] = dbInfo{
		dbms:    GetConfigString("database_read.dbms"),
		host:    GetConfigString("database_read.host"),
		user:    GetConfigString("database_read.user"),
		pass:    GetConfigString("database_read.pass"),
		name:    GetConfigString("database_read.name"),
		logmode: GetConfigBool("database_read.logmode"),
	}
	var master, read *gorm.DB

	for i, v := range info {
		connect := "host=" + v.host + " user=" + v.user + " dbname=" + v.name + " sslmode=disable password=" + v.pass
		db, err := gorm.Open(v.dbms, connect)
		db.LogMode(v.logmode)
		// Disable table name's pluralization globally
		// if set this to true, `User`'s default table name will be `user`, table name setted with `TableName` won't be affected
		db.SingularTable(true)
		if err != nil {
			panic(err)
		}

		if i == DBMaster {
			master = db
		} else if i == DBRead {
			// can't create/update/delete read replica database.
			db.Callback().Create().Before("gorm:create").Register("read_create", database.CreateErrorCallback)
			db.Callback().Update().Before("gorm:update").Register("read_update", database.UpdateErrorCallback)
			db.Callback().Delete().Before("gorm:delete").Register("read_delete", database.DeleteErrorCallback)
			read = db
		}
	}
	return &SQL{master, read}
}
