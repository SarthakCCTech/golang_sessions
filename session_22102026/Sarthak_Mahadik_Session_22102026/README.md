# Session 22-10-2026: Todo API Project

A beginner-friendly Go project for a **Todo List API**, built using standard design and automated code generation workflows.

---

## 📌 What Was Done

1. **User Stories (`todo_user_stories/`)**
   - Defined user requirements and acceptance criteria for creating, viewing, updating, deleting, searching, filtering, and archiving todos.

2. **Database Design (`data_model_diagram/`)**
   - Modeled **Users** and **Todos** tables with relationships in DBML (`todo.dbml` & visual diagram `todo_dbml.svg`).
   - Created a MySQL migration script (`migration.sql`) to build the database tables.

3. **API Specification (`swagger.yaml`)**
   - Wrote an OpenAPI 3.0 specification defining REST endpoints for `/todos` (create, read, update, delete, search, filter, archive).

4. **Go Code Generation (`api/api.gen.go`)**
   - Used `oapi-codegen` to automatically convert `swagger.yaml` into Go data structs and standard `net/http` server routing boilerplate.
   - Documented the setup steps in `oapi_codegen_go_project_steps.docx`.

---

## 📁 Project Structure

```text
├── api/
│   └── api.gen.go                      # Generated Go structs & HTTP server interface
├── data_model_diagram/
│   ├── todo.dbml                       # Database schema definition
│   ├── todo_dbml.svg                   # Visual database diagram
│   └── migration.sql                   # MySQL table creation script
├── todo_user_stories/
│   ├── todo_user_stories.md            # Detailed user stories
│   └── todo_user_stories.html          # HTML report of user stories
├── swagger.yaml                        # OpenAPI 3.0 API blueprint
├── oapi-codegen.yaml                   # Code generator configuration
├── oapi_codegen_go_project_steps.docx  # Step-by-step setup guide
├── go.mod                              # Go module definition
└── README.md                           # Project summary
```

---

## 🚀 How to Re-generate Go Code

From this folder, run:

```powershell
oapi-codegen -config .\oapi-codegen.yaml .\swagger.yaml
```
