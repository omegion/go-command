package command_test

import (
	"testing"

	"github.com/omegion/go-command"
	"github.com/omegion/go-command/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	dummyCommand = "dummy"
)

type Foo struct {
	Commander command.Interface
}

func (f Foo) Run() error {
	_, err := f.Commander.Output(dummyCommand)
	if err != nil {
		return err
	}

	return nil
}

func TestCommand_Output(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mocks.NewMockInterface(ctrl)

	m.EXPECT().Output(dummyCommand).Return([]byte{}, nil)

	foo := Foo{
		Commander: m,
	}

	err := foo.Run()
	assert.Equal(t, err, nil)
}
