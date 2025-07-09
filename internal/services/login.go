package services

import (
	"context"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"
	"time"

	"github.com/pkg/errors"

	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	UserRepository interfaces.IUserRepository
}

func (s *LoginService) Login(ctx context.Context, req models.LoginRequest) (models.LoginResponse, error) {
	var (
		res models.LoginResponse
		now = time.Now()
	)
	userDetail, err := s.UserRepository.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return res, errors.Wrap(err, "failed to get user by username")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(req.Password)); err != nil {
		return res, errors.Wrap(err, "failed to compare password")
	}

	token, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "token", now)
	if err != nil {
		return res, errors.Wrap(err, "failed to generate token")
	}

	refreshToken, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "refresh_token", now)
	if err != nil {
		return res, errors.Wrap(err, "failed to generate refresh token")
	}

	userSession := &models.UserSession{
		UserID:              userDetail.ID,
		Token:               token,
		RefreshToken:        refreshToken,
		TokenExpired:        now.Add(helpers.MapTypeToken["token"]),
		RefreshTokenExpired: now.Add(helpers.MapTypeToken["refresh_token"]),
	}

	err = s.UserRepository.InsertNewUserSession(ctx, userSession)
	if err != nil {
		return res, errors.Wrap(err, "failed to insert user session")
	}

	res.UserID = userDetail.ID
	res.Username = userDetail.Username
	res.FullName = userDetail.FullName
	res.Email = userDetail.Email
	res.Token = token
	res.RefreshToken = refreshToken

	return res, nil
}
