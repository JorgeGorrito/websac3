package dto

type HigherEducationInstitution struct {
	Snies                 string `json:"snies"`
	SniesParent           string `json:"snies_parent"`
	Name                  string `json:"name"`
	Ownership             string `json:"ownership"`
	InstitutionalCategory string `json:"institutional_category"`
	Municipality          string `json:"municipality"`
	Department            string `json:"department"`
}
