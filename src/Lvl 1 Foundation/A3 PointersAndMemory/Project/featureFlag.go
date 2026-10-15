package main

import "fmt"

// Define the Flag type.
// A flag has a Name (string), an Enabled switch (bool), and a Rollout percentage (int).
type Flag struct {
	Name    string
	Enabled bool
	Rollout int
}

// Define the Store type.
// It owns a slice of Flags.
type Store struct {
	Flags []Flag
}

// loadFlags returns the flag set. In a real service this would read from a
// database; here the rows are hardcoded so the project stays runnable.
func loadFlags() []Flag {
	return []Flag{
		{Name: "new-checkout", Enabled: true, Rollout: 25},
		{Name: "dark-mode", Enabled: false, Rollout: 0},
	}
}

func main() {
	store := Store{Flags: loadFlags()}

	for _, f := range store.Flags {
		fmt.Println(f.Name, f.Enabled, f.Rollout)
	}
}
