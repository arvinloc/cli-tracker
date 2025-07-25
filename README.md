# 📌 CLI Tracker

> A simple CLI application written in Go for tracking tasks with JSON-based storage.

---

## ⚙️ Features

- ✅ Add new tasks
- ✏️ Update task status (`in progress` → `done`)
- 📋 List tasks by status
- 💾 Store tasks in `tasks.json`

---

## 🚀 Installation & Run
```bash
git clone https://github.com/your-username/cli-tracker.git
cd cli-tracker
go run cmd/main.go [флаги]
```

---

## Flags & Commands

| Flag              | Description                                  | Example usage                               |
|-------------------|----------------------------------------------|---------------------------------------------|
| `-add`            | Add a new task                               | `go run cmd/main.go -add "Сделать домашку"` |
| `-list`           | Show tasks by status (`done`, `in progress`) | `go run cmd/main.go -list "done"`           |
| `-mark` + `-list` | Update status by ID                          | `go run cmd/main.go -mark 1 -list "done"`   |

> ⚠️ Status can only be done or in progress.

---

## 📁 Usage Examples

### ➕ Add a Task

```bash
go run cmd/main.go -add "Прочитать книгу по Go"
```

**Output:**
```
✅ Task successfully created (ID:1)
```

---

### ✅ Update Task Status

```bash
go run cmd/main.go -mark 1 -list "done"
```

**Output:**
```
✅ Task updated: &{ID:1 Description:Прочитать книгу по Go Status:done}
```

---

### 📋 View Tasks by Status

```bash
go run cmd/main.go -list "done"
```

**Output:**
```
📋 Tasks:
- [1] Read a Go programming book (done)
```

---

## 📄 License
This project is licensed under the [MIT](./LICENSE).
[GitHub](https://github.com/arvinloc)