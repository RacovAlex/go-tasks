package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка соединения с сервером: 404")
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Вы вошли в чат")

	// Ввод имени
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите имя: ")
	name, _ := reader.ReadString('\n')
	// TODO: Подумать, какие могут возникать ошибки и как их лучше обрабатывать
	conn.Write([]byte(name))

	// Чтение с сервера в отдельной горутине (чтобы получать сообщения)
	go readMessages(conn, strings.TrimSpace(name))

	for {
		msg, _ := reader.ReadString('\n')
		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("Ваше сообщение не было отправлено, произошла ошибка: ", err)
		}

	}
}

// readMessages обрабатывает сообщения с сервера
func readMessages(conn net.Conn, name string) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Не удалось получить сообщение, произошла ошибка: ", err)
			os.Exit(1)
		}
		if len(msg) == 0 {
			// TODO:
		}
		user := strings.Split(msg, ":")[0]
		if user != name {
			fmt.Print(msg)
		}

	}
}
