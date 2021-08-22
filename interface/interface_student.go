package interfaces

import "testing/models"

type ServiceStudent interface {
	ServicesGetStudent() (result models.Response)
}
