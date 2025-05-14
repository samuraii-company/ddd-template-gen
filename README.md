# 🧱 ddd-template-gen

Генератор структуры проекта в стиле **Domain-Driven Design (DDD)** с поддержкой **Go** и **Python**. Позволяет создавать только указанные в YAML-файле директории и модули — никакой лишней генерации.

---

## 🚀 Возможности

- ✅ Поддержка **Go** и **Python**
- ✅ Генерация строго по YAML-конфигурации
- ✅ Чистая структура DDD с разбивкой на слои
- ✅ Автоматическое добавление `main.go` / `main.py` в точке входа проекта

---

## 🔧 Установка

Соберите бинарный файл:

```bash
go build -o dddgen
```

## 📦 Использование

./dddgen <language> <project-name> [config.yaml]

| Аргумент         | Описание                                       |
| ---------------- | ---------------------------------------------- |
| `language`       | Язык проекта: `go` или `python`                |
| `project-name`   | Имя создаваемой директории проекта             |
| `structure.yaml` | *(опционально)* Путь к YAML-файлу конфигурации |


**Примеры**

**Создание проекта на Go со структурой по умолчанию:**

```bash
./dddgen go myproject
```

**Создание Python-проекта по кастомному YAML:**
```bash
./dddgen python awesome_project config.yaml
```

## 🧾 Пример YAML-конфигурации

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

## 🗂 Пример сгенерированной структуры

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
└── README.md
```