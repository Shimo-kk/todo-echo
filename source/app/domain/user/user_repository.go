package user

type IUserRepositry interface {
	Insert(entity *UserEntity) (*UserEntity, error)
	NotExists(email string) error
	FindById(id int) (*UserEntity, error)
	FindByEmail(email string) (*UserEntity, error)
	Update(entity *UserEntity) (*UserEntity, error)
	DeleteById(id int) error
}
