package mysql

import (
	"edance/db/domain"
	"github.com/candbright/gin-util/xlog"
)

func (db DB) ListAllSong() ([]domain.Song, error) {
	var result []domain.Song
	if err := db.Find(&result).Error; err != nil {
		return nil, xlog.Wrap(db.Error)
	}
	return result, nil
}

func (db DB) ListSongByMode(mode string, modeDetail string) ([]domain.Song, error) {
	var result []domain.Song
	if err := db.Where("MODE = ?", mode).Where("MODE_DETAIL = ?", modeDetail).Find(&result).Error; err != nil {
		return nil, xlog.Wrap(db.Error)
	}
	return result, nil
}

func (db DB) GetSongById(songId string) (domain.Song, error) {
	var result domain.Song
	if err := db.Where("ID = ?", songId).Take(&result).Error; err != nil {
		return domain.Song{}, xlog.Wrap(db.Error)
	}
	return result, nil
}

func (db DB) AddSong(song domain.Song) error {
	if err := db.Create(&song).Error; err != nil {
		return xlog.Wrap(err)
	}
	return nil
}

func (db DB) DeleteSong(songId string) error {
	if err := db.Where("ID = ?", songId).Delete(&domain.Song{}).Error; err != nil {
		return xlog.Wrap(err)
	}
	return nil
}
