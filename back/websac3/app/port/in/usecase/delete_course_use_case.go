package usecase

type DeleteCourseUseCase interface {
	Execute(courseID uint, userID uint, lang string) error
}
