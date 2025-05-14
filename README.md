# 🧱 ddd-template-gen

A **Domain-Driven Design (DDD)** project structure generator with support for Go and Python.

Creates only directories and modules specified in the YAML configuration file — no unnecessary generation.

---

## 🚀 Features

- ✅ **Go** and **Python** support
- ✅ Generation strictly based on YAML configuration
- ✅ Clean DDD structure with layer separation
- ✅ Automatic `main.go` / `main.py` creation at project entry point

---

## 🔧 Installation

Build the binary:

```bash
go build -o dddgen
```

## 📦 Usage

./dddgen <language> <project-name> [config.yaml]

| Argument         | Description                                       |
| ---------------- | ---------------------------------------------- |
| `language`       | Project language: `go` or `python`             |
| `project-name`   | Name of the project directory to be created    |
| `structure.yaml` | *(optional)* Path to YAML configuration file     |


**Examples**

**Create a Go project with default structure:**

```bash
./dddgen go myproject
```

**Create a Python project with custom YAML configuration:**
```bash
./dddgen python awesome_project structure.yaml
```

## 🧾 Example YAML Configuration

```bash 
structure:
  root_dirs:
    - cmd
    - pkg
    - configs
    - deployments
    - tests
    - docs
    - api
    - internal

  domain_dirs:
    - models
    - repositories
    - services
    - events
    - exceptions

  application:
    dirs:
      - commands
      - queries
      - handlers
      - validators

  infrastructure:
    dirs:
      - database
      - messaging
      - cache
      - logging

  interfaces:
    dirs:
      - http
      - grpc
      - cli
      - events
```

## 🗂 Example Generated Structure

```bash 
project/
├── cmd/
│   └── main.go (или main.py)
├── pkg/
├── configs/
├── deployments/
├── tests/
├── docs/
├── api/
├── internal/
│   ├── domain/
│   │   ├── models/
│   │   ├── repositories/
│   │   └── ...
│   ├── application/
│   │   └── handlers/
│   ├── infrastructure/
│   │   └── database/
│   └── interfaces/
│       └── http/
```