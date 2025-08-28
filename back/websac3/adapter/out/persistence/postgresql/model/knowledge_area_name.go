package model

type KnowledgeAreaName struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`
	Lang string `gorm:"type:varchar(2);not null;index" json:"lang"`

	KnowledgeAreaID uint          `gorm:"not null" json:"ka_id"`
	KnowledgeArea   KnowledgeArea `gorm:"foreignKey:KnowledgeAreaID" json:"-"`
}
