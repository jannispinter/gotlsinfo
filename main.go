package main

import (
	"os"
	"crypto/tls"
	"fmt"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <hostname> <port>\n", os.Args[0])
		os.Exit(1)
	}

	done := make(chan bool)

	var SERVER_NAME string = os.Args[1]
	var SERVER_PORT string = os.Args[2]

	var cipherMap = *GetCipherMap()

	for cipher, cipherName := range cipherMap {
		go func(cipher uint16, cipherName string) {
			conn, err := tls.Dial("tcp", SERVER_NAME+":"+SERVER_PORT, &tls.Config{
				CipherSuites: []uint16{cipher},
			})
			if err != nil {
				// TODO: Could also fail due to various other reasons
				fmt.Printf("Cipher: %s is NOT supported.\n", cipherName)
			} else {
				fmt.Printf("Cipher: %s is supported.\n", cipherName)
				conn.Close()
			}
			done <- true
		}(cipher, cipherName)
	}

	// wait for all goroutines to finish
	for i := 0; i < len(cipherMap); i++ {
		<-done
	}

}
