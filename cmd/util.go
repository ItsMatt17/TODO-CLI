package cmd

import (
	// "encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
	// "strconv"
	// "time"
)

func VerifyDatabase() bool {
	f, err := os.Open("tasks.csv")
	if err != nil {
		return false
	}
	defer f.Close()
	return true
}

func CreateDatabase() error {
	f, err := os.Create("tasks.csv")
	if err != nil {
		return fmt.Errorf("lol there's an error")
	}
	defer f.Close()

	return nil

}

func fetchNextId() int {
	return len(MustFetchTasks()) + 1
}

func FetchTaskById(id uint8) (*Todo, error) {
	tasks := MustFetchTasks()

	for _, task := range tasks {
		if id == task.Id {
			return task, nil
		}

	}

	return &Todo{}, fmt.Errorf("cannot fetch task")

}

func CreateNewTask(task *Todo) {

	tasks := MustFetchTasks()

	tasks = append(tasks, task)

	file, err := os.OpenFile("tasks.csv", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	Must(err)
	defer file.Close()

	err = gocsv.MarshalFile(&tasks, file)
	if err != nil {
		fmt.Println("Could not output data to file")
	}

}

func MustFetchTasks() []*Todo {
	// records := mustFetchCSV()
	var tasks []*Todo
	file, err := os.Open("tasks.csv")
	Must(err)

	defer file.Close()

	err = gocsv.UnmarshalFile(file, &tasks)
	Must(err)

	return tasks
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func MustValue[T interface{}](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
