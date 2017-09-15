package server

import (
	"perfumepb"

	"golang.org/x/net/context"
)

func (s *server) CreateEntry(context.Context, *perfumepb.Entry) (*perfumepb.Entry, error) {

}
func (s *server) DeleteEntry(context.Context, *perfumepb.Entry) (*perfumepb.Entry, error) {

}
func (s *server) UpdateEntry(context.Context, *perfumepb.Entry) (*perfumepb.Entry, error) {

}
func (s *server) GetEntryList(context.Context, *perfumepb.GetEntryListReq) (*perfumepb.EntryListResponse, error) {

}
func (s *server) GetEntry(context.Context, *perfumepb.GetEntryReq) (*perfumepb.Entry, error) {

}
