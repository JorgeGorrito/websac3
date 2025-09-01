package model

type CourseNatureName struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Lang string `gorm:"type:varchar(2);not null;index" json:"lang"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`

	CourseNatureID uint         `gorm:"not null;index;" json:"course_nature_id"`
	CourseNature   CourseNature `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
