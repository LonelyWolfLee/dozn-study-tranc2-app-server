package transaction

import (
	"fmt"
	"net"
	"time"
)

func Client() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Faield to Dial : ", err)
	}
	fmt.Println("Connect Server !")

	defer conn.Close()

	go func(c net.Conn) {
		send := "Transaction"
		for {
			_, err = c.Write([]byte(send))
			if err != nil {
				fmt.Println("Failed to write data : ", err)
				break
			}

			time.Sleep(1 * time.Second)
		}
	}(conn)

	go func(c net.Conn) {
		recv := make([]byte, 4096)

		for {
			n, err := c.Read(recv)
			if err != nil {
				fmt.Println("Failed to Read data : ", err)
				break
			}

			fmt.Println(string(recv[:n]))
		}
	}(conn)

	fmt.Scanln()
}
