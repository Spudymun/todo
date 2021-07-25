package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/Spudymun/todo"
	"github.com/Spudymun/todo/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "dfafsf8u89yhfudis"
	signingKey = "erewrsfewr345325ewrwer325df"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

// Структура в которой есть поле интерфейса авторизации из репозитария
type AuthService struct {
	repo repository.Authorization
}

// Внедрение зависимотси от repository.Authorization
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// Имплементация метода CreateUser который будет передавать структуру юзера еще на слой ниже(в репозиторий)
func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	// get user from DB
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})

	// Возвращение подписаного токена с помощью метода SignedString на основе ключа подписи, который используеться при расшифровке токена
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	// Привод обьек токена в которм есть поле Claims типа интерфейс к нашей структуре
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	// %x отображает строку в виде шестнадцатеричного исчисления. salt случайный набор символо который добавляеться к хешу
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
