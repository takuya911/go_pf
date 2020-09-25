package interactor

import (
	"context"
	"strconv"

	"github.com/takuya911/go_pf/services/user/domain"
	"github.com/takuya911/go_pf/services/user/errors"
	"github.com/takuya911/go_pf/services/user/usecase/repository"
)

type userInteractor struct {
	userRepository repository.UserRepository
}

// NewUserInteractor function
func NewUserInteractor(u repository.UserRepository) *userInteractor {
	return &userInteractor{u}
}

// UserUsecase interface
type UserUsecase interface {
	GetUserByID(ctx context.Context, userID int64) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) (*domain.TokenPair, error)
	UpdateUser(ctx context.Context, user *domain.User) error
}

func (i *userInteractor) GetUserByID(ctx context.Context, userID int64) (*domain.User, error) {
	result, err := i.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (i *userInteractor) Login(ctx context.Context, email string, password string) (*domain.User, *domain.TokenPair, error) {
	user, err := i.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, nil, err
	}

	if err := compareHashAndPass(user.Password, password); err != nil {
		return nil, nil, err
	}

	tokenPair, err := genTokenPair(strconv.FormatInt(user.ID, 10))
	if err != nil {
		return nil, nil, err
	}

	return user, tokenPair, nil
}

func (i *userInteractor) CreateUser(ctx context.Context, user *domain.User) (*domain.TokenPair, error) {

	emailUsedUser, err := i.userRepository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		// recodeない場合もエラーになって帰ってくるので、とりあえず無視
	}
	if emailUsedUser != nil && user.ID != emailUsedUser.ID {
		// email利用しているユーザーがいて、IDが別である場合はerr
		return nil, errors.UserAlreadyExists
	}

	encryptedPass, err := genEncryptedPass(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = encryptedPass

	if err := i.userRepository.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return genTokenPair(strconv.FormatInt(user.ID, 10))

}

func (i *userInteractor) UpdateUser(ctx context.Context, formUser *domain.User) (*domain.User, *domain.User, error) {

	// IDからユーザー情報取得
	existUser, err := i.userRepository.GetUserByID(ctx, formUser.ID)
	if err != nil {
		return nil, nil, err
	}
	// emailを変更する場合は、変更後のEmailを他の人が使っていないか確認
	if formUser.Email != existUser.Email {
		user, err := i.userRepository.GetUserByEmail(ctx, formUser.Email)
		if err != nil {
			if formUser.ID != user.ID {
				return nil, nil, errors.UserAlreadyExists
			}
		}
	}

	// password 暗号化
	encryptedPass, err := genEncryptedPass(formUser.Password)
	if err != nil {
		return nil, nil, err
	}
	formUser.Password = encryptedPass

	// update
	bUser, aUser, err := i.userRepository.UpdateUser(ctx, formUser)
	if err != nil {
		return nil, nil, err
	}
	return bUser, aUser, nil
}
