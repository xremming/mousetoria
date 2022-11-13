package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

	"gameserver/database"
	"gameserver/proto"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

func main() {
	logger := zerolog.New(os.Stderr)

	if len(os.Args) != 2 {
		logger.Fatal().Msg("Usage: mousetoria <grpc-socket-path>")
	}

	ctx := logger.WithContext(context.Background())

	db, err := database.NewDatabase(ctx, "file::memory:")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			logger.Err(err).Msg("failed to close db connection")
		}
	}()

	server := grpc.NewServer()
	proto.RegisterDatabaseServer(server, db)

	addr, err := net.ResolveUnixAddr("unix", os.Args[1])
	if err != nil {
		logger.Err(err).Msg("failed to resolve unix addr")
	}

	conn, err := net.ListenUnix("unix", addr)
	if err != nil {
		logger.Err(err).Msg("failed to listen unix address")
	}
	conn.SetUnlinkOnClose(true)

	go func() {
		err = server.Serve(conn)
		if err != nil {
			logger.Err(err).Msg("failed to serve connections")
		}
	}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	<-exit
	logger.Info().Msg("exiting")
	server.GracefulStop()
}
