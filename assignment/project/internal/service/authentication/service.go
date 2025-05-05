package authentication

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	ct "golang-project/internal/contract"
	"golang-project/internal/model"
	repo "golang-project/internal/repository"
	svc "golang-project/internal/service"
	"golang-project/static"
	"golang-project/util/hashing"
)

// service represents the implementation of service.Authentication
type service struct {
	userRepo repo.User
	hash     hashing.Algorithm
}

// NewService returns a new implementation of service.Authentication
func NewService(userRepo repo.User, hash hashing.Algorithm) svc.Authentication {
	return &service{
		userRepo: userRepo,
		hash:     hash,
	}
}

// SignIn executes the user authentication logic
func (s *service) SignIn(r *ct.SignInRequest) (*ct.SignInResponse, error) {
	user, err := s.userRepo.ReadByEmail(r.Email)
	if err != nil {
		return nil, err
	}

	err = s.hash.Compare([]byte(user.Password), []byte(r.Password))
	if err != nil {
		return nil, err
	}

	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	return prepareSignInResponse(user, token), nil
}

// generateToken returns the JWT token based on the information from model.User
func (s *service) generateToken(user *model.User) (string, error) {
	secret := []byte(viper.GetString(static.EnvAuthSecret))
	customClaim := &ct.CustomClaim{
		StandardClaims: jwt.StandardClaims{
			Audience:  viper.GetString(static.EnvAuthAudience),
			ExpiresAt: time.Now().Unix() + viper.GetInt64(static.EnvAuthLifeTime),
			Id:        uuid.NewString(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    viper.GetString(static.EnvAuthIssuer),
			NotBefore: time.Now().Unix(),
			Subject:   viper.GetString(static.EnvAuthSubject),
		},
		UserID:    user.ID,
		UserEmail: user.Email,
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, customClaim).SignedString(secret)
}

// SignUp handles the user registration process
func (s *service) SignUp(r *ct.SignUpRequest) (*ct.SignUpResponse, error) {
	// Check if email already exists
	existingUser, err := s.userRepo.ReadByEmail(r.Email)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, static.ErrCheckEmailFailed
		}
	}

	if existingUser != nil {
		return nil, static.ErrEmailAlreadyExists
	}

	// Hash the password
	hashedPassword, err := s.hash.Generate([]byte(r.Password))
	if err != nil {
		return nil, static.ErrPasswordHashingFailed
	}

	// Create new user
	user := &model.User{
		Email:      r.Email,
		Password:   string(hashedPassword),
		FirstName:  r.FirstName,
		LastName:   r.LastName,
		Pseudonym:  r.Email,
		IsVerified: false,
	}

	// Save to database
	user, err = s.userRepo.Insert(user)
	if err != nil {
		return nil, static.ErrSaveUserFailed
	}

	return &ct.SignUpResponse{
		User: user,
	}, nil
}
