package agent

import (
	"errors"
	"github.com/candbright/edance/db"
	"github.com/candbright/edance/db/domain"
	"github.com/candbright/util/xlog"
	"gorm.io/gorm"
)

type SongManager struct {
	db db.DB
}

func NewSongManager(db db.DB) (*SongManager, error) {
	manager := &SongManager{db: db}
	return manager, nil
}

func (manager *SongManager) ListAllSong() ([]domain.Song, error) {
	return manager.db.ListAllSong()
}

func (manager *SongManager) ListSongByMode(mode string, modeDetail string) ([]domain.Song, error) {
	return manager.db.ListSongByMode(mode, modeDetail)
}

func (manager *SongManager) GetSongById(songId string) (domain.Song, error) {
	return manager.db.GetSongById(songId)
}

func (manager *SongManager) AddSong(song domain.Song) error {
	_, err := manager.GetSongById(song.Id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return xlog.Wrap(err)
	}
	err = manager.db.AddSong(song)
	if err != nil {
		return xlog.Wrap(err)
	}
	return nil
}

func (manager *SongManager) DeleteSong(songId string) error {
	return manager.db.DeleteSong(songId)
}
