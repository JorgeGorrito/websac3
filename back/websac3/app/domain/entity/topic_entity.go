package entity

type Topic struct {
	ID   uint
	Name string

	KnowledgeAreaID uint
	KnowledgeArea   *KnowledgeArea
}
