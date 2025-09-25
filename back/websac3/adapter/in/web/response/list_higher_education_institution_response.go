package response

type ListHigherEducationInstitutionResponse struct {
	Snies        uint   `json:"snies"`
	Name         string `json:"name"`
	Department   string `json:"department"`
	Municipality string `json:"municipality"`
}
