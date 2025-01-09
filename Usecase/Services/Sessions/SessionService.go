package Sessions

import (
	sessions "auth-server/Domain/Models/Sessions"
	"os"
	"time"
)

type ISessionService interface {
	CreateSession(userId string) (sessions.Session, sessions.Session, error)
	DeleteSession(userId string) error
	RefreshSession(userId string, token string) (sessions.Session, sessions.Session, error)
	ValidateSession(userId string, token string) error
}

type SessionService struct {
	sessionRepository sessions.ISessionRepository
}

var (
	secretKey          = os.Getenv("JWT_SECRET")         // JWT署名用シークレットキー
	refreshSecretKey   = os.Getenv("JWT_REFRESH_SECRET") // リフレッシュトークン用のシークレットキー
	accessTokenExpiry  = time.Minute * 15                // アクセストークンの有効期限(15分)
	refreshTokenExpiry = time.Hour * 24 * 7              // リフレッシュトークンの有効期限(7日)
)

func NewSessionService(sessionRepository sessions.ISessionRepository) ISessionService {
	return &SessionService{sessionRepository}
}

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

func (ss *SessionService) RefreshSession(userId string, token string) (sessions.Session, sessions.Session, error) {
	session := sessions.Session{Token: token}
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

func (ss *SessionService) ValidateSession(userId string, token string) error {
	session := sessions.Session{Token: token}
	err := session.ValidateToken(userId, secretKey)
	if err != nil {
		return err
	}
	return nil
}
