package seeders

import "websac3/app/port/out/persistence/db"

type essentialData struct {
	essentialSeeders []Seeder
}

func EssentialData() Seeder {
	return &essentialData{
		essentialSeeders: []Seeder{
			AccessRequestStatuses(),
			IdentificationTypes(),
			Municipalities(),
			Departments(),
			InstitutionalCategories(),
			Ownerships(),
			HigherEducationInstitutions(),
			Modules(),
			Actions(),
			Permissions(),
			Roles(),
			DefaultUsers(),
			KnowledgeAreas(),
			Topics(),
			DurationUnits(),
		},
	}
}

func (e *essentialData) Seed(ctx db.Context) error {
	for _, seeder := range e.essentialSeeders {
		if err := seeder.Seed(ctx); err != nil {
			return err
		}
	}
	return nil
}
