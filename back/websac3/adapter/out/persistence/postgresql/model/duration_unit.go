package model

type DurationUnit struct {
	ID    uint               `gorm:"primaryKey" json:"id"`
	Names []DurationUnitName `gorm:"foreignKey:DurationID" json:"names"`
}
