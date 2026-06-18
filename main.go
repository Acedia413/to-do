package main

import (
	"bufio"
	"fmt"
	"os"
)

type task struct {
	Name, Description string
	Complete          bool
}

func scanTask(scanner *bufio.Scanner) task {
	var t task
	fmt.Println("Введите название задачи:")

	if scanner.Scan() {
		t.Name = scanner.Text()
	}

	err := scanner.Err()
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
	}

	fmt.Println("Введите описание задачи:")

	if scanner.Scan() {
		t.Description = scanner.Text()
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
	}

	return t
}

func createTasks(scanner *bufio.Scanner) []task {
	tasks := make([]task, 0)
	fmt.Println("Добавить задачу?")
	scanner.Scan()
	for scanner.Text() == "да" {
		t := scanTask(scanner)
		tasks = append(tasks, t)
		fmt.Println("Добавить ещё задачу?")
		scanner.Scan()
	}
	return tasks
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tasks := createTasks(scanner)
	fmt.Println(tasks)
}
