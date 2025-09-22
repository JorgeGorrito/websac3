package mapper

import (
	"websac3/adapter/in/web/request"
	"websac3/adapter/in/web/response"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/command"
)

func registerCourseMappers() {
	// CourseNature mappers
	RegisterMapFunc(
		func(cn *model.CourseNature) (entity.CourseNature, error) {
			var names []entity.CourseNatureName
			for _, name := range cn.Names {
				names = append(names, entity.CourseNatureName{
					ID:             name.ID,
					Lang:           name.Lang,
					Name:           name.Name,
					CourseNatureID: name.CourseNatureID,
				})
			}
			return entity.CourseNature{
				ID:    cn.ID,
				Names: names,
			}, nil
		},
	)

	// CourseNature mapper with language support
	RegisterMapFunc(
		func(cn *model.CourseNature) (entity.CourseNature, error) {
			return mapCourseNatureWithLanguage(cn, "en") // Default to English
		},
	)

	// CourseNature to ListCourseNaturesResponse mapper
	RegisterMapFunc(
		func(courseNatureEntity *entity.CourseNature) (response.ListCourseNaturesResponse, error) {
			var name string
			if len(courseNatureEntity.Names) > 0 {
				name = courseNatureEntity.Names[0].Name
			}
			return response.ListCourseNaturesResponse{
				ID:   courseNatureEntity.ID,
				Name: name,
			}, nil
		},
	)

	// CourseType mappers
	RegisterMapFunc(
		func(ct *model.CourseType) (entity.CourseType, error) {
			var names []entity.CourseTypeName
			for _, name := range ct.Names {
				names = append(names, entity.CourseTypeName{
					ID:           name.ID,
					Lang:         name.Lang,
					Name:         name.Name,
					CourseTypeID: name.CourseTypeID,
				})
			}
			return entity.CourseType{
				ID:    ct.ID,
				Names: names,
			}, nil
		},
	)

	// CourseType to ListCourseTypesResponse mapper
	RegisterMapFunc(
		func(courseTypeEntity *entity.CourseType) (response.ListCourseTypesResponse, error) {
			var name string
			if len(courseTypeEntity.Names) > 0 {
				name = courseTypeEntity.Names[0].Name
			}
			return response.ListCourseTypesResponse{
				ID:   courseTypeEntity.ID,
				Name: name,
			}, nil
		},
	)

	// CourseType mapper with language support
	RegisterMapFunc(
		func(ct *model.CourseType) (entity.CourseType, error) {
			return mapCourseTypeWithLanguage(ct, "en") // Default to English
		},
	)

	RegisterMapFunc(
		func(ct *model.CourseTopic) (entity.CourseTopic, error) {
			var topicPtr *entity.Topic
			if ct.Topic.ID != 0 {
				if topic, err := Map[model.Topic, entity.Topic](&ct.Topic); err == nil {
					topicPtr = &topic
				}
			}

			return entity.CourseTopic{
				CourseID:   ct.CourseID,
				TopicID:    ct.TopicID,
				Topic:      topicPtr,
				StudyHours: float32(ct.StudyHours),
			}, nil
		},
	)

	RegisterMapFunc(
		func(courseModel *model.Course) (entity.Course, error) {
			var courseTopics []entity.CourseTopic
			for _, ct := range courseModel.CourseTopics {
				mapped, err := Map[model.CourseTopic, entity.CourseTopic](&ct)
				if err != nil {
					return entity.Course{}, err
				}
				courseTopics = append(courseTopics, mapped)
			}

			var naturePtr *entity.CourseNature
			if courseModel.Nature.ID != 0 {
				if nature, err := Map[model.CourseNature, entity.CourseNature](&courseModel.Nature); err == nil {
					naturePtr = &nature
				}
			}

			var typePtr *entity.CourseType
			if courseModel.Type.ID != 0 {
				if courseType, err := Map[model.CourseType, entity.CourseType](&courseModel.Type); err == nil {
					typePtr = &courseType
				}
			}

			var degreeProgramPtr *entity.DegreeProgram
			if courseModel.DegreeProgram.ID != 0 {
				if dp, err := Map[model.DegreeProgram, entity.DegreeProgram](&courseModel.DegreeProgram); err == nil {
					degreeProgramPtr = &dp
				}
			}

			var userCreatorPtr *entity.User
			if courseModel.UserCreator.ID != 0 {
				if uc, err := Map[model.User, entity.User](&courseModel.UserCreator); err == nil {
					userCreatorPtr = &uc
				}
			}

			return entity.Course{
				ID:              courseModel.ID,
				Name:            courseModel.Name,
				Code:            courseModel.Code,
				Credits:         courseModel.Credits,
				PeriodNumber:    courseModel.PeriodNumber,
				NatureID:        courseModel.NatureID,
				Nature:          naturePtr,
				TypeID:          courseModel.TypeID,
				Type:            typePtr,
				IsCybersecurity: courseModel.IsCybersecurity,
				DegreeProgramID: courseModel.DegreeProgramID,
				DegreeProgram:   degreeProgramPtr,
				CourseTopics:    courseTopics,
				CreatedBy:       courseModel.CreatedBy,
				UserCreator:     userCreatorPtr,
			}, nil
		},
	)

	RegisterMapFunc(
		func(courseEntity *entity.Course) (model.Course, error) {
			return model.Course{
				ID:                          courseEntity.ID,
				Name:                        courseEntity.Name,
				Code:                        courseEntity.Code,
				Credits:                     courseEntity.Credits,
				PeriodNumber:                courseEntity.PeriodNumber,
				NatureID:                    courseEntity.NatureID,
				TypeID:                      courseEntity.TypeID,
				IsCybersecurity:             courseEntity.IsCybersecurity,
				ContainsCybersecurityTopics: courseEntity.ContainsCybersecurityTopics(),
				DegreeProgramID:             courseEntity.DegreeProgramID,
				CreatedBy:                   courseEntity.CreatedBy,
			}, nil
		},
	)

	RegisterMapFunc(
		func(request *request.CreateCourseRequest) (command.CreateCourseCommand, error) {
			var courseTopics []command.CreateCourseTopicCommand
			for _, topic := range request.CourseTopics {
				courseTopics = append(courseTopics, command.CreateCourseTopicCommand{
					TopicID:    topic.TopicID,
					StudyHours: topic.StudyHours,
				})
			}

			return command.CreateCourseCommand{
				Name:                        request.Name,
				Code:                        request.Code,
				Credits:                     request.Credits,
				PeriodNumber:                request.PeriodNumber,
				NatureID:                    request.NatureID,
				TypeID:                      request.TypeID,
				IsCybersecurity:             request.IsCybersecurity,
				ContainsCybersecurityTopics: request.ContainsCybersecurityTopics,
				DegreeProgramID:             request.DegreeProgramID,
				CourseTopics:                courseTopics,
			}, nil
		},
	)

	RegisterMapFunc(
		func(command *command.CreateCourseCommand) (entity.Course, error) {
			var courseTopics []entity.CourseTopic
			for _, t := range command.CourseTopics {
				courseTopics = append(courseTopics, entity.CourseTopic{
					TopicID:    t.TopicID,
					StudyHours: float32(t.StudyHours),
				})
			}
			return entity.Course{
				Name:            command.Name,
				Code:            command.Code,
				Credits:         command.Credits,
				PeriodNumber:    command.PeriodNumber,
				NatureID:        command.NatureID,
				TypeID:          command.TypeID,
				IsCybersecurity: command.IsCybersecurity,
				DegreeProgramID: command.DegreeProgramID,
				CourseTopics:    courseTopics,
				CreatedBy:       command.CreatedBy,
			}, nil
		},
	)

	RegisterMapFunc(
		func(courseModel *model.Course) (response.ListCourseByDegreeProgramResponse, error) {
			var natureName, typeName, degreeProgramName, creatorName string

			// Get nature name if available
			if courseModel.Nature.ID != 0 && len(courseModel.Nature.Names) > 0 {
				// Get the first available name (could be filtered by language if needed)
				natureName = courseModel.Nature.Names[0].Name
			}

			// Get type name if available
			if courseModel.Type.ID != 0 && len(courseModel.Type.Names) > 0 {
				// Get the first available name (could be filtered by language if needed)
				typeName = courseModel.Type.Names[0].Name
			}

			// Get degree program name if available
			if courseModel.DegreeProgram.ID != 0 {
				degreeProgramName = courseModel.DegreeProgram.Name
			}

			// Get creator name if available
			if courseModel.UserCreator.ID != 0 && courseModel.UserCreator.Person.ID != 0 {
				creatorName = courseModel.UserCreator.Person.Name + " " + courseModel.UserCreator.Person.Lastname
			}

			return response.ListCourseByDegreeProgramResponse{
				ID:                          courseModel.ID,
				Name:                        courseModel.Name,
				Code:                        courseModel.Code,
				Credits:                     courseModel.Credits,
				PeriodNumber:                courseModel.PeriodNumber,
				NatureID:                    courseModel.NatureID,
				NatureName:                  natureName,
				TypeID:                      courseModel.TypeID,
				TypeName:                    typeName,
				IsCybersecurity:             courseModel.IsCybersecurity,
				ContainsCybersecurityTopics: len(courseModel.CourseTopics) > 0,
				DegreeProgramID:             courseModel.DegreeProgramID,
				DegreeProgramName:           degreeProgramName,
				CreatedBy:                   courseModel.CreatedBy,
				CreatorName:                 creatorName,
			}, nil
		},
	)

	RegisterMapFunc(
		func(courseEntity *entity.Course) (response.GetCourseByIDResponse, error) {
			var natureName, typeName, degreeProgramName, creatorName string

			// Get nature name if available
			if courseEntity.Nature != nil && len(courseEntity.Nature.Names) > 0 {
				natureName = courseEntity.Nature.Names[0].Name
			}

			// Get type name if available
			if courseEntity.Type != nil && len(courseEntity.Type.Names) > 0 {
				typeName = courseEntity.Type.Names[0].Name
			}

			// Get degree program name if available
			if courseEntity.DegreeProgram != nil {
				degreeProgramName = courseEntity.DegreeProgram.Name
			}

			// Get creator name if available
			if courseEntity.UserCreator != nil && courseEntity.UserCreator.Person != nil {
				creatorName = courseEntity.UserCreator.Person.Name + " " + courseEntity.UserCreator.Person.Lastname
			}

			// Map course topics
			var courseTopicsResponse []response.CourseTopicResponse
			for _, courseTopic := range courseEntity.CourseTopics {
				var topicName string
				if courseTopic.Topic != nil {
					topicName = courseTopic.Topic.Name
				}
				courseTopicsResponse = append(courseTopicsResponse, response.CourseTopicResponse{
					TopicID:    courseTopic.TopicID,
					TopicName:  topicName,
					StudyHours: courseTopic.StudyHours,
				})
			}

			return response.GetCourseByIDResponse{
				ID:                          courseEntity.ID,
				Name:                        courseEntity.Name,
				Code:                        courseEntity.Code,
				Credits:                     courseEntity.Credits,
				PeriodNumber:                courseEntity.PeriodNumber,
				NatureID:                    courseEntity.NatureID,
				NatureName:                  natureName,
				TypeID:                      courseEntity.TypeID,
				TypeName:                    typeName,
				IsCybersecurity:             courseEntity.IsCybersecurity,
				ContainsCybersecurityTopics: courseEntity.ContainsCybersecurityTopics(),
				DegreeProgramID:             courseEntity.DegreeProgramID,
				DegreeProgramName:           degreeProgramName,
				CreatedBy:                   courseEntity.CreatedBy,
				CreatorName:                 creatorName,
				CourseTopics:                courseTopicsResponse,
			}, nil
		},
	)

	// Language-specific Course mapper
	RegisterMapFunc(
		func(courseModel *model.Course) (entity.Course, error) {
			return mapCourseWithLanguage(courseModel, "en") // Default to English
		},
	)
}

// Helper function to get name by language
func getNameByLanguage(names []entity.CourseNatureName, lang string) string {
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

func getTypeNameByLanguage(names []entity.CourseTypeName, lang string) string {
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

// Helper function to map course nature with specific language
func mapCourseNatureWithLanguage(cn *model.CourseNature, lang string) (entity.CourseNature, error) {
	// Get name by language
	var name string
	for _, cnName := range cn.Names {
		if cnName.Lang == lang {
			name = cnName.Name
			break
		}
	}
	// If no name found for the language, use the first available
	if name == "" && len(cn.Names) > 0 {
		name = cn.Names[0].Name
	}

	// Create names array with the selected name
	var names []entity.CourseNatureName
	if name != "" {
		names = append(names, entity.CourseNatureName{
			ID:             cn.Names[0].ID,
			Lang:           lang,
			Name:           name,
			CourseNatureID: cn.ID,
		})
	}

	return entity.CourseNature{
		ID:    cn.ID,
		Names: names,
	}, nil
}

// Helper function to map course type with specific language
func mapCourseTypeWithLanguage(ct *model.CourseType, lang string) (entity.CourseType, error) {
	// Get name by language
	var name string
	for _, ctName := range ct.Names {
		if ctName.Lang == lang {
			name = ctName.Name
			break
		}
	}
	// If no name found for the language, use the first available
	if name == "" && len(ct.Names) > 0 {
		name = ct.Names[0].Name
	}

	// Create names array with the selected name
	var names []entity.CourseTypeName
	if name != "" {
		names = append(names, entity.CourseTypeName{
			ID:           ct.Names[0].ID,
			Lang:         lang,
			Name:         name,
			CourseTypeID: ct.ID,
		})
	}

	return entity.CourseType{
		ID:    ct.ID,
		Names: names,
	}, nil
}

// Language-specific mapper functions
func GetCourseNatureMapperWithLanguage(lang string) func(*model.CourseNature) (entity.CourseNature, error) {
	return func(cn *model.CourseNature) (entity.CourseNature, error) {
		return mapCourseNatureWithLanguage(cn, lang)
	}
}

func GetCourseTypeMapperWithLanguage(lang string) func(*model.CourseType) (entity.CourseType, error) {
	return func(ct *model.CourseType) (entity.CourseType, error) {
		return mapCourseTypeWithLanguage(ct, lang)
	}
}

func GetCourseTopicMapperWithLanguage(lang string) func(*model.CourseTopic) (entity.CourseTopic, error) {
	return func(ct *model.CourseTopic) (entity.CourseTopic, error) {
		var topicPtr *entity.Topic
		if ct.Topic.ID != 0 {
			topic, err := mapTopicWithLanguage(&ct.Topic, lang)
			if err == nil {
				topicPtr = &topic
			}
		}

		return entity.CourseTopic{
			CourseID:   ct.CourseID,
			TopicID:    ct.TopicID,
			Topic:      topicPtr,
			StudyHours: float32(ct.StudyHours),
		}, nil
	}
}

func GetCourseMapperWithLanguage(lang string) func(*model.Course) (entity.Course, error) {
	return func(courseModel *model.Course) (entity.Course, error) {
		return mapCourseWithLanguage(courseModel, lang)
	}
}

// Helper function to map course with language support
func mapCourseWithLanguage(courseModel *model.Course, lang string) (entity.Course, error) {
	var natureName, typeName string

	// Get nature name by language
	if courseModel.Nature.ID != 0 {
		natureName = getCourseNatureNameByLanguageFromModel(&courseModel.Nature, lang)
	}

	// Get type name by language
	if courseModel.Type.ID != 0 {
		typeName = getCourseTypeNameByLanguageFromModel(&courseModel.Type, lang)
	}

	// Map degree program
	var degreeProgram *entity.DegreeProgram
	if courseModel.DegreeProgram.ID != 0 {
		dp, err := Map[model.DegreeProgram, entity.DegreeProgram](&courseModel.DegreeProgram)
		if err != nil {
			return entity.Course{}, err
		}
		degreeProgram = &dp
	}

	// Map user creator
	var userCreator *entity.User
	if courseModel.UserCreator.ID != 0 {
		uc, err := Map[model.User, entity.User](&courseModel.UserCreator)
		if err != nil {
			return entity.Course{}, err
		}
		userCreator = &uc
	}

	// Map course topics with language support
	var courseTopics []entity.CourseTopic
	for _, courseTopic := range courseModel.CourseTopics {
		courseTopicEntity, err := mapCourseTopicWithLanguage(&courseTopic, lang)
		if err != nil {
			return entity.Course{}, err
		}
		courseTopics = append(courseTopics, courseTopicEntity)
	}

	return entity.Course{
		ID:              courseModel.ID,
		Name:            courseModel.Name,
		Code:            courseModel.Code,
		Credits:         courseModel.Credits,
		PeriodNumber:    courseModel.PeriodNumber,
		NatureID:        courseModel.NatureID,
		Nature:          &entity.CourseNature{ID: courseModel.NatureID, Names: []entity.CourseNatureName{{Lang: lang, Name: natureName}}},
		TypeID:          courseModel.TypeID,
		Type:            &entity.CourseType{ID: courseModel.TypeID, Names: []entity.CourseTypeName{{Lang: lang, Name: typeName}}},
		IsCybersecurity: courseModel.IsCybersecurity,
		DegreeProgramID: courseModel.DegreeProgramID,
		DegreeProgram:   degreeProgram,
		CourseTopics:    courseTopics,
		CreatedBy:       courseModel.CreatedBy,
		UserCreator:     userCreator,
	}, nil
}

// Helper function to map course topic with language support
func mapCourseTopicWithLanguage(courseTopicModel *model.CourseTopic, lang string) (entity.CourseTopic, error) {
	var topicName, knowledgeAreaName string
	var knowledgeAreaID uint

	// Get topic name by language
	if courseTopicModel.Topic.ID != 0 {
		topicName = getTopicNameByLanguageFromModel(&courseTopicModel.Topic, lang)
		knowledgeAreaID = courseTopicModel.Topic.KnowledgeAreaID

		// Get knowledge area name by language
		if courseTopicModel.Topic.KnowledgeArea.ID != 0 {
			knowledgeAreaName = getKnowledgeAreaNameByLanguageFromModel(&courseTopicModel.Topic.KnowledgeArea, lang)
		}
	}

	// Create topic entity with language-specific name
	var topic *entity.Topic
	if courseTopicModel.Topic.ID != 0 {
		topic = &entity.Topic{
			ID:              courseTopicModel.Topic.ID,
			Name:            topicName,
			KnowledgeAreaID: knowledgeAreaID,
			KnowledgeArea: &entity.KnowledgeArea{
				ID:   knowledgeAreaID,
				Name: knowledgeAreaName,
			},
		}
	}

	return entity.CourseTopic{
		CourseID:   courseTopicModel.CourseID,
		TopicID:    courseTopicModel.TopicID,
		Topic:      topic,
		StudyHours: float32(courseTopicModel.StudyHours),
	}, nil
}

// Helper functions to get names by language from models
func getTopicNameByLanguageFromModel(topicModel *model.Topic, lang string) string {
	for _, name := range topicModel.Names {
		if name.Lang == lang {
			return name.Name
		}
	}
	// If no name found for the language, use the first available
	if len(topicModel.Names) > 0 {
		return topicModel.Names[0].Name
	}
	return ""
}

func getKnowledgeAreaNameByLanguageFromModel(kaModel *model.KnowledgeArea, lang string) string {
	for _, name := range kaModel.Names {
		if name.Lang == lang {
			return name.Name
		}
	}
	// If no name found for the language, use the first available
	if len(kaModel.Names) > 0 {
		return kaModel.Names[0].Name
	}
	return ""
}

func getCourseNatureNameByLanguageFromModel(cnModel *model.CourseNature, lang string) string {
	for _, name := range cnModel.Names {
		if name.Lang == lang {
			return name.Name
		}
	}
	// If no name found for the language, use the first available
	if len(cnModel.Names) > 0 {
		return cnModel.Names[0].Name
	}
	return ""
}

func getCourseTypeNameByLanguageFromModel(ctModel *model.CourseType, lang string) string {
	for _, name := range ctModel.Names {
		if name.Lang == lang {
			return name.Name
		}
	}
	// If no name found for the language, use the first available
	if len(ctModel.Names) > 0 {
		return ctModel.Names[0].Name
	}
	return ""
}
