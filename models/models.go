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
// https://willnorris.com/2014/go-rest-apis-and-pointers/
type TodoItem struct {
	ID        int64      `db:"id, primaryket, autoincrement" json:"id"`
	Desc      string     `db:"description" json:"description"`
	Created   time.Time  `db:"created_at" json:"created_at"`
	Updated   time.Time  `db:"last_updated_at" json:"last_updated_at"`
	Done      bool       `db:"done" json:"done"`
	Deleted   bool       `db:"deleted" json:"deleted"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}
type TodoItemModel struct{}

var queryLimit = "3"

// TODO: validate env var type
func (m TodoItemModel) Init() {
	queryLimit = os.Getenv("TODO_DB_QUERY_MAX_LIMIT")
}
func (m TodoItemModel) SelectTodoItemWhereDeletedIsFalse() (items []TodoItem, err error) {
	_, err = db.GetDB().Select(&items,
		"SELECT id, description, created_at, last_updated_at, done, deleted, deleted_at "+
			"FROM todo.item "+
			"WHERE deleted=false "+
			"ORDER BY id DESC LIMIT "+queryLimit)
	return items, err
}

func (m TodoItemModel) SelectTodoItem(queryValues map[string][]string) (items []TodoItem, err error) {
	query, err := decodeQuery(queryValues)
	if err != nil {
		return nil, err
	}
	_, err = db.GetDB().Select(&items,
		"SELECT id, description, created_at, last_updated_at, done deleted, deleted_at "+
			"FROM todo.item "+
			query+
			"ORDER BY id DESC LIMIT "+queryLimit)
	return items, err
}

func (m TodoItemModel) InsertTodoItem(form forms.CreateTodoItemForm) (todoItemId int64, err error) {
	operation, err := db.GetDB().Exec(
		"INSERT INTO todo.item(description, done) " +
			"VALUES(\"" + form.Desc + "\", false)")
	if err != nil {
		return 0, err
	}
	todoItemId, err = operation.LastInsertId()
	return todoItemId, err
}

func (m TodoItemModel) UpdateTodoItemSetDoneById(todoItemId int64, form forms.UpdateDoneTodoItemForm) (err error) {
	operation, err := db.GetDB().Exec(
		"UPDATE todo.item " +
			"SET done=" + strconv.FormatBool(*form.Done) +
			" WHERE id=" + fmt.Sprintf("%d", todoItemId))
	if err != nil {
		return err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return errors.New("updated 0 records")
	}

	return err
}

func (m TodoItemModel) UpdateTodoItemSetDescById(todoItemId int64, form forms.UpdateDescTodoItemForm) (err error) {
	operation, err := db.GetDB().Exec(
		"UPDATE todo.item " +
			"SET description=\"" + form.Desc + "\" " +
			"WHERE id=" + fmt.Sprintf("%d", todoItemId))
	if err != nil {
		return err
	}

	success, _ := operation.RowsAffected()
	if success == 0 {
		return errors.New("updated 0 records")
	}

	return err
}

func (m TodoItemModel) UpdateTodoItemSetDeletedIsFalseById(todoItemId int64) (success bool, err error) {
	tString := time.Now().Format(time.RFC3339)
	operation, err := db.GetDB().Exec(
		"UPDATE todo.item " +
			"SET deleted=true, deleted_at=\"" + tString + "\" " +
			"WHERE id=" + fmt.Sprintf("%d", todoItemId))
	if err != nil {
		return false, err
	}
	sqlSuccess, _ := operation.RowsAffected()
	if sqlSuccess == 0 {
		return false, errors.New("updated 0 records")
	}
	return true, nil
}

// NOTE: https://stackoverflow.com/questions/1676551/best-way-to-test-if-a-row-exists-in-a-mysql-table
func (m TodoItemModel) CheckRowExistenceById(todoItemId int64) (result bool, err error) {
	err = db.GetDB().SelectOne(&result, "SELECT EXISTS(SELECT 1 FROM todo.item WHERE id="+fmt.Sprint(todoItemId)+" LIMIT 1) AS 'exists'")
	if err != nil {
		return false, err
	}
	return result, nil
}

func (m TodoItemModel) CheckRowExistenceByIdAndDeleted(todoItemId int64, deleted bool) (result bool, err error) {
	err = db.GetDB().SelectOne(&result, "SELECT EXISTS(SELECT 1 FROM todo.item WHERE id="+fmt.Sprint(todoItemId)+" AND deleted="+strconv.FormatBool(deleted)+" LIMIT 1) AS 'exists'")
	if err != nil {
		return false, err
	}
	return result, nil
}

//custom helpers

func decodeQuery(querys map[string][]string) (result string, err error) {
	var interResultString string

	for key, val := range querys {
		switch key {
		case "showDeleted":
			//oneof validation cannot contain spaces (https://github.com/go-playground/validator/issues/525)
			if val[0] == "" {
				interResultString = interResultString + "deleted=false "
				break
			}
			show, parseErr := strconv.ParseBool(val[0])
			if parseErr != nil {
				return "", parseErr
			}
			if !show {
				interResultString = interResultString + "deleted=false "
			} else {
				interResultString = interResultString + ""
			}
		}
	}
	if len(interResultString) > 0 {
		interResultString = "WHERE " + interResultString
	}
	return interResultString, nil
}

// TODO: check if used, if not used delete helpers
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
