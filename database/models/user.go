package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:100;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Blogs    []Blog `gorm:"foreignKey:UserID"` // ความสัมพันธ์ One-to-Many
}
