package models

// Bootstrap :
type Bootstrap struct {
	ID          int    `json:"id" example:"123547"`
	Title       string `json:"title" example:"Ingeniero en Sistemas" validate:"nonzero"`
	Description string `json:"description" example:"Se busca ingeniero en sistemas computacionales" validate:"nonzero,max=50"`
}
