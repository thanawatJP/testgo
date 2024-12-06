package models

type Blog struct {
	ID       uint   `gorm:"primaryKey"`
	Title    string `json:"title" form:"title" gorm:"size:255;not null"`
	Content  string `json:"content" form:"content" gorm:"type:text"`
	UserID   uint   `json:"user_id" form:"user_id" gorm:"not null"` // Foreign Key เชื่อมกับ User
	ImageURL string `json:"image_url" form:"image_url" gorm:"type:varchar(255)"`
}
