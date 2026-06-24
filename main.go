package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func deleteTasks(scanner *bufio.Scanner, tasks []task) []task {
	var i int
	fmt.Println("Хотите удалить задачу?")
	scanner.Scan()
	if scanner.Text() == "да" {
		fmt.Println("Введите номер задачи:")
		_, err := fmt.Scanln(&i)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			return tasks
		} else {
			if i > 0 && i <= len(tasks) {
				tasks = slices.Delete(tasks, i-1, i)
				return tasks
			} else {
				fmt.Println("Несуществующий номер задачи.")
				return tasks
			}
		}
	} else {
		return tasks
	}
	
}

func completeCheck(b bool) string {
	var check string
	switch {
	case b == true:
		check = "выполнено"
	case b == false:
		check = "не выполнено"
	}

	return check
}

func viewTasks(tasks []task) {
	fmt.Println("Вывод списка задач:")
	for i := range len(tasks) {
		fmt.Printf("\nНазвание задачи:%s\nОписание задачи:%s\nСтатус задачи:%s", tasks[i].Name, tasks[i].Description, completeCheck(tasks[i].Complete))
	}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tasks := createTasks(scanner)
	viewTasks(tasks)
	deleteTasks(scanner, tasks)
	viewTasks(tasks)
}
