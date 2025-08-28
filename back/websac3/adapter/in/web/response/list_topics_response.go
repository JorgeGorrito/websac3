package response

type ListTopicResponse struct {
	ID                uint   `json:"id"`
	Name              string `json:"name"`
	KnowledgeAreaID   uint   `json:"knowledge_area_id"`
	KnowledgeAreaName string `json:"knowledge_area_name"`
}
