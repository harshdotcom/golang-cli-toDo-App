package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks []Task

func main() {
	fmt.Println("Welcome to the To-Do App!")

	loadTasksFromFile()

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Save & Exit")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var title string
			fmt.Print("Enter task name: ")
			fmt.Scanln(&title)
			addTask(title)
		case 2:
			listTasks()
		case 3:
			var id int
			listTasks()
			fmt.Print("Enter task ID to mark done: ")
			fmt.Scanln(&id)
			if id > 0 {
				markTaskDone(id)
			}
		case 4:
			var id int
			listTasks()
			fmt.Print("Enter task ID to delete: ")
			fmt.Scanln(&id)
			if id > 0 {
				deleteTask(id)
			}
		case 5:
			saveTasksToFile()
			fmt.Println("Tasks saved. Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func addTask(title string) {

	newTask := Task{
		Id:    len(tasks) + 1,
		Title: title,
		Done:  false,
	}

	tasks = append(tasks, newTask)
	fmt.Println("Task added successfully!!")
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	fmt.Println("Your tasks")

	for _, task := range tasks {
		status := "Pending"
		if task.Done {
			status = "Done"
		}
		fmt.Printf("%d. %s [%s]\n", task.Id, task.Title, status)
	}
}

func markTaskDone(id int) {
	for i, task := range tasks {
		if task.Id == id {
			tasks[i].Done = true
			fmt.Println("Task Marked as done!")
			return
		}
	}
	fmt.Println("Task not found")
}

func deleteTask(id int) {
	for i, task := range tasks {
		if task.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task deleted sucessfully!")
			return
		}
	}
	fmt.Println("Task not found.")
}

func saveTasksToFile() {
	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error Writing to file:", err)
	}
}

func loadTasksFromFile() {
	file, err := os.Open("tasks.json")
	if err != nil {
		fmt.Println("No existing tasks found. Starting fresh!")
		return
	}
	defer file.Close()

	data := json.NewDecoder(file)
	err = data.Decode(&tasks)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Tasks loaded successfully!")
}
