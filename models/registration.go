package models

type Registration struct {
	ID      int64 `json:"id" gorm:"primaryKey" gorm:"autoIncrement"`
	UserID  int64 `json:"userId"`
	EventID int64 `json:"eventId"`

	User  User  `json:"user" gorm:"foreignKey:UserID"`
	Event Event `json:"event" gorm:"foreignKey:EventID"`
}
