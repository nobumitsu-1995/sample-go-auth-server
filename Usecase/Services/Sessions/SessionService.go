package Sessions

import (
	sessions "auth-server/Domain/Models/Sessions"
	users "auth-server/Domain/Models/Users"
	"time"
)

type ISessionService interface {
	CreateSession(user users.User) (sessions.Session, sessions.Session, error)
	DeleteSession(user users.User) error
	ValidateToken(userID, token, secret string) error
}

type SessionService struct {
	sessionRepository sessions.ISessionRepository
}

var (
	secretKey          = "my_secret_key"         // JWT署名用シークレットキー
	refreshSecretKey   = "my_refresh_secret_key" // リフレッシュトークン用のシークレットキー
	accessTokenExpiry  = time.Minute * 15        // アクセストークンの有効期限(15分)
	refreshTokenExpiry = time.Hour * 24 * 7      // リフレッシュトークンの有効期限(7日)
)

func (ss *SessionService) CreateSession(userId string) (sessions.Session, sessions.Session, error) {
	accessToken, err := sessions.CreateSession(userId, secretKey, accessTokenExpiry)
	if err != nil {
		return sessions.Session{}, sessions.Session{}, err
	}

	refreshToken, err := sessions.CreateSession(userId, refreshSecretKey, refreshTokenExpiry)
	if err != nil {
		return sessions.Session{}, sessions.Session{}, err
	}

	err = ss.sessionRepository.Save(refreshToken)
	if err != nil {
		return sessions.Session{}, sessions.Session{}, err
	}
	return accessToken, refreshToken, nil
}

func (ss *SessionService) DeleteSession(userId string) error {
	err := ss.sessionRepository.Delete(userId)
	if err != nil {
		return err
	}
	return nil
}

func (ss *SessionService) RefreshSession(userId string, session sessions.Session, secret string) (sessions.Session, sessions.Session, error) {
	err := session.ValidateToken(userId, refreshSecretKey)
	if err != nil {
		return sessions.Session{}, sessions.Session{}, err
	}

	err = ss.sessionRepository.Delete(userId)
	if err != nil {
		return sessions.Session{}, sessions.Session{}, err
	}

	accessToken, refreshToken, err := ss.CreateSession(userId)
	if err != nil {
		return sessions.Session{}, sessions.Session{}, err
	}
	return accessToken, refreshToken, nil
}
