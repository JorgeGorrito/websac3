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
