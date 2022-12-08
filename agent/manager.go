package agent

import "edance/db"

var (
	songManager *SongManager
)

func InitManager() {
	var err error
	db, err := db.GetDb()
	if err != nil {
		panic(err)
	}
	songManager, err = NewSongManager(db)
	if err != nil {
		panic(err)
	}
}
