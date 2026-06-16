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

func main() {
	var t task
	tasks := make([]task, 0)
	scanner := bufio.NewScanner(os.Stdin)
	t = scanTask(scanner)
	tasks = append(tasks, t)
}
