# To-Do List (CLI)

## Description
A simple console application for managing a list of tasks. The functions of adding, deleting, editing, and marking tasks are implemented. The data is saved in a JSON file.

## Technology stack
- **Programming language:** Go
- **Data Storage:** JSON file
- **Libraries used:**
- `encoding/json' — working with JSON
- `bufio` and `os` — data input and output
- `errors` — error handling
- `github.com/aquasecurity/table ` — formatting the output of tasks in a table
  - `strconv' — converting numbers to strings
  - `time` — working with dates and times

## Installation
1. Make sure that you have Go (version 1.18+) installed. 
2. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/todo-cli.git
   cd todo-cli
   ```
3. Assemble and run the program:
   ```sh
   go run main.go
   ```

## Usage
When you start the program, a menu with available actions will be displayed in the terminal.:
```
Welcome to the admin panel of your toDoList
1. Create task
2. Delete task
3. Update status
4. Edit task
5. Show task(s)
6. Exit
Select an option:
```
Select the desired action by entering the appropriate number.

### Basic commands:
- **Add Task**: Enter 1, then enter the task title.
- **Delete Task**: Enter 2, then specify the task number.
- **Mark completion**: Enter 3, then the number of the completed task.
- **Edit task**: Enter 4, then the number and the new task title.
- **View Tasks**: Enter 5 to see a list of all tasks.
- **Exit**: Enter 6 to exit with the changes saved.

## Methods
### `add(title string)`
Adds a new task to the list.
### `delete(index int) error`
Deletes an issue by index. If the index is incorrect, it returns an error.
### `toggle(index int) error`
Changes the status of the task (completed/not completed). If the task becomes completed, the completion time is saved.
### `edit(index int, title string) error`
Edits the task title by index. If the index is incorrect, it returns an error.
### `print()`
Displays a list of tasks in the form of a table with the headings: `#`, `Title`, `Completed`, `Created At`, `Completed At'.

## Project structure
```
.
├── main.go        # Основной файл с логикой обработки команд
├── storage.go     # Файл для работы с JSON-хранилищем
├── todo.go        # Определение структуры задачи и функций управления
├── todos.json     # Файл для хранения списка задач
├── go.mod         # Модульные зависимости проекта
├── go.sum         # Контрольные суммы зависимостей
```

## Data storage format (todos.json)
```json
[
    {
        "Title": "Buy Milk",
        "Completed": false,
        "CreatedAt": "2025-03-27T12:27:08.055146+05:00",
        "CompletedAt": null
    }
]
```

## Автор
[Kairat Bagakhan]
