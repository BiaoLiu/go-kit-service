package userprofile_service

import (
	"errors"
	"context"
)

type ProfileService interface {
	PostProfile(ctx context.Context, p Profile) error
	GetProfile(ctx context.Context, id string) (Profile, error)
}

type profileService struct {
	profiles map[string]Profile // in-memory data store
}

func (ps *profileService) PostProfile(ctx context.Context, p Profile) error {
	if _, ok := ps.profiles[p.ID]; ok {
		return errors.New("profile already exists")
	}
	ps.profiles[p.ID] = p
	return nil
}

func (ps *profileService) GetProfile(ctx context.Context, id string) (Profile, error) {
	p, ok := ps.profiles[id]
	if !ok {
		return Profile{}, errors.New("profile not found")
	}
	return p, nil
}

type Profile struct {
	ID   string
	Name string
}
