package forms

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type TodoItemForm struct{}

type CreateTodoItemForm struct {
	Desc string `form:"description" json:"description" binding:"required,max=50"`
}

type UpdateDescTodoItemForm struct {
	Desc string `form:"description" json:"description" binding:"required,max=50"`
}

//NOTE: boolean binding https://github.com/gin-gonic/gin/issues/814#issuecomment-294636138

type UpdateDoneTodoItemForm struct {
	Done *bool `form:"done" json:"done" binding:"required"`
}

type GetQueryFormString struct {
	ShowDeleted string `form:"showDeleted" binding:"omitempty,oneof=true false"`
}

func (f TodoItemForm) Desc(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "len 0 not allowed"
		}
		return errMsg[0]
	case "max":
		return "exceed 50 char"
	default:
		return "unknown error at forms"
	}
}

func (f TodoItemForm) CheckDesc(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Desc" {
				return f.Desc(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}

func (f TodoItemForm) CheckDoneFlag(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later - unmarshall"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "done" {
				return "boolean required"
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later - update done"
}
