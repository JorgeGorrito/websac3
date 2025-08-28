package model

type TopicName struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Lang string `gorm:"type:varchar(2);not null;index" json:"lang"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`

	TopicID uint  `gorm:"not null;index;" json:"topic_id"`
	Topic   Topic `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
