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
	DeleteUser(ctx context.Context, user *domain.User) error
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

func (i *userInteractor) CreateUser(ctx context.Context, formUser *domain.User) (*domain.TokenPair, error) {

	// recodeない場合もエラーになって帰ってくるので、とりあえず無視
	emailUsedUser, err := i.userRepository.GetUserByEmail(ctx, formUser.Email)
	if emailUsedUser != nil {
		// すでに該当するemailを利用しているユーザーがいる場合、error返す
		return nil, errors.EmailAlreadyUsed
	}

	encryptedPass, err := genEncryptedPass(formUser.Password)
	if err != nil {
		return nil, err
	}
	formUser.Password = encryptedPass

	if err := i.userRepository.CreateUser(ctx, formUser); err != nil {
		return nil, err
	}

	return genTokenPair(strconv.FormatInt(formUser.ID, 10))

}

func (i *userInteractor) UpdateUser(ctx context.Context, formUser *domain.User) (*domain.User, *domain.User, error) {
	// 更新前ユーザー取得
	bUser, err := i.userRepository.GetUserByID(ctx, formUser.ID)
	if err != nil {
		return nil, nil, err
	}

	// emailに更新がある場合、更新後のemailアドレスを使っているユーザーがいないか確認
	if bUser.Email != formUser.Email {
		emailUsedUser, _ := i.userRepository.GetUserByEmail(ctx, formUser.Email)
		if emailUsedUser != nil && formUser.ID != emailUsedUser.ID {
			// すでに該当するemailを利用しているユーザーがいて、IDが別である場合error返す
			return nil, nil, errors.EmailAlreadyUsed
		}
	}

	// password 暗号化
	encryptedPass, err := genEncryptedPass(formUser.Password)
	if err != nil {
		return nil, nil, err
	}
	formUser.Password = encryptedPass

	// update
	aUser, err := i.userRepository.UpdateUser(ctx, formUser)
	if err != nil {
		return nil, nil, err
	}
	return bUser, aUser, nil
}

func (i *userInteractor) DeleteUser(ctx context.Context, formUser *domain.User) (bool, error) {
	// 更新前ユーザー取得
	bUser, err := i.userRepository.GetUserByID(ctx, formUser.ID)
	if err != nil {
		return false, err
	}

	if err := compareHashAndPass(bUser.Password, formUser.Password); err != nil {
		return false, err
	}

	// 削除
	result, err := i.userRepository.DeleteUser(ctx, formUser.ID)
	if err != nil {
		return false, err
	}

	return result, nil
}
