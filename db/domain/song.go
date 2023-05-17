package domain

type Song struct {
	Id           string `gorm:"column:id;primaryKey" json:"id"`
	SongName     string `gorm:"column:song_name" json:"song_name"`
	SingerName   string `gorm:"column:singer_name" json:"singer_name"`
	Index        int    `gorm:"column:index" json:"index"`
	Difficult    string `gorm:"column:difficult" json:"difficult"`
	Mode         string `gorm:"column:mode" json:"mode"`
	ModeDetail   string `gorm:"column:mode_detail" json:"mode_detail"`
	DesignerName string `gorm:"column:designer_name" json:"designer_name"`
	Description  string `gorm:"column:description" json:"description"`
}

func (song Song) TableName() string {
	return "domain_song"
}
