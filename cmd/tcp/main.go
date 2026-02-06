package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	ls, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer ls.Close()

	for {
		con, err := ls.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("Conexão estabelecida")

		go func(con net.Conn) {
			for {
				data, _ := bufio.NewReader(con).ReadString('\n')
				fmt.Println("Dado recebido:", data)
				if strings.Contains(data, "quit") {
					break
				}
				con.Write([]byte("Sua mensagem foi recebida com sucesso.\n"))
			}
			con.Close()
			fmt.Println("Conexão foi encerrada")
		}(con)
	}
}
