package users

import (
	"errors"
	"testing"
)

// mockStore records calls for verification.
type mockStore struct {
	getCalls  []int
	saveCalls []User

	// configure return values
	getReturn User
	getErr    error
}

func (m *mockStore) GetUser(id int) (User, error) {
	m.getCalls = append(m.getCalls, id)
	if m.getErr != nil {
		return User{}, m.getErr
	}
	return m.getReturn, nil
}

func (m *mockStore) SaveUser(u User) error {
	m.saveCalls = append(m.saveCalls, u)
	return nil
}

func TestGreetUser_VerifiesCall(t *testing.T) {
	m := &mockStore{getReturn: User{ID: 2, Name: "Bob"}}
	g, err := GreetUser(m, 2)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if g != "Hello, Bob" {
		t.Fatalf("greet = %q", g)
	}
	if len(m.getCalls) != 1 || m.getCalls[0] != 2 {
		t.Fatalf("expected GetUser called once with 2, got %v", m.getCalls)
	}
}

func TestGreetUser_ErrorPropagates(t *testing.T) {
	m := &mockStore{getErr: errors.New("db down")}
	_, err := GreetUser(m, 3)
	if err == nil {
		t.Fatalf("expected error")
	}
}
