package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Welcome to the admin panel of your toDoList")
		fmt.Println("1. Create task")
		fmt.Println("2. Delete task")
		fmt.Println("3. Update status")
		fmt.Println("4. Edit task")
		fmt.Println("5. Show task(s)")
		fmt.Println("6. Exit")
		fmt.Print("Select an option: ")

		var input int
		fmt.Scanln(&input)
		switch input {
		case 1:
			fmt.Println("Input title")
			scanner.Scan()
			title := scanner.Text()
			todos.add(title)
		case 2:
			fmt.Println("Input index of the task")
			var index int
			fmt.Scanln(&index)
			todos.delete(index)
		case 3:
			fmt.Println("Which task did you do?")
			var index int
			fmt.Scanln(&index)
			todos.toggle(index)
		case 4:
			fmt.Println("Which one do you want to edit (input index)?")
			var index int
			fmt.Scanln(&index)
			fmt.Println("Which new title do you want to set?")
			scanner.Scan()
			title := scanner.Text()
			todos.edit(index, title)
		case 5:
			todos.print()
		case 6:
			fmt.Println("Exiting...")
			storage.Save(todos)
			return
		default:
			fmt.Println("Invalid action. Exiting...")
		}
	}
	storage.Save(todos)
}
