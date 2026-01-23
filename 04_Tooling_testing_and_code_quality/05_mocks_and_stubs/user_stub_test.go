package users

import (
	"fmt"
	"testing"
)

// stubStore provides predefined responses for testing.
type stubStore struct {
	users map[int]User
}

func (s *stubStore) GetUser(id int) (User, error) {
	if u, ok := s.users[id]; ok {
		return u, nil
	}
	return User{}, fmt.Errorf("not found")
}

func (s *stubStore) SaveUser(u User) error {
	s.users[u.ID] = u
	return nil
}

func TestGreetUser_WithStub(t *testing.T) {
	s := &stubStore{users: map[int]User{1: {ID: 1, Name: "Alice"}}}
	got, err := GreetUser(s, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "Hello, Alice" {
		t.Fatalf("unexpected greeting: %q", got)
	}
}
