package todo

import (
	"encoding/json"
	"errors"
	"os"
	"slices"
	"time"
)

// TODO
/*
	- struct
	- add task
	- complete task
	- update task
	- delete task
	- load items
	- store items
*/

// struct for todos task
type item struct {
	ID			int
	Title		string
	Task		string
	Done		bool
	CreatedAt	time.Time
}

// List of items
type Todos []item

var nextID int

// add task function
func (t *Todos) AddTask(title, task string) {
	todo := item{
		ID: nextID,
		Title: title,
		Task: task,
		Done: false,
		CreatedAt: time.Now(),
	}

	nextID++

	*t = append(*t, todo)
}

// update task to complete task
func (t *Todos) Complete(id int) error {
	err := t.checkID(id)

	if err != nil {
		return err
	}

	ls := *t

	ls[id].Done = true

	return nil
}

// Update tasks
func (t *Todos) Update(id int, title, task string) error{
	err := t.checkID(id)

	if err != nil {
		return err
	}

	ls := *t

	ls[id].Title = title
	ls[id].Task = task
	ls[id].CreatedAt = time.Now()

	return nil
}

// Delete tasks
func (t *Todos) Delete(id int) error {
	err := t.checkID(id)

	ls := *t

	if err != nil {
		return err
	}

	*t = slices.Delete(ls, id, id + 1)

	return nil
}

// Load items
/* 
	- readfile from laptop
	- check for error
	- unmarshal with JSON
	- update nextID

*/
func (t *Todos) Load(filename string) error {
	// read file from laptop
	data, err := os.ReadFile(filename)

	// check if error in ReadFile
	if err != nil {
		return err
	}

	// umarshal
	err = json.Unmarshal(data, t)
	if err != nil {
		return err
	}

	// update NextID
	if len(*t) > 0 {
		maxID := (*t)[0].ID
		for _, todo := range *t {
			if todo.ID > maxID {
				maxID = todo.ID
			}
		}

		nextID = maxID + 1
	}
	return nil
}

// Store items
/* 
	- marshal *t
	- use os library to store
*/
func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

// Implement DRY by making a function to check ID
func (t *Todos) checkID(id int) error {
	ls := *t
	if id >= len(ls) || id < 0 {
		return errors.New("Invalid ID")
	}

	return nil
}

