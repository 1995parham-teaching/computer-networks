/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 26-03-2020
 * |
 * | File Name:     echo-server.go
 * +===============================================
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	// here we use 127.0.0.1 as our listening address and 1378 as our listening port
	// change 127.0.0.1 to 0.0.0.0 and describe what happens
	l, err := net.Listen("tcp", "127.0.0.1:1378")
	if err != nil {
		log.Fatalf("listen failed %s", err)
	}

	// here we have a listener, actually a TCP listener
	// infinitely looping for accepting new connection
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("accept failed: %s", err)
		}

		// conn is new client that must be handled
		// when we are handling conn others must wait
		log.Printf("Hello %s", conn.RemoteAddr())

		// we can use the low-level functions to read from the sockets
		// but as you saw in java you can wrap-up the streams to have streams with more functions
		rd := bufio.NewReader(conn)
		wr := bufio.NewWriter(conn)

		for {
			s, err := rd.ReadString('\n')
			if err != nil {
				log.Printf("read from %s failed: %s", conn.RemoteAddr(), err)
				break
			}

			// let client to gracefully close the connection with "BYE" command
			if strings.Compare(s, "BYE\r\n") == 0 {
				break
			}

			fmt.Printf("client (%s): %s", conn.RemoteAddr(), s)

			if _, err := wr.WriteString(s); err != nil {
				log.Printf("read from %s failed: %s", conn.RemoteAddr(), err)
				break
			}

			// IMPORTANT
			wr.Flush()
		}

		// close connection and wait for another connection
		conn.Close()
	}
}
