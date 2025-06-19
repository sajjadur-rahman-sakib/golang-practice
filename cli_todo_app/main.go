package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func displayMenu() {
	fmt.Println("Welcome to the CLI task manager")
	fmt.Println("1. Add Task")
	fmt.Println("2. View Tasks")
	fmt.Println("3. Mark as Completed")
	fmt.Println("4. Delete Task")
	fmt.Println("5. Exit")
}

func main() {

	var tasks []string
	reader := bufio.NewScanner(os.Stdin)

	displayMenu()

	for {
		fmt.Println("Please enter your choice: ")

		reader.Scan()
		input := strings.TrimSpace(reader.Text())
		choice, error := strconv.Atoi(input)

		if error != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			return
		}

		switch choice {
		case 1:
			// Add Task
			fmt.Println("Enter task: ")
			reader.Scan()
			task := strings.TrimSpace(reader.Text())

			// Add the task to the list
			tasks = append(tasks, task)
			fmt.Println("Task added successfully!")

		case 2:
			// View Tasks
			if len(tasks) == 0 {
				fmt.Println("No tasks available.")
			} else {
				fmt.Println("Tasks:")
				for index, task := range tasks {
					fmt.Printf("%d - %s \n", index+1, task)
				}
			}

		case 3:
			// Mark as Completed
			if len(tasks) == 0 {
				fmt.Println("No tasks available.")
			}

			fmt.Println("Enter task number to complete: ")

			reader.Scan()
			taskIdString := strings.TrimSpace(reader.Text())
			taskId, error := strconv.Atoi(taskIdString)

			if error == nil && taskId > 0 && taskId <= len(tasks) {
				tasks[taskId-1] = tasks[taskId-1] + " (Completed)"
				fmt.Println("Task completed successfully!")
			} else {
				fmt.Println("Invalid input. Please enter a valid number.")
			}

		case 4:
			// Delete Task
			if len(tasks) == 0 {
				fmt.Println("No tasks available.")
			}

			fmt.Println("Enter task number to delete: ")

			reader.Scan()
			taskIdString := strings.TrimSpace(reader.Text())
			taskId, error := strconv.Atoi(taskIdString)

			if error == nil && taskId > 0 && taskId <= len(tasks) {
				tasks = append(tasks[:taskId-1], tasks[taskId:]...)
				fmt.Println("Task deleted successfully!")
			} else {
				fmt.Println("Invalid input. Please enter a valid number.")
			}

		case 5:
			// Exit
			fmt.Println("Exiting the task manager. Goodbye!")
			os.Exit(1)

		default:
			fmt.Println("Invalid choice. Please select a valid option from the menu.")

		}
	}

}
