package user

import (
	"time"
	"todo/app/core"

	"golang.org/x/crypto/bcrypt"
)

type UserEntity struct {
	id        int
	createdAt time.Time
	updatedAt time.Time
	name      string
	email     string
	password  string
}

// エンティティの作成
func NewEntity(name string, email string, password string) (*UserEntity, error) {
	// バリデーション
	if err := validateName(name); err != nil {
		return nil, err
	}
	if err := validateEmail(email); err != nil {
		return nil, err
	}
	if err := validatePassword(password); err != nil {
		return nil, err
	}

	// パスワードの暗号化
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, core.NewError(core.SystemError, "パスワードの暗号化に失敗しました。->"+err.Error())
	}

	return &UserEntity{name: name, email: email, password: string(hashed)}, nil
}

// パスワードの検証
func (ue *UserEntity) VerifyPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(ue.password), []byte(password)); err != nil {
		return core.NewError(core.BadRequestError, "パスワードが正しくありません。")
	}

	return nil
}

// ユーザー名の変更
func (ue *UserEntity) ChangeName(name string) error {
	// バリデーション
	if err := validateName(name); err != nil {
		return err
	}

	ue.name = name
	return nil
}

// パスワードの変更
func (ue *UserEntity) ChangePassword(password string) error {
	// バリデーション
	if err := validatePassword(password); err != nil {
		return err
	}

	// パスワードの暗号化
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return core.NewError(core.SystemError, "パスワードの暗号化に失敗しました。->"+err.Error())
	}

	ue.password = string(hashed)
	return nil
}

func (ue *UserEntity) GetId() int {
	return ue.id
}

func (ue *UserEntity) GetCreatedAt() time.Time {
	return ue.createdAt
}

func (ue *UserEntity) GetUpdatedAt() time.Time {
	return ue.updatedAt
}

func (ue *UserEntity) GetName() string {
	return ue.name
}

func (ue *UserEntity) GetEmail() string {
	return ue.email
}

func (ue *UserEntity) GetPassword() string {
	return ue.password
}
