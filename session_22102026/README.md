# Session 22-10-2026: Todo Application Deliverables

## 📋 Overview

This session focuses on the end-to-end design, data modeling, API contract specification, and automated Go HTTP server boilerplate code generation for the **Todo List Application**.

---

## 📁 Repository Structure

```text
session_22102026/
├── README.md                                    # Session overview
├── Sarthak_Mahadik_Project_Deliverables_Session_22102026.zip
└── Sarthak_Mahadik_Session_22102026/
    ├── README.md                                # Detailed documentation
    ├── go.mod                                   # Go module definition (todo-api)
    ├── oapi-codegen.yaml                        # Configuration for oapi-codegen tool
    ├── oapi_codegen_go_project_steps.docx       # Step-by-step setup and generation guide
    ├── swagger.yaml                             # OpenAPI 3.0.3 specification for Todo API
    ├── swagger1.yml                             # Supporting Swagger draft
    ├── api/
    │   └── api.gen.go                           # Generated Go models & standard HTTP server handlers
    ├── data_model_diagram/
    │   ├── migration.sql                        # MySQL DDL migration script generated from DBML
    │   ├── todo.dbml                            # Database Markup Language (DBML) schema definition
    │   └── todo_dbml.svg                        # Vector ER diagram visualizing Users & Todos
    └── todo_user_stories/
        ├── old-user-stories.md                  # Initial draft of user stories
        ├── todo_user_stories.html               # Formatted HTML report of user stories
        └── todo_user_stories.md                 # Detailed user stories and acceptance criteria
```

---

## 🚀 Key Tasks & Deliverables Completed

### 1. Requirements Gathering & User Stories (`todo_user_stories/`)
- Drafted and refined user stories detailing features and behavior for the Todo system:
  - **US-001: Create a Todo** — Input validation, priority, category, and feedback.
  - **US-002: View Todos** — Displaying ordered tasks, statuses, and empty states.
  - **US-003: Update a Todo** — Modifying existing todo titles, categories, priorities.
  - **US-004: Delete a Todo** — Soft/hard removal with user feedback.
  - **US-005: Search Todos** — Finding items by title or ID.
  - **US-006: Filter Todos** — Filtering by completion status, priority level, category, and archived state.
  - **US-007: Mark Status** — Toggling status between `pending` and `completed`.
- Delivered in both Markdown (`todo_user_stories.md`) and styled HTML (`todo_user_stories.html`).

---

### 2. Database Schema Design & Migration (`data_model_diagram/`)
- **DBML Schema (`todo.dbml`)**:
  - Defined the relational schema connecting `Users` and `Todos`.
  - Configured foreign key constraint: `Todos.user_id > Users.user_id`.
  - Configured field constraints, auto-increment primary keys, default statuses (`'Pending'`), and timestamps.
- **ER Diagram (`todo_dbml.svg`)**:
  - Visualized the entity-relationship model and cardinality between tables.
- **SQL Migration Script (`migration.sql`)**:
  - Generated MySQL-compatible DDL migration:
    - `CREATE TABLE Users (...)`
    - `CREATE TABLE Todos (...)`
    - `ALTER TABLE Todos ADD FOREIGN KEY (user_id) REFERENCES Users(user_id);`

---

### 3. API Contract Specification (`swagger.yaml`)
- Authored the standard OpenAPI 3.0.3 specification (`swagger.yaml`) defining the RESTful interface:
  - `GET /todos` — Retrieve todos with search query and query filters (`status`, `priority`, `category`, `archived`).
  - `POST /todos` — Create a new todo item with input validation.
  - `GET /todos/{id}` — Retrieve a single todo by path parameter `id`.
  - `PUT /todos/{id}` — Full update of a specific todo item.
  - `DELETE /todos/{id}` — Delete a specific todo.
  - `PATCH /todos/{id}/archive` — Archive an active todo item.
- Reusable schemas:
  - `TodoInput`: Required `title`, optional `description`, `status` enum (`pending`, `completed`), `priority` enum (`low`, `medium`, `high`), and `category`.
  - `Todo`: Inherits `TodoInput` and includes read-only fields `id`, `archived`, `createdAt`, `updatedAt`.

---

### 4. Go HTTP Boilerplate Code Generation (`oapi-codegen`)
- **Configuration (`oapi-codegen.yaml`)**:
  - Package name: `api`
  - Output path: `api/api.gen.go`
  - Enabled generators: `models: true`, `std-http-server: true`.
- **Go Module (`go.mod`)**:
  - Initialized module `todo-api` targeting Go 1.22+.
- **Generated Code (`api/api.gen.go`)**:
  - Generated with `oapi-codegen/v2` (v2.8.1).
  - Includes:
    - Go types: `Todo`, `TodoInput`, `GetTodosParams`, `TodoPriority`, `TodoStatus`.
    - `ServerInterface`: The Go interface for HTTP handlers (`GetTodos`, `CreateTodo`, `GetTodo`, `UpdateTodo`, `DeleteTodo`, `ArchiveTodo`).
    - Standard library `net/http` routing adapters and multiplexer registration (`Handler`, `HandlerFromMux`, `HandlerWithOptions`).
- **Workflow Guide (`oapi_codegen_go_project_steps.docx`)**:
  - Step-by-step guide documenting prerequisites, installation, command execution, and troubleshooting tips.

---

## 🛠️ How to Re-generate the Code

From the session directory (`session_22102026/Sarthak_Mahadik_Session_22102026`):

```bash
# 1. Install oapi-codegen (if not already installed)
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

# 2. Run the code generator using the YAML config
oapi-codegen -config .\oapi-codegen.yaml .\swagger.yaml
```
