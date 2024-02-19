package task

type ITaskRepository interface {
	Insert(entity *TaskEntity) (*TaskEntity, error)
	FindAll(userId int) (*[]TaskEntity, error)
	FindById(id int) (*TaskEntity, error)
	Update(entity *TaskEntity) (*TaskEntity, error)
	DeleteById(id int) error
}
