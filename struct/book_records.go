package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Client struct {
	Name  string
	Phone string
	Email string
}

var clients []Client

func main() {
	loadClientsFromFile("clients.csv")

	for {
		fmt.Println()
		fmt.Println("1. Add a client")
		fmt.Println("2. Save clients to file")
		fmt.Println("3. Search clients by name")
		fmt.Println("4. Display all clients")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addClient()
		case 2:
			saveClientsToFile("clients.csv")
		case 3:
			searchClientsByName()
		case 4:
			displayAllClients()
		case 5:
			fmt.Println("Exiting program.")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func addClient() {
	var client Client

	fmt.Print("Enter client's name: ")
	fmt.Scanln(&client.Name)

	fmt.Print("Enter client's phone number: ")
	fmt.Scanln(&client.Phone)

	fmt.Print("Enter client's email address: ")
	fmt.Scanln(&client.Email)

	clients = append(clients, client)
	fmt.Println("Client added successfully.")
}

func saveClientsToFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Name", "Phone", "Email"}
	writer.Write(headers)

	for _, client := range clients {
		row := []string{client.Name, client.Phone, client.Email}
		writer.Write(row)
	}

	fmt.Println("Clients saved to file.")
}

func loadClientsFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for _, record := range records {
		client := Client{
			Name:  record[0],
			Phone: record[1],
			Email: record[2],
		}
		clients = append(clients, client)
	}

	fmt.Println("Clients loaded from file.")
}

func searchClientsByName() {
	var name string
	fmt.Print("Enter name to search for: ")
	fmt.Scanln(&name)

	fmt.Println("Search results:")
	for _, client := range clients {
		if strings.Contains(strings.ToLower(client.Name), strings.ToLower(name)) {
			fmt.Printf("Name: %s, Phone: %s, Email: %s\n", client.Name, client.Phone, client.Email)
		}
	}
}

func displayAllClients() {
	fmt.Println("All clients:")
	for _, client := range clients {
		fmt.Printf("Name: %s, Phone: %s, Email: %s\n", client.Name, client.Phone, client.Email)
	}
}