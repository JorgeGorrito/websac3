package service

import (
	"websac3/app/domain/entity"
	"websac3/app/domain/errs"
	"websac3/app/port/out/message"
	"websac3/app/port/out/persistence"
	"websac3/app/port/out/persistence/db"
)

type GetCourseTopicsByCourseIDService struct {
	getCourseTopicsPort persistence.GetCourseTopicsPort
	persistenceManager  db.Manager
	msgProvider         message.Provider
}

func NewGetCourseTopicsByCourseIDService(
	getCourseTopicsPort persistence.GetCourseTopicsPort,
	persistenceManager db.Manager,
	msgProvider message.Provider,
) *GetCourseTopicsByCourseIDService {
	return &GetCourseTopicsByCourseIDService{
		getCourseTopicsPort: getCourseTopicsPort,
		persistenceManager:  persistenceManager,
		msgProvider:         msgProvider,
	}
}

func (s *GetCourseTopicsByCourseIDService) Execute(
	courseID uint,
	lang string,
) ([]entity.CourseTopic, *entity.Course, error) {
	var (
		err          error
		courseTopics []entity.CourseTopic = make([]entity.CourseTopic, 0)
		course       *entity.Course
	)

	err = s.persistenceManager.ExecuteNonTransactional(
		func(dbCtx db.Context) error {
			courseTopics, course, err = s.getCourseTopicsPort.GetTopicsWithCourseByCourseIDAndLang(courseID, lang, dbCtx)
			if err != nil {
				courseTopics = nil
				course = nil
				return err
			}

			if len(courseTopics) == 0 {
				return errs.NewNotFoundError(
					s.msgProvider.
						WithLang(lang).
						GetMessage("get_course_topics_by_course_id", "not_found"),
				)
			}

			return nil
		},
	)

	return courseTopics, course, err
}
