package sessions

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type Session struct {
	UserID string
	Token  string
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type ISessionRepository interface {
	Save(session Session) error
	Delete(id string) error
}

func CreateSession(userID string, secret string, expiry time.Duration) (Session, error) {
	token, err := generateToken(userID, secret, expiry)
	if err != nil {
		return Session{}, err
	}

	return Session{
		UserID: userID,
		Token:  token,
	}, nil
}

func generateToken(userID, secret string, expiry time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": expiry,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *Session) ValidateToken(userID, secret string) error {
	claims := &Claims{}
	parsedToken, err := jwt.ParseWithClaims(s.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return err
	}
	if !parsedToken.Valid {
		return errors.New("invalid token")
	}
	if claims.Subject != userID {
		return errors.New("user ID mismatch")
	}
	return nil
}
