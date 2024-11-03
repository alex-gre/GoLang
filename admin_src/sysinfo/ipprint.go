package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	// Выполняем команду ipconfig на Windows или ip addr на Linux/MacOS
	cmd := exec.Command("ipconfig")
	// или cmd := exec.Command("ip addr")

	// Создаем пайп для вывода команды
	out, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()

	// Запускаем команду
	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Читаем вывод команды
	bytes, err := ioutil.ReadAll(out)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Выводим вывод команды
	fmt.Println(string(bytes))

	// Останавливаем команду
	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
	}
}
