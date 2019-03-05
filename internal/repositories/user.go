package repositories

import (
	"ects/internal/models"
	"errors"
	"sync"
)

type Query func(user models.User) bool

type UserRepository interface {
	Exec(query Query, action Query, limit int, mode int) (ok bool)
	Select(query Query) (user models.User, found bool)
	SelectMany(query Query, limit int) (results []models.User)
	InsertOrUpdate(user models.User) (updated models.User, err error)
	Delete(query Query, limit int) (deleted bool)
}

type userMemoryRepository struct {
	source map[int64]models.User
	mu     sync.RWMutex
}

func NewUserRepository(source map[int64]models.User) UserRepository {
	return &userMemoryRepository{source: source}
}

const (
	// ReadOnlyMode will RLock(read) the data .
	ReadOnlyMode = iota
	// ReadWriteMode will Lock(read/write) the data.
	ReadWriteMode
)

func (r *userMemoryRepository) Exec(query Query, action Query, limit int, mode int) (ok bool) {
	loops := 0

	if mode == ReadOnlyMode {
		r.mu.RLock()
		defer r.mu.RUnlock()
	} else {
		r.mu.Lock()
		defer r.mu.Unlock()
	}

	for _, movie := range r.source {
		ok = query(movie)
		if ok {
			if action(movie) {
				loops++
				if limit >= loops {
					break // break
				}
			}
		}
	}

	return
}

func (r *userMemoryRepository) Select(query Query) (user models.User, found bool) {
	found = r.Exec(query, func(m models.User) bool {
		user = m
		return true
	}, 1, ReadOnlyMode)

	// set an empty models.User if not found at all.
	if !found {
		user = models.User{}
	}

	return
}


func (r *userMemoryRepository) SelectMany(query Query, limit int) (results []models.User) {
	r.Exec(query, func(m models.User) bool {
		results = append(results, m)
		return true
	}, limit, ReadOnlyMode)

	return
}

func (r *userMemoryRepository) InsertOrUpdate(user models.User) (models.User, error) {
	id := user.ID

	if id == 0 { // Create new action
		var lastID int64
		// find the biggest ID in order to not have duplications
		// in productions apps you can use a third-party
		// library to generate a UUID as string.
		r.mu.RLock()
		for _, item := range r.source {
			if item.ID > lastID {
				lastID = item.ID
			}
		}
		r.mu.RUnlock()

		id = lastID + 1
		user.ID = id

		// map-specific thing
		r.mu.Lock()
		r.source[id] = user
		r.mu.Unlock()

		return user, nil
	}

	// Update action based on the user.ID,
	// here we will allow updating the poster and genre if not empty.
	// Alternatively we could do pure replace instead:
	// r.source[id] = user
	// and comment the code below;
	current, exists := r.Select(func(m models.User) bool {
		return m.ID == id
	})

	if !exists { // ID is not a real one, return an error.
		return models.User{}, errors.New("failed to update a nonexistent user")
	}

	// or comment these and r.source[id] = m for pure replace
	if user.Avatar != "" {
		current.Avatar = user.Avatar
	}

	if user.Password != "" {
		current.Password = user.Password
	}

	// map-specific thing
	r.mu.Lock()
	r.source[id] = current
	r.mu.Unlock()

	return user, nil
}

func (r *userMemoryRepository) Delete(query Query, limit int) bool {
	return r.Exec(query, func(m models.User) bool {
		delete(r.source, m.ID)
		return true
	}, limit, ReadWriteMode)
}
