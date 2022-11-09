package main

import (
	"context"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"

	"github.com/rs/zerolog"
	"github.com/xremming/mousetoria/db"
	"github.com/xremming/mousetoria/ledger"
)

type Database struct{ db.Database }

func (d *Database) InsertTransaction(tx ledger.Transaction, value *int64) (err error) {
	*value, err = d.Database.InsertTransaction(context.TODO(), tx)
	log.Printf("InsertTransaction(%v) = %v", tx, *value)
	return
}

func main() {
	logger := zerolog.New(os.Stderr)
	ctx := logger.WithContext(context.Background())

	db, err := db.NewDatabase(ctx, "file::memory:")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}()

	addr, err := net.ResolveUnixAddr("unix", "/tmp/mousetoria.sock")
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenUnix("unix", addr)
	if err != nil {
		log.Fatal(err)
	}

	listener := Database{db}
	rpc.Register(&listener)

	for {
		conn, err := inbound.Accept()
		if err != nil {
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
