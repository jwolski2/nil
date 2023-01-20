package server

import "errors"

type storage struct {
	activeChallenges map[authID]*challenge
	users            map[userID]*ys
}

func newStorage() *storage {
	return &storage{
		activeChallenges: map[authID]*challenge{},
		users:            map[userID]*ys{},
	}
}

func (s *storage) hasUser(user string) bool {
	_, ok := s.users[userID(user)]
	return ok
}

func (s *storage) store(user string, ys *ys) error {
	// Check if user already exists.
	if _, ok := s.users[userID(user)]; ok {
		return errors.New("user already exists")
	}

	// Register.
	s.users[userID(user)] = ys

	return nil
}

func (s *storage) storeChallenge(ch *challenge) error {
	// Check if challenge already exists.
	if _, ok := s.activeChallenges[authID(ch.authIDStr())]; ok {
		return errors.New("challenge already exists")
	}

	s.activeChallenges[authID(ch.authIDStr())] = ch

	return nil
}
