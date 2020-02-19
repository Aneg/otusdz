package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

//Цель: Реализовать утилиту envdir на Go. Эта утилита позволяет запускать программы получая переменные окружения из определенной директории. См man envdir Пример go-envdir /path/to/evndir command arg1 arg2
//Завести в репозитории отдельный пакет (модуль) для этого ДЗ
//Реализовать функцию вида ReadDir(dir string) (map[string]string, error), которая сканирует указанный каталог и возвращает все переменные окружения, определенные в нем.
//Реализовать функцию вида RunCmd(cmd []string, env map[string]string) int , которая запускает программу с аргументами (cmd) c переопределнным окружением.
//Реализовать функцию main, анализирующую аргументы командной строки и вызывающую ReadDir и RunCmd

func main() {
	flag.Parse()
	if len(flag.Args()) < 2 {
		log.Fatalf("не верное количество аргументов")
	}

	stdout, err := RunCmd(flag.Args()[1:], map[string]string{"USER": "ne_gena"})
	if err != nil {
		log.Fatal(stdout)
	}
	fmt.Println(string(stdout))
}

// сканирует указанный каталог и возвращает все переменные окружения, определенные в нем
func ReadDir(dir string) (map[string]string, error) {
	file, err := os.Open(dir)
	if err != nil {
		return nil, errors.New("другие ошибки, например нет прав")
	}
	defer file.Close()
	//b := make([]byte, 1024*1024)
	//read, err := io.ReadFull(file, b)
	//if err != nil {
	//
	//}

	cmd := exec.Command("cd ..")

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	return nil, nil
}

// запускает программу с аргументами (cmd) c переопределнным окружением
func RunCmd(cmd []string, env map[string]string) ([]byte, error) {
	command := exec.Command(cmd[0], cmd[1:]...)
	for name, val := range env {
		command.Env = append(command.Env, name+"="+val)
	}
	return command.Output()
}
