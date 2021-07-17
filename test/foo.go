package test

import (
	"io"
	"time"

	"github.com/google/uuid"
)

//go:generate getter -type Foo,Bar

type Foo struct {
	A string
	B int64
	C *string
	D *int
	E uuid.UUID
	F chan *Bar
	G Bar
	H <-chan int
	J map[string]*string
	K map[int]Bar
	M map[int64]uuid.UUID
	N map[Bar]map[uuid.UUID]string
	O time.Time
	Q io.Writer
	R func(string) uuid.UUID
	S func(string, int, map[string]uuid.UUID) chan map[string]io.WriteCloser
}

type Bar struct {
	S string
}
