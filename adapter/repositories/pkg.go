package repositories

import (
	"base-go/common/config"
	"base-go/common/logger"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const CtxKeyTransaction = "ctx_key_transaction"

func NewGormdb(cnf *config.Config) (db *gorm.DB) {
	dbCnf := cnf.DBconfig
	logger.Info("Connecting to database backend %s at %s@%h:%d/%s", dbCnf.DBbackend, dbCnf.Username, dbCnf.Host, dbCnf.Port, dbCnf.DBname)
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbCnf.Username, dbCnf.Password, dbCnf.Host, dbCnf.Port, dbCnf.DBname)
	switch dbCnf.DBbackend {
	case "mysql":
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN: connString,
		}), &gorm.Config{})
		if err != nil {
			logger.Error("Fatal: %s", err.Error())
			panic(err)
		}
		if db == nil {
			err = fmt.Errorf("null DB")
			panic(err)
		}
		// setup underlying *sql.DB
		sqlDB, err := db.DB()
		if err != nil {
			logger.Error("Fatal: %s", err.Error())
			panic(err)
		}

		sqlDB.SetConnMaxIdleTime(dbCnf.ConnMaxLifetime)
		sqlDB.SetMaxOpenConns(dbCnf.MaxOpenConns)
		sqlDB.SetMaxIdleConns(dbCnf.MaxIdleConns)

		if err = sqlDB.Ping(); err != nil {
			logger.Error("Fatal: DB ping failed, error: %s", err.Error())
			panic(err)
		}

		return db
	default:
		logger.Error("Fatal: Unknown DB backend")
		panic(fmt.Errorf("unknown DB backend"))
	}
}
