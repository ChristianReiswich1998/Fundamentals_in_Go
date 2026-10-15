package main

import (
	"fmt"
	"slices"
)

type Flag3 struct {
	Name    string
	Enabled bool
	Rollout int
}

type Store3 struct {
	Flags []Flag3
}

// Write Find: return *Flag for the matching name, or nil.
func (s *Store3) Find(name string) *Flag3 {
	for i := 0; i < len(s.Flags); i++ {
		if s.Flags[i].Name != name {
			return nil
		}
	}

	return &Flag3{}
}

// Write Remove: delete the first flag with the matching name.
func (s *Store3) Remove(name string) {
	for i := 0; i < len(s.Flags); i++ {
		if s.Flags[i].Name == name {
			s.Flags = slices.Delete(s.Flags, i, i+1)
			return
		}
	}
}

// Set is already provided.
func (s *Store3) Set(name string, enabled bool) {
	for i := 0; i < len(s.Flags); i++ {
		if s.Flags[i].Name == name {
			s.Flags[i].Enabled = enabled
			return
		}
	}
	s.Flags = append(s.Flags, Flag3{Name: name, Enabled: enabled})
}

func main() {
	store := &Store3{}
	store.Set("new-checkout", true)
	store.Set("dark-mode", false)
	store.Set("legacy-banner", true)

	if f := store.Find("new-checkout"); f != nil {
		f.Rollout = 50
	}
	store.Remove("legacy-banner")

	for _, f := range store.Flags {
		fmt.Println(f.Name, f.Enabled, f.Rollout)
	}
}
