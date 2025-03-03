package cmd

import (
	// "encoding/csv"
	"fmt"
	// "io"
	"os"

	"github.com/gocarina/gocsv"
	// "strconv"
	// "time"
)

const (
	PATH = "tasks.csv"
)

func VerifyDatabase() bool {
	f, err := os.Open(PATH)
	if err != nil {
		return false
	}
	defer f.Close()
	return true
}

func CreateDatabase() error {
	f, err := os.Create(PATH)
	if err != nil {
		return fmt.Errorf("lol there's an error")
	}
	defer f.Close()

	return nil

}

func fetchNextId() int {
	return len(MustFetchTasks()) + 1
}

func FetchTaskById(id uint8) (*Todo, int, error) {
	tasks := MustFetchTasks()

	for idx, task := range tasks {
		if id == task.Id {
			return task, idx, nil
		}

	}

	return &Todo{}, 0, fmt.Errorf(("the task by such ID does not exist"))

}

// func printTasks(tasks *[]*Todo) {

// 	writer := gocsv.DefaultCSVWriter(os.Stdout)

// 	err := gocsv.MarshalCSV(tasks, writer)
// 	if err != nil {
// 		fmt.Println("Could not output data to file")
// 	}
// }

func DeleteTask(idx int) {
	tasks := MustFetchTasks()
	t := append(tasks[:idx], tasks[idx+1:]...)

	writeTasks(t)

}

func UpdateTask(task *Todo, idx int) {

	tasks := MustFetchTasks()

	tasks[idx] = task

	// printTasks(&tasks)
	defer writeTasks(tasks)
}

func CreateNewTask(task *Todo) {

	tasks := MustFetchTasks()

	tasks = append(tasks, task)

	defer writeTasks(tasks)

}

func writeTasks(tasks []*Todo) {
	file, err := os.OpenFile(PATH, os.O_WRONLY, os.ModePerm)
	Must(err)
	defer file.Close()

	file.Truncate(0) // Will not overwrite content without

	err = gocsv.MarshalFile(tasks, file)
	if err != nil {
		fmt.Println("Could not output data to file")
	}
}

func MustFetchTasks() []*Todo {
	var tasks []*Todo
	file, err := os.Open(PATH)
	Must(err)

	defer file.Close()

	err = gocsv.UnmarshalFile(file, &tasks)
	Must(err)

	return tasks
}

func Must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(9)
	}
}

func MustValue[T interface{}](v T, err error) T {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(9)
	}
	return v
}
