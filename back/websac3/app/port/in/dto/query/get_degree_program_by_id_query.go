package query

// GetDegreeProgramByIDQuery representa la consulta para obtener el detalle de un programa de grado por ID
type GetDegreeProgramByIDQuery struct {
	DegreeProgramID uint `json:"degree_program_id"`
}
