package repositories

import (
	"api-bootstrap-echo/models"
)

// IBootstrapRepository :
type IBootstrapRepository interface {
	Get(int) (models.Bootstrap, error)
	Save(models.Bootstrap) (bool, error)
}
