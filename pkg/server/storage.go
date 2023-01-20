package server

import "errors"

type storage struct {
	users map[string]*ys
}

func newStorage() *storage {
	return &storage{map[string]*ys{}}
}

func (s *storage) store(user string, ys *ys) error {
	// Check if user already exists.
	if _, ok := s.users[user]; ok {
		return errors.New("user already exists")
	}

	// Register.
	s.users[user] = ys

	return nil
}
