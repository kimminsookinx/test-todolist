package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/kimminsookinx/test-todolist/db"
)

type TodoItem struct {
	ID      int64     `db:"id, primaryket, autoincrement" json:"id"`
	Desc    string    `db:"description" json:"description"`
	Created time.Time `db:"created_at" json:"created_at"`
	Updated time.Time `db:"last_updated_at" json:"last_updated_at"`
	Done    bool      `db:"done_flag" json:"done_flag"`
}
type TodoItemModel struct{}

func (m TodoItemModel) TodoItemList() (items []TodoItem, err error) {
	//"SELECT id, description, created_at, last_updated_at, done_flag FROM todo.item ORDER BY id DESC LIMIT 500"
	_, err = db.GetDB().Select(&items, "SELECT id, description, created_at, last_updated_at, done_flag FROM todo.item ORDER BY id DESC LIMIT 500")
	return items, err
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
