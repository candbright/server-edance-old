package mysql

import (
	"fmt"
	"github.com/candbright/edance"
	"github.com/candbright/edance/db/domain"
	"github.com/candbright/util/xlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func (db *DB) InitTables() error {
	if db.DB == nil {
		return xlog.Wrap(edance.ErrNilDB)
	}
	err := db.DB.AutoMigrate(&domain.Song{})
	if err != nil {
		return xlog.Wrap(err)
	}
	return nil
}

func NewDb() (*DB, error) {
	const (
		userName = "root"
		password = "1q2w3e@4R"
		dbName   = "edance"
		params   = "?charset-utf8&parseTime=True&loc=Local"
	)
	var (
		err error
		ssh string
	)
	if edance.DbIp == "" || edance.DbPort == "" {
		ssh = ""
	} else {
		ssh = fmt.Sprintf("tcp(%s:%s)", edance.DbIp, edance.DbPort)
	}
	dsn := fmt.Sprintf("%s:%s@%s/%s%s", userName, password, ssh, dbName, params)
	dbConn, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,  // default size for string fields
		DisableDatetimePrecision:  true, // disable datetime precision, witch not supported before MySQL 5.6
		DontSupportRenameIndex:    true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true, // `change` when rename column, rename column not supported before MuSQL 8, MariaDB
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		return nil, xlog.Wrap(edance.ErrDBOpenFailed(err))
	}
	db := &DB{dbConn}
	err = db.InitTables()
	if err != nil {
		return nil, xlog.Wrap(edance.ErrDBInitTablesFailed(err))
	}
	return &DB{dbConn}, nil
}
