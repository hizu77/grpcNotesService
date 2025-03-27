package main

import (
	"net"
	"notes/internal/api"
	"notes/internal/repository"
	"notes/pkg/generated/proto/note"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	r := repository.New()
	noteApi := api.New(r)

	lsn, err := net.Listen("tcp", ":8082")
	if err != nil {
		panic(err)
	}

	noteServer := grpc.NewServer()
	reflection.Register(noteServer)

	note.RegisterNoteManagerServer(noteServer, noteApi)

	if err := noteServer.Serve(lsn); err != nil {
		panic(err)
	}
}
