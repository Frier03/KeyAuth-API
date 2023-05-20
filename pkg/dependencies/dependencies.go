package dependencies

import (
	"github.com/Frier03/KeyAuth-API/pkg/services"
)

type Dependencies struct {
	BadgerService *services.BadgerService
}

// Function to create new instance of Dependencies
func NewDependencies() (*Dependencies, error) {

	// Create badger service dep
	badgerService, err := services.NewBadgerService()
	if err != nil {
		return nil, err
	}

	deps := &Dependencies{
		BadgerService: badgerService,
	}

	return deps, nil
}
