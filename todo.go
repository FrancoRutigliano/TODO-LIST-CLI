package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
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
	table := simpletable.New()
	var cells [][]*simpletable.Cell
	if len(*t) == 0 {
		table.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: "Mensaje"},
			},
		}

		cells = append(cells, *&[]*simpletable.Cell{
			{Text: "No hay tareas agendadas aún"},
		})

		table.Body = &simpletable.Body{Cells: cells}

	} else {
		table.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: "#"},
				{Align: simpletable.AlignCenter, Text: "NAME"},
				{Align: simpletable.AlignCenter, Text: "STATUS"},
				{Align: simpletable.AlignCenter, Text: "CREATED_AT"},
				{Align: simpletable.AlignCenter, Text: "COMPLETED_AT"},
			},
		}

		for idx, item := range *t {
			idx++

			task := blue(item.Task)
			done := blue("no finalizada")
			if item.Done {
				task = green(fmt.Sprintf("\u2705 %s", item.Task))
				done = green("Finalizada")
			}

			cells = append(cells, *&[]*simpletable.Cell{
				{Text: fmt.Sprintf("%d", idx)},
				{Text: task},
				{Text: done},
				{Text: item.CreatedAt.Format(time.RFC822)},
				{Text: item.CompletedAt.Format(time.RFC822)},
			})
		}

		table.Body = &simpletable.Body{Cells: cells}

		table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("Tenes pendientes %d todos", t.CountPending()))},
		}}

	}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func (t *Todos) CountPending() int {
	total := 0

	for _, item := range *t {
		if !item.Done {
			total++
		}
	}

	return total
}
