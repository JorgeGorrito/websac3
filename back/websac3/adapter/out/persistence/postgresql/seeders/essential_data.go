package seeders

import (
	"websac3/app/port/out/persistence"
)

type essentialData struct {
	essentialSeeders []Seeder
}

func EssentialData() Seeder {
	return &essentialData{
		essentialSeeders: []Seeder{
			Roles(),
			AccessRequestStatuses(),
			IdentificationTypes(),
			Municipalities(),
			Departments(),
			InstitutionalCategories(),
			Ownerships(),
			HigherEducationInstitutions(),
		},
	}
}

func (e *essentialData) Seed(ctx persistence.Context) error {
	for _, seeder := range e.essentialSeeders {
		if err := seeder.Seed(ctx); err != nil {
			return err
		}
	}
	return nil
}
