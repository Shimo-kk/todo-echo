package user

import "time"

type User struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Email     string
	Password  string
}

// エンティティをDTOに変換
func ToDtoFromEntity(ue *UserEntity) *User {
	return &User{
		Id:        ue.id,
		CreatedAt: ue.createdAt,
		UpdatedAt: ue.updatedAt,
		Name:      ue.name,
		Email:     ue.email,
		Password:  ue.password,
	}
}

// DTOをエンティティに変換
func (ud *User) ToEntity() *UserEntity {
	return &UserEntity{
		id:        ud.Id,
		createdAt: ud.CreatedAt,
		updatedAt: ud.UpdatedAt,
		name:      ud.Name,
		email:     ud.Email,
		password:  ud.Password,
	}
}
