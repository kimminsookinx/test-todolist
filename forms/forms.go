package forms

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type TodoItemForm struct{}

type CreateTodoItemForm struct {
	Desc string `form:"description" json:"description" binding:"required,max=1000"`
}

func (f TodoItemForm) Desc(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "len 0 not allowed"
		}
		return errMsg[0]
	case "max":
		return "exceed 1000 char"
	default:
		return "unknown error at forms"
	}
}

func (f TodoItemForm) PostTodoItem(err error) string {
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
