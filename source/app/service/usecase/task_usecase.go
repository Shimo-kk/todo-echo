package usecase

import (
	"todo/app/domain/task"
	"todo/app/service/schema"
)

type ITaskUsecase interface {
	CreateTask(userId int, data schema.TaskCreateModel) (*schema.TaskReadModel, error)
	GetTask(id int) (*schema.TaskReadModel, error)
	GetAllTask(userId int) (*[]schema.TaskReadModel, error)
	UpdateTask(data schema.TaskUpdateModel) (*schema.TaskReadModel, error)
	DoneTask(id int) (*schema.TaskReadModel, error)
	DeleteTask(id int) error
}

type taskUsecase struct {
	taskRepository task.ITaskRepository
}

// タスクユースケースの作成
func NewTaskUsecase(taskRepository task.ITaskRepository) ITaskUsecase {
	return &taskUsecase{taskRepository: taskRepository}
}

// タスクの作成
func (tu *taskUsecase) CreateTask(userId int, data schema.TaskCreateModel) (*schema.TaskReadModel, error) {
	//タスクエンティティの作成
	taskEntity, err := task.NewEntity(userId, data.Title)
	if err != nil {
		return nil, err
	}

	// タスクの挿入
	inserted, err := tu.taskRepository.Insert(taskEntity)
	if err != nil {
		return nil, err
	}

	// スキーマへ変換
	result := schema.TaskReadModel{
		Id:       inserted.GetId(),
		UserId:   inserted.GetUserId(),
		Title:    inserted.GetTitle(),
		DoneFlag: inserted.GetDoneFlag(),
	}

	return &result, nil
}

// タスクの取得
func (tu *taskUsecase) GetTask(id int) (*schema.TaskReadModel, error) {
	// タスクを取得
	taskEntity, err := tu.taskRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	// スキーマへ変換
	result := schema.TaskReadModel{
		Id:       taskEntity.GetId(),
		UserId:   taskEntity.GetUserId(),
		Title:    taskEntity.GetTitle(),
		DoneFlag: taskEntity.GetDoneFlag(),
	}

	return &result, nil
}

// タスクを全件取得
func (tu *taskUsecase) GetAllTask(userId int) (*[]schema.TaskReadModel, error) {
	// タスクを全件取得
	taskEntirys, err := tu.taskRepository.FindAll(userId)
	if err != nil {
		return nil, err
	}

	// スキーマへ変換
	result := []schema.TaskReadModel{}
	for _, entity := range *taskEntirys {
		model := schema.TaskReadModel{
			Id:       entity.GetId(),
			UserId:   entity.GetUserId(),
			Title:    entity.GetTitle(),
			DoneFlag: entity.GetDoneFlag(),
		}
		result = append(result, model)
	}

	return &result, nil
}

// タスクを更新
func (tu *taskUsecase) UpdateTask(data schema.TaskUpdateModel) (*schema.TaskReadModel, error) {
	// タスクを取得
	taskEntiry, err := tu.taskRepository.FindById(data.Id)
	if err != nil {
		return nil, err
	}

	// タイトルを更新
	if err := taskEntiry.ChangeTitle(data.Title); err != nil {
		return nil, err
	}

	// タスクを更新
	updated, err := tu.taskRepository.Update(taskEntiry)
	if err != nil {
		return nil, err
	}

	// スキーマへ変換
	result := schema.TaskReadModel{
		Id:       updated.GetId(),
		UserId:   updated.GetUserId(),
		Title:    updated.GetTitle(),
		DoneFlag: updated.GetDoneFlag(),
	}
	return &result, nil
}

// タスクの完了
func (tu *taskUsecase) DoneTask(id int) (*schema.TaskReadModel, error) {
	// タスクを取得
	taskEntiry, err := tu.taskRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	// タスクを完了
	taskEntiry.Done()

	// タスクを更新
	updated, err := tu.taskRepository.Update(taskEntiry)
	if err != nil {
		return nil, err
	}

	// スキーマへ変換
	result := schema.TaskReadModel{
		Id:       updated.GetId(),
		UserId:   updated.GetUserId(),
		Title:    updated.GetTitle(),
		DoneFlag: updated.GetDoneFlag(),
	}
	return &result, nil
}

// タスクの削除
func (tu *taskUsecase) DeleteTask(id int) error {
	// タスクの削除
	if err := tu.taskRepository.DeleteById(id); err != nil {
		return nil
	}

	return nil
}
