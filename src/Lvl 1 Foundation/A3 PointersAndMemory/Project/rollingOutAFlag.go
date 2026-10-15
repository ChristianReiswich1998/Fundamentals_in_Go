package main

import (
	"fmt"
	"slices"
)

type Flag4 struct {
	Name    string
	Enabled bool
	Rollout int
}

type Store4 struct {
	Flags []Flag4
}

// Write IsEnabled: return whether the named flag is on for this user.
// Unknown flag or disabled flag means off. Otherwise roll out by bucket.
func (s *Store4) IsEnabled(name string, userID int) bool {
	for i := 0; i < len(s.Flags); i++ {
		if s.Flags[i].Name == name {
			if s.Flags[i].Enabled == false {
				return false
			} else if userID%100 < s.Flags[i].Rollout {
				return true
			} else if s.Flags[i].Rollout == 100 {
				for j := 0; j < len(s.Flags); j++ {
					s.Flags[j].Enabled = true
				}
			}
		}
	}
	return false
}

// --- Store methods from earlier lessons. No need to change these. ---

func (s *Store4) Set(name string, enabled bool) {
	for i := range s.Flags {
		if s.Flags[i].Name == name {
			s.Flags[i].Enabled = enabled
			return
		}
	}
	s.Flags = append(s.Flags, Flag4{Name: name, Enabled: enabled})
}

func (s *Store4) Find(name string) *Flag4 {
	for i := range s.Flags {
		if s.Flags[i].Name == name {
			return &s.Flags[i]
		}
	}
	return nil
}

func (s *Store4) Remove(name string) {
	for i := range s.Flags {
		if s.Flags[i].Name == name {
			s.Flags = slices.Delete(s.Flags, i, i+1)
			return
		}
	}
}

func main() {
	store := &Store4{}
	store.Set("new-checkout", true)
	if f := store.Find("new-checkout"); f != nil {
		f.Rollout = 30
	}

	for _, id := range []int{7, 42, 95} {
		fmt.Println(id, store.IsEnabled("new-checkout", id))
	}
}
