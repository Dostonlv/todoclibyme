package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	Id          int
	Title       string
	Description string
	Completed   bool
}

var todos []Todo

func main() {
	fmt.Println("Welcome to Todo CLI")
	fmt.Println("1. Create Todo")
	fmt.Println("2. Update Todo")
	fmt.Println("3. Read Todo")
	fmt.Println("4. Delete Todo")
	fmt.Println("5. Exit")
	fmt.Println("Enter your choice")
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	switch choice {
	case "1":
		createTodo()
	case "2":
		updateTodo()
	case "3":
		readTodo()
	case "4":
		deleteTodo()
	case "5":
		os.Exit(0)
	default:
		fmt.Println("Invalid choice")
	}
}

func createTodo() {
	fmt.Println("Enter title")
	reader := bufio.NewReader(os.Stdin)
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	fmt.Println("Enter description")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)
	todo := Todo{Id: len(todos) + 1, Title: title, Description: description, Completed: false}
	todos = append(todos, todo)
	fmt.Println("Todo created successfully")
	main()
}

func updateTodo() {
	fmt.Println("Enter todo id")
	reader := bufio.NewReader(os.Stdin)
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)
	todoId, _ := strconv.Atoi(id)
	for i, todo := range todos {
		if todo.Id == todoId {
			fmt.Println("Enter title")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			fmt.Println("Enter description")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSpace(description)
			todos[i].Title = title
			todos[i].Description = description
			fmt.Println("Todo updated successfully")
			main()
		}
	}
	fmt.Println("Todo not found")
	main()
}

func readTodo() {
	fmt.Println("Enter todo id")
	reader := bufio.NewReader(os.Stdin)
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)
	todoId, _ := strconv.Atoi(id)
	for _, todo := range todos {
		if todo.Id == todoId {
			fmt.Println("Title: ", todo.Title)
			fmt.Println("Description: ", todo.Description)
			fmt.Println("Completed: ", todo.Completed)
			main()
		}
	}
	fmt.Println("Todo not found")
	main()
}

func deleteTodo() {
	fmt.Println("Enter todo id")
	reader := bufio.NewReader(os.Stdin)
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)
	todoId, _ := strconv.Atoi(id)
	for i, todo := range todos {
		if todo.Id == todoId {
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Println("Todo deleted successfully")
			main()
		}
	}
	fmt.Println("Todo not found")
	main()
}
