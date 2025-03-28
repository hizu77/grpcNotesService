package api

import (
	"context"
	"log"
	"math/rand/v2"
	"notes/internal/model"
	"notes/internal/repository"
	"notes/pkg/generated/proto/note"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NoteApi struct {
	note.UnimplementedNoteManagerServer
	repostitory repository.NoteRepository
}

func New(r repository.NoteRepository) *NoteApi {
	return &NoteApi{
		repostitory: r,
	}
}

func (n *NoteApi) CreateNote(ctx context.Context, request *note.CreateNoteRequest) (*note.CreateNoteResponse, error) {
	noteId := rand.N[int64](1e9)

	log.Print(request.Note)

	err := n.repostitory.CreateNote(ctx, &model.Note{
		ID:     noteId,
		UserID: request.UserId,
		Info:   request.Note.Data,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, errors.Wrap(err, "lol").Error())
	}

	return &note.CreateNoteResponse{
		NoteId: noteId,
	}, nil
}

func (n *NoteApi) GetNoteById(ctx context.Context, request *note.GetNoteByIdRequest) (*note.GetNoteByIdResponse, error) {
	v, err := n.repostitory.GetNoteById(ctx, request.NoteId)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.Wrap(err, "you").Error())
	}

	return &note.GetNoteByIdResponse{
		Note: &note.Note{
			Data: v.Info,
		},
	}, nil
}
