package model

type DurationUnitName struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"varchar(64); not null; unique" json:"name"`
	Lang       string `gorm:"varchar(2); not null; index" json:"lang"`
	DurationID uint   `json:"duration_unit_id"`
}
