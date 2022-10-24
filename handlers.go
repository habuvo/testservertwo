package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/habuvo/testserverone/proto"
	"google.golang.org/grpc"
)

type handler struct {
	client proto.PersonServiceClient
}

func NewHandler(c *grpc.ClientConn) *handler {
	return &handler{proto.NewPersonServiceClient(c)}
}

func (h *handler) getPerson(w http.ResponseWriter, r *http.Request) {
	person, err := h.client.GetPerson(r.Context(), &proto.GetRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.WriteString(w, fmt.Sprintf("got person with name %s\n", person.Name))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func (h *handler) setPerson(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Unimplemented\n")
}
