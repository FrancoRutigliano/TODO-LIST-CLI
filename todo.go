package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []Item

func (t *Todos) Add(task string) {
	todo := Item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	ls := *t

	if index <= 0 || index > len(ls) {
		return errors.New("index Invalido")
	}

	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

func (t *Todos) Delete(index int) error {

	ls := *t

	if index <= 0 || index > len(ls) {
		return errors.New("index invalido")
	}

	*t = append(ls[:index-1], ls[index:]...)

	return nil
}

func (t *Todos) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todos) Store(filename string) error {
	data, err := json.MarshalIndent(t, "", "")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)

}

func (t *Todos) Print() {
	for i, item := range *t {
		i++
		fmt.Printf("%d - %s\n", i, item.Task)
	}
}
