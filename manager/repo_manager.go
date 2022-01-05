package manager

import (
	"enigmacamp.com/completetesting/repository"
)

type RepoManager interface {
	StudentRepo() repository.IStudentRepository
}

type repoManager struct {
	infra Infra
}

func (rm *repoManager) StudentRepo() repository.IStudentRepository {
	return repository.NewStudentRepository(rm.infra.SqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{infra: infra}
}
