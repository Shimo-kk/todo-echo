package usecase

import (
	"todo/app/domain/user"
	"todo/app/service/schema"
)

type IAuthUsecase interface {
	SignUp(data schema.SignUpModel) error
	SignIn(data schema.SignInModel) (*schema.UserReadModel, error)
}

type authUsecase struct {
	userRepository user.IUserRepositry
}

// 認証ユースケースの作成
func NewAuthUsecase(userRepository user.IUserRepositry) IAuthUsecase {
	return &authUsecase{userRepository}
}

// サインアップ
func (au *authUsecase) SignUp(data schema.SignUpModel) error {
	// E-mailで存在していないか確認する
	if err := au.userRepository.NotExists(data.Email); err != nil {
		return err
	}

	// ユーザーエンティティを作成
	userEntity, err := user.NewEntity(data.Name, data.Email, data.Password)
	if err != nil {
		return err
	}

	// ユーザーの挿入
	if _, err := au.userRepository.Insert(userEntity); err != nil {
		return err
	}

	return nil
}

// サインイン
func (au *authUsecase) SignIn(data schema.SignInModel) (*schema.UserReadModel, error) {
	// E-mailでユーザーを取得
	userEntity, err := au.userRepository.FindByEmail(data.Email)
	if err != nil {
		return nil, err
	}

	// パスワードの検証
	if err := userEntity.VerifyPassword(data.Password); err != nil {
		return nil, err
	}

	// スキーマへ変換
	result := schema.UserReadModel{
		Id:    userEntity.GetId(),
		Name:  userEntity.GetName(),
		Email: userEntity.GetEmail(),
	}

	return &result, nil
}
