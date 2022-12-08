package db

import (
	"edance"
	"edance/db/domain"
	"edance/db/mysql"
)

var db DB

type DB interface {
	InitTables() error
	ListAllSong() ([]domain.Song, error)
	ListSongByMode(mode string, modeDetail string) ([]domain.Song, error)
	GetSongById(songId string) (domain.Song, error)
	AddSong(song domain.Song) error
	DeleteSong(songId string) error
}

func GetDb() (DB, error) {
	if db == nil {
		switch edance.DbType {
		case "mysql":
			mysqlDb, err := mysql.NewDb()
			if err != nil {
				return nil, err
			}
			db = mysqlDb
		default:
			return nil, edance.ErrUnsupportedDBType
		}
	}
	return db, nil
}
