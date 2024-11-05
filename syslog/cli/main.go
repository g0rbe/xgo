package main

import (
	"fmt"
	"os"

	"github.com/g0rbe/xgo/syslog"
)

func main() {

	addr := "127.0.0.1:514"

	if len(os.Args) > 1 {
		addr = os.Args[1]
	}

	srv, err := syslog.NewServer(addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start server: %s\n", err)
		os.Exit(1)
	}

	for {

		msg, err := srv.Read()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read: %s\n", err)
			break
		}

		fmt.Printf("%s\n", msg)
	}

	srv.Close()
}
