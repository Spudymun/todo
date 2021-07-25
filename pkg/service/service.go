package service

import (
	"github.com/Spudymun/todo"
	"github.com/Spudymun/todo/pkg/repository"
)

// Интерфейсы для сущностей
type Authorization interface {
	// Возвращает id пользователя созданного в БД юзера и ошибку
	CreateUser(user todo.User) (int, error)
	// Возвращает стрингу token и ошибку
	GenerateToken(username, password string) (string, error)
	// Принимает токен и возращает id пользователя при успешном парсинге
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

// Слой сервисов обращаеться к БД, поэтому вводим зависимость от указателя на структуру репозитория
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
