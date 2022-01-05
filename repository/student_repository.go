package repository

import (
	"enigmacamp.com/completetesting/model"
	"enigmacamp.com/completetesting/util/logger"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type IStudentRepository interface {
	GetAll() ([]model.Student, error)
	GetOneByName(name string) ([]model.Student, error)
	GetOneById(idCard string) (*model.Student, error)
	CreateOne(student model.Student) (*model.Student, error)
}

type StudentRepository struct {
	db *sqlx.DB
}

func NewStudentRepository(resource *sqlx.DB) IStudentRepository {
	studentRepository := &StudentRepository{db: resource}
	return studentRepository
}

func (s *StudentRepository) GetAll() ([]model.Student, error) {
	students := []model.Student{}
	err := s.db.Select(&students, "SELECT * FROM M_STUDENT")
	if err != nil {
		logger.Log.Error().Err(err).Str("DOMAIN", "Student").Msg("Failed Get All")
		return nil, err
	}
	logger.Log.Debug().Msg(fmt.Sprintf("%v", students))
	return students, nil
}

func (s *StudentRepository) GetOneByName(name string) ([]model.Student, error) {
	students := []model.Student{}
	err := s.db.Select(&students, "SELECT * FROM M_STUDENT WHERE name like '%$1%'", name)
	if err != nil {
		logger.Log.Error().Err(err).Str("DOMAIN", "Student").Msg("Failed Get One By Name")
		return nil, err
	}
	return students, nil
}

func (s *StudentRepository) GetOneById(idCard string) (*model.Student, error) {
	student := model.Student{}
	err := s.db.Get(&student, "SELECT * FROM M_STUDENT WHERE id_card=$1", idCard)
	if err != nil {
		logger.Log.Error().Err(err).Str("DOMAIN", "Student").Msg("Failed Get One By Id")
		return nil, err
	}
	return &student, nil
}

func (s *StudentRepository) CreateOne(student model.Student) (*model.Student, error) {
	lastInsertId := 0
	err := s.db.QueryRow("INSERT INTO M_STUDENT(name,gender,age,join_date,id_card,senior) VALUES($1,$2,$3,$4,$5,$6) RETURNING id", student.Name, student.Gender, student.Age, student.JoinDate, student.IdCard, student.Senior).Scan(&lastInsertId)
	if err != nil {
		logger.Log.Error().Err(err).Str("DOMAIN", "Student").Msg("Failed Create")
		return nil, err
	}
	student.Id = lastInsertId
	return &student, nil
}
