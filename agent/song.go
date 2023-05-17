package agent

import (
	"github.com/candbright/edance"
	"github.com/candbright/edance/db/domain"
	"github.com/candbright/util/xgin"
	"github.com/candbright/util/xlog"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SongW struct {
	Song domain.Song `json:"song"`
}

type SongsW struct {
	Songs []domain.Song `json:"songs"`
}

func restListSong(context *gin.Context) {
	songName := context.Query("song_name")
	xgin.GET(context, func() (interface{}, error) {
		songs, err := songManager.ListAllSong()
		if err != nil {
			return nil, xlog.Wrap(err)
		}
		result := make([]domain.Song, 0)
		for _, song := range songs {
			if songName != "" && song.SongName != songName {
				continue
			}
			result = append(result, song)
		}
		return SongsW{result}, nil
	})
}

func restGetSongById(context *gin.Context) {
	songId := context.Param("song_id")
	xgin.GET(context, func() (interface{}, error) {
		song, err := songManager.GetSongById(songId)
		if err != nil {
			return nil, xlog.Wrap(err)
		}
		return SongW{song}, nil
	})
}

func restAddSong(context *gin.Context) {
	xgin.POST(context, func(receive interface{}) (interface{}, error) {
		receiveData := receive.(*SongW).Song
		if receiveData.Id == "" {
			randomId, err := uuid.NewRandom()
			if err != nil {
				return nil, xlog.Wrap(edance.ErrRandomUuid(err))
			}
			receiveData.Id = randomId.String()
		}
		err := songManager.AddSong(receiveData)
		if err != nil {
			return nil, xlog.Wrap(err)
		}
		after, err := songManager.GetSongById(receiveData.Id)
		if err != nil {
			return nil, xlog.Wrap(err)
		}
		return SongW{after}, nil
	}, &SongW{})
}

func restUpdateSong(context *gin.Context) {

}

func restDeleteSong(context *gin.Context) {
	id := context.Param("song_id")
	xgin.DELETE(context, func() (interface{}, error) {
		err := songManager.DeleteSong(id)
		if err != nil {
			return nil, xlog.Wrap(err)
		}
		return nil, nil
	})
}
