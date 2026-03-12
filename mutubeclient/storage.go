package mutubeclient

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

// Store handles the persistent storage and tracking of strings.
type Store struct {
	mu       sync.RWMutex
	filePath string
	seen     map[string]struct{}
}

// NewStore initializes a store, loads existing strings from the file,
// and prepares it for concurrent use.
func NewStore(path string) (*Store, error) {
	s := &Store{
		filePath: path,
		seen:     make(map[string]struct{}),
	}

	// Load existing data so we have history across restarts
	if err := s.load(); err != nil {
		return nil, err
	}

	return s, nil
}

// Save adds a string to the store and persists it to disk if it's new.
func (s *Store) Save(val string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if we've already seen this to avoid duplicates/unnecessary writes
	if _, exists := s.seen[val]; exists {
		return nil
	}

	// Persist to disk (Append mode)
	f, err := os.OpenFile(s.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer f.Close()

	if _, err := f.WriteString(val + "\n"); err != nil {
		return fmt.Errorf("could not write string: %w", err)
	}

	// Update in-memory cache
	s.seen[val] = struct{}{}
	return nil
}

// Exists checks if a string has been stored previously.
func (s *Store) Exists(val string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exists := s.seen[val]
	return exists
}

// load reads the file from disk into the in-memory map.
func (s *Store) load() error {
	f, err := os.Open(s.filePath)
	if os.IsNotExist(err) {
		return nil // New store, nothing to load
	}
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s.seen[scanner.Text()] = struct{}{}
	}
	return scanner.Err()
}
