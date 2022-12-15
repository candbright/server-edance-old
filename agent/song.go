package agent

import (
	"github.com/candbright/edance/db/domain"
	"github.com/candbright/util/xgin"
	"github.com/candbright/util/xlog"
	"github.com/gin-gonic/gin"
)

type SongW struct {
	Song domain.Song `json:"song"`
}

func restListAllSong(context *gin.Context) {
	xgin.GET(context, func(pathParams map[string]string) (interface{}, error) {
		songs, err := songManager.ListAllSong()
		if err != nil {
			return nil, xlog.Wrap(err)
		}
		return songs, nil
	})
}

func restGetSongById(context *gin.Context) {

}
func restAddSong(context *gin.Context) {
	xgin.POST(context, func(receive interface{}, pathParams map[string]string) (interface{}, error) {
		receiveData := receive.(*SongW)
		err := songManager.AddSong(receiveData.Song)
		if err != nil {
			return nil, xlog.Wrap(err)
		}
		return songManager.songs[0], nil
	}, &SongW{})
}

func restUpdateSong(context *gin.Context) {

}

func restDeleteSong(context *gin.Context) {
	xgin.DELETE(context, func(pathParams map[string]string) (interface{}, error) {
		id := pathParams["song_id"]
		err := songManager.DeleteSong(id)
		if err != nil {
			return nil, xlog.Wrap(err)
		}
		return nil, nil
	})
}
