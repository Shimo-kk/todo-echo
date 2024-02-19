package task

import (
	"todo/app/core"
	"unicode/utf8"
)

// タイトルのバリデーション
func validateTitle(value string) error {
	if value == "" {
		return core.NewError(core.ValidationError, "タイトルは空の値を入力することはできません。")
	}

	if utf8.RuneCountInString(value) > 50 {
		return core.NewError(core.ValidationError, "タイトルは50文字より大きい値を入力できません。")
	}

	return nil
}
