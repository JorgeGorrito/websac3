package mapper

import (
	"websac3/adapter/in/web/request"
	"websac3/adapter/out/persistence/postgresql/model"
	"websac3/app/domain/entity"
	"websac3/app/port/in/dto/command"
)

func registerCourseMappers() {
	RegisterMapFunc(
		func(courseModel *model.Course) (entity.Course, error) {
			return entity.Course{
				ID:                          courseModel.ID,
				Name:                        courseModel.Name,
				Code:                        courseModel.Code,
				Credits:                     courseModel.Credits,
				PeriodNumber:                courseModel.PeriodNumber,
				NatureID:                    courseModel.NatureID,
				TypeID:                      courseModel.TypeID,
				IsCybersecurity:             courseModel.IsCybersecurity,
				ContainsCybersecurityTopics: courseModel.ContainsCybersecurityTopics,
				DegreeProgramID:             courseModel.DegreeProgramID,
				CreatedBy:                   courseModel.CreatedBy,
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
				ContainsCybersecurityTopics: courseEntity.ContainsCybersecurityTopics,
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
			return entity.Course{
				Name:                        command.Name,
				Code:                        command.Code,
				Credits:                     command.Credits,
				PeriodNumber:                command.PeriodNumber,
				NatureID:                    command.NatureID,
				TypeID:                      command.TypeID,
				IsCybersecurity:             command.IsCybersecurity,
				ContainsCybersecurityTopics: command.ContainsCybersecurityTopics,
				DegreeProgramID:             command.DegreeProgramID,
				CreatedBy:                   command.CreatedBy,
			}, nil
		},
	)
}
