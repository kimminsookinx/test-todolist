package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/kimminsookinx/test-todolist/db"
	"github.com/kimminsookinx/test-todolist/forms"
)

// TODO: if more models are added, divide into {modelName}.go
type TodoItem struct {
	ID        int64     `db:"id, primaryket, autoincrement" json:"id"`
	Desc      string    `db:"description" json:"description"`
	Created   time.Time `db:"created_at" json:"created_at"`
	Updated   time.Time `db:"last_updated_at" json:"last_updated_at"`
	Done      bool      `db:"done" json:"done"`
	Deleted   bool      `db:"deleted" json:"deleted"`
	DeletedAt time.Time `db:"deleted_at" json:"deleted_at"`
}
type TodoItemModel struct{}

var queryLimit = "3"

// TODO: validate env var type
func (m TodoItemModel) Init() {
	queryLimit = os.Getenv("TODO_DB_QUERY_MAX_LIMIT")
}
func (m TodoItemModel) SelectTodoItemWhereDeletedIsFalse() (items []TodoItem, err error) {
	_, err = db.GetDB().Select(&items,
		"SELECT id, description, created_at, last_updated_at, done, deleted, deleted_at"+
			"FROM todo.item "+
			"WHERE deleted=false "+
			"ORDER BY id DESC LIMIT "+queryLimit)
	return items, err
}

func (m TodoItemModel) SelectTodoItem() (items []TodoItem, err error) {
	_, err = db.GetDB().Select(&items,
		"SELECT id, description, created_at, last_updated_at, done deleted, deleted_at"+
			"FROM todo.item "+
			"ORDER BY id DESC LIMIT "+queryLimit)
	return items, err
}

func (m TodoItemModel) Post(form forms.CreateTodoItemForm) (todoItemId int64, err error) {
	operation, err := db.GetDB().Exec("INSERT INTO todo.item(description, done) VALUES(\"" + form.Desc + "\", false)")
	if err != nil {
		return 0, err
	}
	todoItemId, err = operation.LastInsertId()
	return todoItemId, err
}

func (m TodoItemModel) UpdateDone(todoItemId int64, form forms.UpdateDoneTodoItemForm) (err error) {
	var boolString string
	if *form.Done {
		boolString = "true"
	} else {
		boolString = "false"
	}
	operation, err := db.GetDB().Exec("UPDATE todo.item SET done=" + boolString + " WHERE id=" + fmt.Sprintf("%d", todoItemId))
	if err != nil {
		return err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return errors.New("updated 0 records")
	}

	return err
}

func (m TodoItemModel) UpdateDesc(todoItemId int64, form forms.UpdateDescTodoItemForm) (err error) {
	operation, err := db.GetDB().Exec("UPDATE todo.item SET description=\"" + form.Desc + "\" WHERE id=" + fmt.Sprintf("%d", todoItemId))
	if err != nil {
		return err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return errors.New("updated 0 records")
	}

	return err
}

// helpers (https://github.com/Massad/gin-boilerplate/blob/master/models/util.go)
type JSONRaw json.RawMessage

func (j *JSONRaw) MarshalJSON() ([]byte, error) {
	return *j, nil
}

// UnmarshalJSON ...
func (j *JSONRaw) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*j = append((*j)[0:0], data...)
	return nil
}

// DataList ....
type DataList struct {
	Data JSONRaw `db:"data" json:"data"`
	Meta JSONRaw `db:"meta" json:"meta"`
}
