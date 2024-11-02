package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "crypto/md5"
)

func main() {
    // Открываем файл для чтения
    file, err := os.Open("prog.exe")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    // Читаем содержимое файла
    data, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Вычисляем MD5 хеш
    hash := md5.Sum(data)

    // Записываем хеш в файл
    err = ioutil.WriteFile("md5sum.txt", hash[:], 0644)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Файл md5sum.txt успешно создан.")
}
