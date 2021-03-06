# go-command

[![GithubBuild](https://img.shields.io/github/workflow/status/omegion/go-command/Code%20Check)](http://pkg.go.dev/github.com/omegion/go-command)
[![Coverage Status](https://coveralls.io/repos/github/omegion/go-command/badge.svg?branch=master)](https://coveralls.io/github/omegion/go-command?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/omegion/go-command)](https://goreportcard.com/report/github.com/omegion/go-command)
[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/omegion/go-command)

`exec.Command` wrapper for better testing.

## How does it work?

```go
const dummyCommand = "dummy"

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
```
