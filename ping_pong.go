package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for {
		c <- "ping"             // envia a palavra "ping" para o canal
		time.Sleep(time.Second) // aguarda 1 segundo
	}
}

func pong(c chan string) {
	for {
		c <- "pong"             // envia a palavra "pong" para o canal
		time.Sleep(time.Second) // aguarda 1 segundo
	}
}

func main() {
	c := make(chan string) // cria um canal de strings

	go ping(c) // inicia a goroutine para a função ping
	go pong(c) // inicia a goroutine para a função pong

	for {
		msg := <-c // recebe a mensagem do canal

		fmt.Println(msg) // exibe a mensagem recebida

		time.Sleep(time.Second) // aguarda 1 segundo antes de exibir a próxima mensagem
	}
}
