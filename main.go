package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type ProjectConfig struct {
	Name      string `yaml:"-"`
	Language  string `yaml:"-"`
	Structure struct {
		RootDirs    []string `yaml:"root_dirs"`
		DomainDirs  []string `yaml:"domain_dirs"`
		Application struct {
			Dirs []string `yaml:"dirs"`
		} `yaml:"application"`
		Infrastructure struct {
			Dirs []string `yaml:"dirs"`
		} `yaml:"infrastructure"`
		Interfaces struct {
			Dirs []string `yaml:"dirs"`
		} `yaml:"interfaces"`
	} `yaml:"structure"`
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: dddgen <language> <project-name> [config-file]")
		fmt.Println("Available languages: go, python")
		fmt.Println("If config-file is not provided, default structure will be used")
		os.Exit(1)
	}

	config := ProjectConfig{
		Language: strings.ToLower(os.Args[1]),
		Name:     os.Args[2],
	}

	if len(os.Args) > 3 {
		configFile := os.Args[3]
		err := loadConfig(&config, configFile)
		if err != nil {
			fmt.Printf("Error loading config file: %v\n", err)
			os.Exit(1)
		}
	} else {
		setDefaultConfig(&config)
	}

	if config.Language != "go" && config.Language != "python" {
		fmt.Println("Error: unsupported language. Available options: go, python")
		os.Exit(1)
	}

	err := createDDDStructure(config)
	if err != nil {
		fmt.Printf("Error creating project structure: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully created DDD project structure for %s (%s)\n", config.Name, config.Language)
}

func loadConfig(config *ProjectConfig, filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return fmt.Errorf("failed to parse config file: %w", err)
	}

	return nil
}

func setDefaultConfig(config *ProjectConfig) {
	// default structure
	config.Structure.RootDirs = []string{
		"cmd",
		"pkg",
		"configs",
		"migrations",
		"scripts",
		"tests",
		"docs",
	}

	config.Structure.DomainDirs = []string{
		"entities",
		"valueobjects",
		"events",
		"repositories",
		"services",
	}

	config.Structure.Application.Dirs = []string{
		"commands",
		"queries",
		"services",
		"dto",
	}

	config.Structure.Infrastructure.Dirs = []string{
		"persistence",
		"messaging",
		"logging",
		"cache",
	}

	config.Structure.Interfaces.Dirs = []string{
		"rest",
		"graphql",
		"grpc",
		"cli",
	}
}

func createDDDStructure(config ProjectConfig) error {
	baseDir := config.Name

	if err := createAllDirectories(baseDir, config); err != nil {
		return err
	}

	switch config.Language {
	case "go":
		return createGoFiles(baseDir, config.Name)
	case "python":
		return createPythonFiles(baseDir, config.Name)
	default:
		return fmt.Errorf("unsupported language: %s", config.Language)
	}
}

func createAllDirectories(baseDir string, config ProjectConfig) error {
	for _, dir := range config.Structure.RootDirs {
		if err := os.MkdirAll(filepath.Join(baseDir, dir), 0755); err != nil {
			return fmt.Errorf("failed to create root directory %s: %w", dir, err)
		}
	}

	layers := map[string][]string{
		"domain":         config.Structure.DomainDirs,
		"application":    config.Structure.Application.Dirs,
		"infrastructure": config.Structure.Infrastructure.Dirs,
		"interfaces":     config.Structure.Interfaces.Dirs,
	}

	for layer, subdirs := range layers {
		layerPath := filepath.Join(baseDir, "internal", layer)
		if err := os.MkdirAll(layerPath, 0755); err != nil {
			return fmt.Errorf("failed to create layer directory %s: %w", layerPath, err)
		}

		for _, subdir := range subdirs {
			subdirPath := filepath.Join(layerPath, subdir)
			if err := os.MkdirAll(subdirPath, 0755); err != nil {
				return fmt.Errorf("failed to create subdirectory %s: %w", subdirPath, err)
			}
		}
	}

	return nil
}

func createGoFiles(baseDir, projectName string) error {
	mainFile := filepath.Join(baseDir, "cmd", "main.go")
	content := `package main

import "fmt"

func main() {
    fmt.Println("` + projectName + ` service started")
}
`
	if err := os.WriteFile(mainFile, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to create main.go: %w", err)
	}

	return createCommonFiles(baseDir, "go")
}

func createPythonFiles(baseDir, projectName string) error {
	if err := createInitPyFiles(baseDir); err != nil {
		return err
	}

	mainFile := filepath.Join(baseDir, "cmd", "main.py")
	mainContent := fmt.Sprintf(`def main():\n    print("%s service started")\n\nif __name__ == "__main__":\n    main()`, projectName)

	if err := os.WriteFile(mainFile, []byte(mainContent), 0644); err != nil {
		return fmt.Errorf("failed to create main.py: %w", err)
	}

	return createCommonFiles(baseDir, "python")
}

func createInitPyFiles(baseDir string) error {
	return filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			initPath := filepath.Join(path, "__init__.py")
			if _, err := os.Stat(initPath); os.IsNotExist(err) {
				if err := os.WriteFile(initPath, []byte("# Package initialization\n"), 0644); err != nil {
					return fmt.Errorf("failed to create %s: %w", initPath, err)
				}
			}
		}
		return nil
	})
}

func createCommonFiles(baseDir, language string) error {
	if language == "python" {
		requirementsFile := filepath.Join(baseDir, "requirements.txt")
		err := os.WriteFile(requirementsFile, []byte("# Project dependencies\n"), 0644)
		if err != nil {
			return fmt.Errorf("failed to create requirements.txt: %w", err)
		}
	}

	return nil
}
