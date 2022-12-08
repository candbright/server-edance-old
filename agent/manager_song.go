package agent

import (
	"edance"
	"edance/db"
	"edance/db/domain"
	"github.com/candbright/gin-util/xlog"
	"github.com/google/uuid"
)

type SongManager struct {
	songs []domain.Song
	db    db.DB
}

func NewSongManager(db db.DB) (*SongManager, error) {
	manager := &SongManager{db: db}
	songs, err := db.ListAllSong()
	if err != nil {
		return nil, xlog.Wrap(err)
	}
	manager.songs = songs
	return manager, nil
}

func (manager *SongManager) ListAllSong() ([]domain.Song, error) {
	//TODO implement me
	panic("implement me")
}

func (manager *SongManager) ListSongByMode(mode string, modeDetail string) ([]domain.Song, error) {
	//TODO implement me
	panic("implement me")
}

func (manager *SongManager) GetSongById(songId string) (domain.Song, error) {
	//TODO implement me
	panic("implement me")
}

func (manager *SongManager) AddSong(song domain.Song) error {
	if song.Id == "" {
		randomId, err := uuid.NewRandom()
		if err != nil {
			return xlog.Wrap(edance.ErrRandomUuid(err))
		}
		song.Id = randomId.String()
	}
	err := manager.db.AddSong(song)
	if err != nil {
		return xlog.Wrap(err)
	}
	manager.songs = append(manager.songs, song)
	return nil
}

func (manager *SongManager) DeleteSong(songId string) error {
	//TODO implement me
	panic("implement me")
}
