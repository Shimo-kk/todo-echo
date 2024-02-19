package repository

import (
	"todo/app/core"
	"todo/app/domain/task"

	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

// タスクリポジトリの作成
func NewTaskRepository(db *gorm.DB) task.ITaskRepository {
	return &taskRepository{db: db}
}

// 挿入
func (tr *taskRepository) Insert(entity *task.TaskEntity) (*task.TaskEntity, error) {
	dto := task.ToDtoFromEntity(entity)
	if err := tr.db.Create(dto).Error; err != nil {
		return nil, core.NewError(core.SystemError, "タスクの作成に失敗しました。->"+err.Error())
	}
	result := dto.ToEntity()
	return result, nil
}

// 全件取得
func (tr *taskRepository) FindAll(userId int) (*[]task.TaskEntity, error) {
	dtos := []task.Task{}
	if err := tr.db.Where(&task.Task{UserId: userId}).Find(&dtos).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, core.NewError(core.NotFoundError, "タスクが存在しません。")
		} else {
			return nil, core.NewError(core.SystemError, "タスクの取得に失敗しました。->"+err.Error())
		}
	}
	result := []task.TaskEntity{}
	for _, dto := range dtos {
		result = append(result, *dto.ToEntity())
	}
	return &result, nil
}

// IDで取得
func (tr *taskRepository) FindById(id int) (*task.TaskEntity, error) {
	dto := task.Task{}
	if err := tr.db.Where(&task.Task{Id: id}).First(&dto).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, core.NewError(core.NotFoundError, "タスクが存在しません。")
		} else {
			return nil, core.NewError(core.SystemError, "タスクの取得に失敗しました。->"+err.Error())
		}
	}
	result := dto.ToEntity()
	return result, nil
}

// 更新
func (tr *taskRepository) Update(entity *task.TaskEntity) (*task.TaskEntity, error) {
	dto := task.ToDtoFromEntity(entity)
	if err := tr.db.Save(dto).Error; err != nil {
		return nil, core.NewError(core.SystemError, "タスクの更新に失敗しました。->"+err.Error())
	}
	result := dto.ToEntity()
	return result, nil
}

// 削除
func (tr *taskRepository) DeleteById(id int) error {
	if err := tr.db.Delete(&task.Task{Id: id}).Error; err != nil {
		return core.NewError(core.SystemError, "タスクの削除に失敗しました。->"+err.Error())
	}
	return nil
}
