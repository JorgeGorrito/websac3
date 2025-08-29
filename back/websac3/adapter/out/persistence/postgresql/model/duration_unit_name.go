package model

type DurationUnitName struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Lang string `gorm:"type:varchar(2);not null;index" json:"lang"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`

	DurationUnitID uint         `gorm:"not null;index;" json:"duration_unit_id"`
	DurationUnit   DurationUnit `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
