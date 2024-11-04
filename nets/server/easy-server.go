package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"os"
	"strings"
)

var (
	clients = make(map[net.Conn]string) // Список клиентов
	// TODO: использовать sync.Map
	newClients = make(chan net.Conn) // Канал для сообщений
	messages   = make(chan string)   // Канал для новых клиентов
)

func main() {
	// Cоздаем "слушателя" на указанном порту 8080, который будет принимать входящие TCP-соединения
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		slog.Error("ошибка при запуске сервера: %w", err)
		os.Exit(1)
	}
	defer func(ln net.Listener) {
		err := ln.Close()
		if err != nil {
			slog.Error("ошибка закрытия соединения: %w", err)
		}
	}(listener)

	slog.Info("Сервер запущен и слушает порт: 8080")

	// Запускаем обработку новых подключений
	go handleNewClients()

	// Запускаем обработку сообщений
	go handleMessages()

	// Ожидаем подключений
	for {
		// Ожидаем соединения клиента
		conn, err := listener.Accept()
		if err != nil {
			slog.Error("Соединение не установлено: %w", err)
			continue
		}
		// Обрабатываем соединение в отдельной горутине
		newClients <- conn
	}
}

// handleNewClients обрабатывает новых клиентов.
func handleNewClients() {
	for {
		conn := <-newClients // Получаем новое соединение клиента
		slog.Info("Новый клиент подключился: ", conn.RemoteAddr())
		go func() {
			defer conn.Close()
			handleConnection(conn) // Обрабатываем соединение клиента
		}()

	}
}

// handleConnection обрабатывает подключения клиента.
func handleConnection(conn net.Conn) {
	// Чтение имени клиента

	nameReader := bufio.NewReader(conn) // Создаем новый буферизированный ридер
	// TODO: Нужно ли закрывать ридеры?
	name, err := nameReader.ReadString('\n') // Читаем строку от клиента
	if err != nil {
		slog.Error("ошибка чтения имени клиента: %w", err)
	}
	name = strings.TrimSpace(name)
	newClientMsg := fmt.Sprintf("Новый пользователь %s подключился к чату", name)
	messages <- newClientMsg
	clients[conn] = name // Вносим имя клиента в список активных соединений

	// Читаем сообщения от клиента по строчно
	reader := bufio.NewReader(conn)
	for {
		getMsg, err := reader.ReadString('\n')
		// Если возникла ошибка, то как правило из-за разрыва соединения
		if err != nil {
			disconnectMsg := fmt.Sprintf("Пользователь %s покинул чат", clients[conn])
			messages <- disconnectMsg
			delete(clients, conn)
			return
		}
		// Обрабатываем полученное сообщение и направляем в очередь сообщений
		msg := fmt.Sprintf("%s: %s", clients[conn], strings.TrimSpace(getMsg))
		messages <- msg
	}
}

// handleMessages обрабатывает сообщения
func handleMessages() {
	for {
		msg := <-messages
		// Для каждого соединения пишем сообщение
		for client := range clients {
			fmt.Fprintln(client, msg)
		}
	}
}
