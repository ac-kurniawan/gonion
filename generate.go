package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	template2 "text/template"
	"unicode"
)

func CreateDir(dirName string) {
	if err := os.MkdirAll(dirName, os.ModePerm); err != nil {
		panic(err)
	}
}

func AddFile(path, content string) {
	if err := os.WriteFile(path, []byte(content), os.ModePerm); err != nil {
		panic(err)
	}
}

func HydrateTemplate(template string, data map[string]interface{}) string {
	tmpl, err := template2.New(template).Parse(template)
	if err != nil {
		panic(err)
	}
	var text bytes.Buffer
	err = tmpl.Execute(&text, data)
	if err != nil {
		panic(err)
	}
	return text.String()
}

func capitalize(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func getModuleName(input string) string {
	split := strings.Split(input, "/")
	length := len(split)

	return capitalize(split[length-1])
}

func Generate(repository string) {
	split := strings.Split(repository, "/")
	length := len(split)
	rootPath := split[length-1]
	rootPath = "./" + rootPath

	directories := []string{CoreDirectory, CoreModelDirectory, AdaptorDirectory, InterfaceDirectory, LibraryDirectory}

	for _, dir := range directories {
		CreateDir(rootPath + "/" + dir)
		AddFile(fmt.Sprintf("./%s/%s/.gitkeep", rootPath, dir), "")
	}

	data := map[string]interface{}{
		"moduleRepository": repository,
		"moduleName":       getModuleName(repository),
	}
	AddFile(rootPath+"/properties.yml", HydrateTemplate(PropertiesFile, data))
	AddFile(rootPath+"/app.go", HydrateTemplate(AppFile, data))
	AddFile(rootPath+"/main.go", HydrateTemplate(MainFile, data))
	AddFile(rootPath+"/core/service.go", HydrateTemplate(CoreServiceFile, data))
	AddFile(rootPath+"/core/repository.go", HydrateTemplate(CoreRepositoryFile, data))
	AddFile(rootPath+"/core/Dockerfile", HydrateTemplate(DockerFile, data))
	AddFile(rootPath+"/core/.gitignore", HydrateTemplate(GitIgnoreFile, data))
	AddFile(rootPath+"/go.mod", HydrateTemplate(GoModFile, data))

	fmt.Println("install dependencies")

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "cd", fmt.Sprintf("./%s", getModuleName(repository)), "&&", "go", "mod", "tidy")

	} else {
		cmd = exec.Command("bash", "-c", "cd", fmt.Sprintf("./%s", getModuleName(repository)), "&&", "go", "mod", "tidy")
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
