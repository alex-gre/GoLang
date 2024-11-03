// go mod init myRastAPI
// REST API на Go
//

package main

// go get github.com/gorilla/mux

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

/*

Первой функцией будет getBooks(), в результате работы которой мы будем видеть список всех книг.
В качестве параметров передаём объект ResponseWriter и указатель на объект Request.

Также мы должны установить заголовок Content-Type, определяющий, с каким типом данных мы будет работать,
а именно – json. Дальше мы формируем ответ и отправляем его пользователю.

*/

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

/*
Функция getBook() будет показывать книгу по определённому id, который мы передаём через URL. 
Для начала передаём знакомые нам параметры запроса и ответа и устанавливаем заголовок. 
Теперь самое главное: нам нужно как-то получить id, передаваемый пользователем, с id книги, 
которая у нас реально существует. Для этого посмотрим на адресную строку:

127.0.0.1:8000/books/1

Чтобы достать GET-параметр 1, воспользуемся функцией библиотеки gorilla/mux
*/

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

/*

Функция createBook() будет отвечать за создание книги.
Для этого мы инициализируем объект структуры Book,
в четвёртой строке мы парсим тело запроса и связываем её с объектом book,
передаваемым по ссылке. Дальше мы формируем случайный ID и включаем новую
книгу в массив books с помощью встроенной функции append.

*/

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

/*
Функции updateBook() и deleteBook() внешне очень похожи: мы также должны получить
переменные запроса с помощью метода Vars(), необходимые для сравнения значений на
изменение и удаление книги. Комбинируя функцию append и оператора среза мы добиваемся того,
 что в первом случае индексируем новый срез с уже обновлённой книгой, который включаем в
 массив с книгами, а во втором – удаляем срез, в котором находится книга.

*/

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

/*
И, наконец, мы должны объявить функцию main – точку входа в наше приложение.
Для начала создадим объект NewRouter() пакета gorilla/mux и добавим две книги в массив.
Теперь осталось прописать пять маршрутов нашего приложения, определить HTTP методы и
заставить приложение слушать порт 8000.
*/

func main() {
	r := mux.NewRouter()
	books = append(books, Book{ID: "1", Title: "Война и Мир", Author: &Author{Firstname: "Лев", Lastname: "Толстой"}})
	books = append(books, Book{ID: "2", Title: "Преступление и наказание", Author: &Author{Firstname: "Фёдор", Lastname: "Достоевский"}})
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}

/*
На запрос 127.0.0.1:8000/books/1 мы должны получить данные о первой книге
Чтобы проверить, что запросы PUT, DELETE и POST работают, воспользуемся Postman.
Пишем json, который хотим отправить на добавление:

Видим, что функция сработала и в нашем списке появилась новая книга.
Теперь попробуем обновить одну из книг. Для этого в Postman выставляем метод PUT, в строке запроса пишем http://127.0.0.1:8000/books/1, где 1 – id обновляемой книги.



REST API на Go
*/
