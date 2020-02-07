package main

import (
	"errors"
)

type Client struct {
	addr string
}
type Server struct {
	numConnClients int
	Clients        []Client
	maxClients     int
}

var NotFound = errors.New("Not Found")

func (s *Server) addClient(c Client) {
	if s.numConnClients == s.maxClients {
		return
	}
	s.Clients = append(s.Clients, c)
	s.numConnClients++
}

func (s *Server) getClient(addr string) (Client, error) {
	var c Client
	for _, c := range s.Clients {
		if c.addr == addr {
			return c, nil
		}
	}
	return c, NotFound
}
