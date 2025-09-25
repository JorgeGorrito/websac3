package mapper

import (
	"websac3/adapter/in/web/response"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
)

func registerFormationLevelMappers() {
	// FormationLevel mappers
	RegisterMapFunc(
		func(fl *model.FormationLevel) (entity.FormationLevel, error) {
			var name string
			if len(fl.Names) > 0 {
				name = fl.Names[0].Name
			}
			return entity.FormationLevel{
				ID:   fl.ID,
				Name: name,
			}, nil
		},
	)

	// FormationLevel to ListFormationLevelsResponse mapper
	RegisterMapFunc(
		func(formationLevelEntity *entity.FormationLevel) (response.ListFormationLevelsResponse, error) {
			return response.ListFormationLevelsResponse{
				ID:   formationLevelEntity.ID,
				Name: formationLevelEntity.Name,
			}, nil
		},
	)
}

// Helper function to get formation level name by language
func getFormationLevelNameByLanguage(names []model.FormationLevelName, lang string) string {
	for _, name := range names {
		if name.Lang == lang {
			return name.Name
		}
	}
	// If no name found for the language, return the first available name
	if len(names) > 0 {
		return names[0].Name
	}
	return ""
}

// Helper function to map formation level with specific language
func mapFormationLevelWithLanguage(fl *model.FormationLevel, lang string) (entity.FormationLevel, error) {
	// Get name by language
	name := getFormationLevelNameByLanguage(fl.Names, lang)

	return entity.FormationLevel{
		ID:   fl.ID,
		Name: name,
	}, nil
}

// Language-specific mapper function
func GetFormationLevelMapperWithLanguage(lang string) func(*model.FormationLevel) (entity.FormationLevel, error) {
	return func(fl *model.FormationLevel) (entity.FormationLevel, error) {
		return mapFormationLevelWithLanguage(fl, lang)
	}
}
