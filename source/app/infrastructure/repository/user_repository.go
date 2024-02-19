package repository

import (
	"todo/app/core"
	"todo/app/domain/user"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// ユーザーリポジトリの作成
func NewUserRepository(db *gorm.DB) user.IUserRepositry {
	return &userRepository{db: db}
}

// 挿入
func (ur *userRepository) Insert(entity *user.UserEntity) (*user.UserEntity, error) {
	dto := user.ToDtoFromEntity(entity)
	if err := ur.db.Create(dto).Error; err != nil {
		return nil, core.NewError(core.SystemError, "ユーザーの作成に失敗しました。->"+err.Error())
	}
	result := dto.ToEntity()
	return result, nil
}

// 存在していないか確認
func (ur *userRepository) NotExists(email string) error {
	dto := user.User{}
	if err := ur.db.Where(&user.User{Email: email}).First(&dto).Error; err != nil && err == gorm.ErrRecordNotFound {
		return core.NewError(core.AlreadyExistsError, "E-mailアドレスがすでに存在しています。")
	}

	return nil
}

// IDで取得
func (ur *userRepository) FindById(id int) (*user.UserEntity, error) {
	dto := user.User{}
	if err := ur.db.Where(&user.User{Id: id}).First(&dto).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, core.NewError(core.NotFoundError, "ユーザーが存在しません。")
		} else {
			return nil, core.NewError(core.SystemError, "ユーザーの取得に失敗しました。->"+err.Error())
		}
	}
	result := dto.ToEntity()
	return result, nil
}

// Emailで取得
func (ur *userRepository) FindByEmail(email string) (*user.UserEntity, error) {
	dto := user.User{}
	if err := ur.db.Where(&user.User{Email: email}).First(&dto).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, core.NewError(core.NotFoundError, "ユーザーが存在しません。")
		} else {
			return nil, core.NewError(core.SystemError, "ユーザーの取得に失敗しました。->"+err.Error())
		}
	}
	result := dto.ToEntity()
	return result, nil
}

// 更新
func (ur *userRepository) Update(entity *user.UserEntity) (*user.UserEntity, error) {
	dto := user.ToDtoFromEntity(entity)
	if err := ur.db.Save(dto).Error; err != nil {
		return nil, core.NewError(core.SystemError, "ユーザーの更新に失敗しました。->"+err.Error())
	}
	result := dto.ToEntity()
	return result, nil
}

// 削除
func (ur *userRepository) DeleteById(id int) error {
	if err := ur.db.Delete(&user.User{Id: id}).Error; err != nil {
		return core.NewError(core.SystemError, "ユーザーの削除に失敗しました。->"+err.Error())
	}
	return nil
}
