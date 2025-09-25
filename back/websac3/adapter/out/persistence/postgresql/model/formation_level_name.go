package model

type FormationLevelName struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Lang string `gorm:"type:varchar(2);not null;index" json:"lang"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`

	FormationLevelID uint           `gorm:"not null;index;" json:"formation_level_id"`
	FormationLevel   FormationLevel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
