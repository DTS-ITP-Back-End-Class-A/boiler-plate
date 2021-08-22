package services

import (
	"testing/models"
	repo "testing/repositories"
)

func ServicesGetStudent() (result models.Response) {

	// di service biasanya kita nglakuin call repositories and do logic
	// statusActive := lang.StatusActive

	repo.RepoGetStudent()
	return result
}
