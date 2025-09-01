package model

type CourseTypeName struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Lang string `gorm:"type:varchar(2);not null;index" json:"lang"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`

	CourseTypeID uint       `gorm:"not null;index;" json:"course_type_id"`
	CourseType   CourseType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
