package domain

type OfficialSong struct {
	Song
	ImageSrc string `gorm:"column:image_src" json:"image_src"`
	VideoUrl string `gorm:"column:video_url" json:"video_url"`
}

func (song OfficialSong) TableName() string {
	return "domain_song_official"
}
