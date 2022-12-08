package domain

type Song struct {
	Id          string `gorm:"column:ID;primaryKey" json:"id"`
	Name        string `gorm:"column:NAME" json:"name"`
	Description string `gorm:"column:DESCRIPTION" json:"description"`
	Mode        string `gorm:"column:MODE" json:"mode"`
	ModeDetail  string `gorm:"column:MODE_DETAIL" json:"mode_detail"`
	Difficult   string `gorm:"column:DIFFICULT" json:"difficult"`
	ImageSrc    string `gorm:"column:IMAGE_SRC" json:"image_src"`
	Url         string `gorm:"column:URL" json:"url"`
}

func (song Song) TableName() string {
	return "SONG"
}
