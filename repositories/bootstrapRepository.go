package repositories

import (
	"api-bootstrap-echo/models"
)

// NewBootstrapRepository :
func NewBootstrapRepository() *BootstrapRepository {
	return &BootstrapRepository{}
}

// BootstrapRepository :
type BootstrapRepository struct {
}

// Get :
func (BootstrapRepository) Get(id int) (models.Bootstrap, error) {
	bootstrap := models.Bootstrap{ID: id, Title: "Ingeniero en Sistemas", Description: "Se busca ingeniero en sistemas"}
	return bootstrap, nil
}

// Save :
func (BootstrapRepository) Save(bootstrap models.Bootstrap) (bool, error) {
	return true, nil
}
