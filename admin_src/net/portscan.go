//Функция сканирования портов на Go можно реализовать с помощью библиотеки net.
//Вот пример кода:

package main

import (
	"fmt"
	"net"
	"sync"
)

func scanner(port int) {
	// Создаем сокет для сканирования
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "192.168.1.1", port))
	if err != nil {
		// Если сокет не был создан, то порт закрыт
		//fmt.Printf("Порт %d закрыт\n", port)

	} else {
		// Если сокет был создан, то порт открыт
		fmt.Printf("Порт %d открыт\n", port)

		// Закрываем сокет
		conn.Close()
		return

		if opErr, ok := err.(*net.OpError); ok && opErr.Temporary() {
			// Если ошибка временная, то мы продолжаем сканирование
			return
		}
	}

}

func main() {
	// Укажите диапазон портов, которые хотите сканировать
	startPort := 1
	endPort := 65535
	//or max value 65535

	var wg sync.WaitGroup

	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			scanner(port)
		}(port)
	}

	wg.Wait()

}

/*
Этот код сканирует все порты в диапазоне от `startPort` до `endPort` и выводит результат в консоль.
 Если сокет был создан для порта, то это означает, что порт открыт, иначе - закрыт.

Обратите внимание, что этот код может занять некоторое время для сканирования всех портов,
особенно если вы сканируете большой диапазон. Кроме того, некоторые системы могут блокировать
сканирование портов, поэтому результаты могут не быть точными.

Чтобы улучшить производительность, вы можете использовать горутинное параллельное
программирование, как показано в коде. Это позволяет сканировать несколько портов
одновременно, что может ускорить процесс сканирования.

Также обратите внимание, что этот код не проверяет, является ли порт в использовании или
занят другим процессом.
*/