package manager

import "enigmacamp.com/completetesting/usecase"

type UseCaseManager interface {
	StudentUseCase() usecase.IStudentUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (uc *useCaseManager) StudentUseCase() usecase.IStudentUseCase {
	return usecase.NewStudentUseCase(uc.repo.StudentRepo())
}
func NewUseCaseManger(infra Infra) UseCaseManager {
	return &useCaseManager{repo: NewRepoManager(infra)}
}