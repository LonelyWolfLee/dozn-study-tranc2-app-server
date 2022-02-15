package transaction

import (
	"dozn/app-server/logging"
	"fmt"
	"net"
)

func client(request string) {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		logging.Error(fmt.Sprintf("Faield to Dial : %v", err))
	}
	logging.Info("Connect Server !")

	defer conn.Close()

	go func(c net.Conn) {
		send := request
		for {
			_, err = c.Write([]byte(send))
			if err != nil {
				logging.Error(fmt.Sprintf("Failed to write data : %v", err))
				break
			}
			break
		}
	}(conn)

	go func(c net.Conn) {
		recv := make([]byte, 4096)

		for {
			n, err := c.Read(recv)
			if err != nil {
				logging.Error(fmt.Sprintf("Failed to Read data : %v", err))
				fmt.Println()
				break
			}
			logging.Info(fmt.Sprintf("Receive data : %s", string(recv[:n])))
			break
		}
	}(conn)
}
