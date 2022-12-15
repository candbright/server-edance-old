package mysql

import (
	"github.com/candbright/edance/db/domain"
	"github.com/candbright/util/xlog"
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
	if err := db.Where("mode = ?", mode).Where("mode_detail = ?", modeDetail).Find(&result).Error; err != nil {
		return nil, xlog.Wrap(db.Error)
	}
	return result, nil
}

func (db DB) GetSongById(songId string) (domain.Song, error) {
	var result domain.Song
	if err := db.Where("id = ?", songId).Take(&result).Error; err != nil {
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
	if err := db.Where("id = ?", songId).Or("name = ?", songId).Delete(&domain.Song{}).Error; err != nil {
		return xlog.Wrap(err)
	}
	return nil
}
