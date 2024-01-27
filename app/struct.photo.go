package app

type Photo struct {
	ID       uint   `gorm:"primaryKey"`
	Title    string `gorm:"not null"`
	Caption  string `gorm:"not null"`
	PhotoUrl string `gorm:"not null"`
	UserID   uint   `gorm:"foreignKey:UserID"`
}
