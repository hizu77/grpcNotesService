package repository

import (
	"context"
	"errors"
	"notes/internal/model"
	"sync"
)

var (
	ErrNoteAlreadyExists = errors.New("note with this ID is already exists")
	ErrNoteDoesntExist   = errors.New("note with this ID doesnt exist")
)

type NoteRepository interface {
	CreateNote(ctx context.Context, note *model.Note) error
	GetNoteById(ctx context.Context, id int64) (*model.Note, error)
}

type noteRepositoryImpl struct {
	mx    *sync.RWMutex
	notes map[int64]*model.Note
}

func (n *noteRepositoryImpl) CreateNote(ctx context.Context, note *model.Note) error {
	n.mx.Lock()
	defer n.mx.Unlock()

	if _, ok := n.notes[note.ID]; !ok {
		n.notes[note.ID] = note
		return nil
	} else {
		return ErrNoteAlreadyExists
	}
}

func (n *noteRepositoryImpl) GetNoteById(ctx context.Context, id int64) (*model.Note, error) {
	n.mx.RLock()
	defer n.mx.RUnlock()

	if v, ok := n.notes[id]; !ok {
		return nil, ErrNoteDoesntExist
	} else {
		return v, nil
	}
}

func New() *noteRepositoryImpl {
	return &noteRepositoryImpl{
		mx:    &sync.RWMutex{},
		notes: make(map[int64]*model.Note),
	}
}
