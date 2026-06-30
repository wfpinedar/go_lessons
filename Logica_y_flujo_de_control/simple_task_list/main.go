package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Name     string
	Complete bool
}

// func addTask(tasks []Task, taskName string) []Task {
// 	tasks = append(tasks, Task{taskName, false})
// 	return tasks
// }

func addTask(tasks []Task, taskName string, status bool) []Task {
	tasks = append(tasks, Task{taskName, status})
	return tasks
}

func viewAllTasks(tasks []Task) {
	var complete int
	var remaining int
	for _, task := range tasks {
		if task.Complete {
			remaining = remaining + 1
		} else {
			remaining = remaining + 1
		}
	}
	fmt.Printf("Total: %d tasks (%d completed, %d remaining)", len(tasks), complete, remaining)
}

func main() {

	// Leer entrada
	// var appName string
	// var userName string
	// fmt.Scanln(&appName)
	// fmt.Scanln(&userName)

	// // TODO: Escribe tu código a continuación
	// // Definir la estructura Task
	// // Crear un slice de estructuras Task
	tasks := []Task{}
	// // Mostrar mensaje de bienvenida y menú
	// menu := "1. Add Task\n2. View Tasks\n3. Complete Task\n4. Remove Task\n5. Exit"
	// fmt.Printf("Welcome to %s, %s!\n", appName, userName)
	// fmt.Println(menu)
	// // Imprimir el recuento actual de tareas
	// fmt.Printf("Current tasks: %d", len(tasks))
	var prevTasks string
	fmt.Scanln(&prevTasks)
	// existingTasks, _ := strconv.Atoi(prevTasks)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	taskNames := scanner.Text()
	// for i := 0; i < existingTasks; i++ {
	// 	taskName := "Existing Task " + strconv.Itoa(i+1)
	// 	tasks = addTask(tasks, taskName, false)
	// }
	taskNamesSplit := strings.Split(taskNames, ",")
	// for _, taskName := range taskNamesSplit {
	// 	if taskName != "" {
	// 		tasks = addTask(tasks, taskName)
	// 	}
	// }
	for _, taskName := range taskNamesSplit {
		if taskName != "" {
			tk := strings.Split(taskName, ":")
			st, _ := strconv.ParseBool(tk[1])
			tasks = addTask(tasks, tk[0], st)
		}
	}
	// for _, task := range tasks {
	// 	fmt.Printf("Task: %s, Completed: %t\n", task.Name, task.Complete)
	// }
	// fmt.Printf("Total tasks: %d", len(tasks))
	viewAllTasks(tasks)

}
