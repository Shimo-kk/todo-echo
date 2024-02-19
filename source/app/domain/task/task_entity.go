package task

import "time"

type TaskEntity struct {
	id        int
	createdAt time.Time
	updatedAt time.Time
	userId    int
	title     string
	doneFlag  bool
}

// エンティティの作成
func NewEntity(userId int, title string) (*TaskEntity, error) {
	// バリデーション
	if err := validateTitle(title); err != nil {
		return nil, err
	}

	return &TaskEntity{userId: userId, title: title, doneFlag: false}, nil
}

// タイトルの変更
func (ue *TaskEntity) ChangeTitle(title string) error {
	// バリデーション
	if err := validateTitle(title); err != nil {
		return err
	}

	ue.title = title
	return nil
}

// タスクの完了
func (ue *TaskEntity) Done() {
	ue.doneFlag = true
}

func (ue *TaskEntity) GetId() int {
	return ue.id
}

func (ue *TaskEntity) GetCreatedAt() time.Time {
	return ue.createdAt
}

func (ue *TaskEntity) GetUpdatedAt() time.Time {
	return ue.updatedAt
}

func (ue *TaskEntity) GetUserId() int {
	return ue.userId
}

func (ue *TaskEntity) GetTitle() string {
	return ue.title
}

func (ue *TaskEntity) GetDoneFlag() bool {
	return ue.doneFlag
}
