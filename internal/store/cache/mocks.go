package cache

import (
	"context"

	"github.com/karthikbhandary2/Social/internal/store"
)

func NewMockStore() Storage {
	return Storage{
		Users: &MockUserStore{},
	}
}

type MockUserStore struct {
	
}

func (m *MockUserStore) Get(ctx context.Context, userID int64) (*store.User, error) {
	// args := m.Called(userID)
	// return nil, args.Error(1)
	return nil, nil
}

func (m *MockUserStore) Set(ctx context.Context, user *store.User) error {
	// args := m.Called(user)
	// return args.Error(0)
	return nil
}

func (m *MockUserStore) Delete(ctx context.Context, userID int64) {
	// m.Called(userID)
}