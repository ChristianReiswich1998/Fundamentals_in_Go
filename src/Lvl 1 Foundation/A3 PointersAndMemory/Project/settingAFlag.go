package main

import "fmt"

type Flag2 struct {
	Name    string
	Enabled bool
	Rollout int
}

type Store2 struct {
	Flags []Flag2
}

// Write Set as a method on *Store.
// If a flag with this name exists, update its Enabled field.
// Otherwise append a new Flag with this name and Enabled value.
func (s *Store2) Set(name string, enabled bool) {

	for i := 0; i < len(s.Flags); i++ {
		if s.Flags[i].Name == name {
			s.Flags[i].Enabled = true
			return
		}
	}

	s.Flags = append(s.Flags, Flag2{Name: name, Enabled: enabled})
}

func main() {
	store := &Store2{}
	store.Set("new-checkout", false)
	store.Set("dark-mode", true)
	store.Set("new-checkout", true) // update, not a second flag

	fmt.Println(len(store.Flags)) // 2
	for _, f := range store.Flags {
		fmt.Println(f.Name, f.Enabled)
	}
}
