package repository

import (
	"fmt"

	"github.com/Spudymun/todo"
	"github.com/jmoiron/sqlx"
)

// Структура которая имплементирует наш интерфейс репозитория
type AuthPostgres struct {
	db *sqlx.DB
}

// Внедрение зависимотси в конструкторе AuthPostgres от sqlx.DB
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// Реализация SQL-запроса на создание юзера
func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	// %s Простой вывод строк. $ - это placeholders для аргументов в методе запроса к БД. RETURNING id возвращает id новой записи
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	// QueryRow() выполняет SQL запрос и возвращает обьек row, котрый хранит в себе инф. о строке в БД. Метод Scan() записывает значение в переменну переданную по ссылке
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// Метод для получения пользователя по его логину и паролю
func (r *AuthPostgres) GetUser(username, password string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	// Метод обьекта БД в который передаем указатель на структуру в которую хотим записать результат выборки
	err := r.db.Get(&user, query, username, password)

	return user, err
}
