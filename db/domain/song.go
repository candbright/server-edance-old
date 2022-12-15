package domain

type Song struct {
	Id          string `gorm:"column:id;primaryKey" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
	Mode        string `gorm:"column:mode" json:"mode"`
	ModeDetail  string `gorm:"column:mode_detail" json:"mode_detail"`
	Difficult   string `gorm:"column:difficult" json:"difficult"`
	ImageSrc    string `gorm:"column:image_src" json:"image_src"`
	Url         string `gorm:"column:url" json:"url"`
}

func (song Song) TableName() string {
	return "domain_song"
}
