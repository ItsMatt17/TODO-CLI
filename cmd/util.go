package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
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

func processCSV(rec [][]string) []Todo {
	var tasks []Todo
	for i, val := range rec {
		if i == 0 {
			continue
		}

		t := Todo{}

		t.Id = uint8(MustValue(strconv.ParseUint(val[0], 10, 8))) // Causes issuse if you don't do it this way lol prob a better way
		t.Task = val[1]
		t.CreationDate = MustValue(strconv.ParseInt(val[2], 10, 64))
		t.Completion = MustValue(strconv.ParseBool(val[3]))

		tasks = append(tasks, t)
	}
	return tasks
}

func fetchNextId() int {
	return len(MustFetchTasks()) + 1
}

func CreateNewTask(name string) {
	task := []string{strconv.FormatInt(int64(fetchNextId()), 10),
		name, strconv.FormatInt(time.Now().Unix(), 10), strconv.FormatBool(false)}

	file, err := os.OpenFile("tasks.csv", os.O_APPEND, 0644)
	Must(err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(task)
	Must(err)
}

func mustFetchCSV() [][]string {
	file, err := os.Open("tasks.csv")
	Must(err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	Must(err)

	return records
}

func MustFetchTasks() []Todo {
	records := mustFetchCSV()

	tasks := processCSV(records)
	fmt.Println(tasks)
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
