package usecase

import (
	"enigmacamp.com/completetesting/model"
	"enigmacamp.com/completetesting/repository"
)

type IStudentUseCase interface {
	NewRegistration(student model.Student) (*model.Student, error)
	FindStudentInfoById(idCard string) (*model.Student, error)
}

type StudentUseCase struct {
	repo repository.IStudentRepository
}

func NewStudentUseCase(studentRepository repository.IStudentRepository) IStudentUseCase {
	return &StudentUseCase{studentRepository}
}

func (s *StudentUseCase) NewRegistration(student model.Student) (*model.Student, error) {
	return s.repo.CreateOne(student)
}

func (s *StudentUseCase) FindStudentInfoById(idCard string) (*model.Student, error) {
	return s.repo.GetOneById(idCard)
}
