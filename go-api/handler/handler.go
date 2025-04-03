package handler

import (
	"github.com/s-union/canalia/generated"
)

var _ generated.ServerInterface = (*Server)(nil)

type Server struct{}

func NewServer() Server {
	return Server{}
}
