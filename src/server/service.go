package server

import (
	"perfumepb"
	"storage"

	"github.com/sirupsen/logrus"

	"golang.org/x/net/context"
)

func (s *server) CreateEntry(ctx context.Context, entry *perfumepb.Entry) (*perfumepb.Entry, error) {
	logrus.Infof("Create Entry: %+v", entry)

	if entry.Link == "" {
		return nil, storage.ErrorInvalidLink
	}

	err := s.storage.AddEntry(entry)
	if err != nil {
		logrus.WithError(err).Errorf("Service can not add entry")
	}

	return entry, nil
}

func (s *server) DeleteEntry(ctx context.Context, entry *perfumepb.Entry) (*perfumepb.Entry, error) {
	logrus.Infof("Delete Entry: %+v", entry)

	if entry.Id == "" {
		return entry, storage.ErrorInvalidId
	}

	err := s.storage.DeleteEntry(entry)

	if err != nil {
		logrus.WithError(err).Errorf("Service can not delete entry")
	}

	return entry, nil
}

func (s *server) UpdateEntry(ctx context.Context, entry *perfumepb.Entry) (*perfumepb.Entry, error) {
	logrus.Infof("Update Entry: %+v", entry)

	if entry.Id == "" {
		return entry, storage.ErrorInvalidId
	}

	err := s.storage.UpdateEntry(entry)

	if err != nil {
		logrus.WithError(err).Errorf("Service can not update entry")
	}

	return entry, nil
}

func (s *server) GetEntryList(ctx context.Context, req *perfumepb.GetEntryListReq) (*perfumepb.EntryListResponse, error) {
	logrus.Infof("Get Entry List:")

	response, err := s.storage.GetEntryList()

	if err != nil {
		logrus.WithError(err).Errorf("Service can not get entry list")
	}

	return response, nil
}

func (s *server) GetEntry(ctx context.Context, req *perfumepb.GetEntryReq) (*perfumepb.Entry, error) {
	logrus.Infof("Get Entry: %+v", req)

	if req.Id == "" {
		return nil, storage.ErrorInvalidId
	}

	entry, err := s.storage.GetEntry(req.Id)

	if err != nil {
		logrus.WithError(err).Errorf("Service can not get entry with id")
	}

	return entry, nil
}
