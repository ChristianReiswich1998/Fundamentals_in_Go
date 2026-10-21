package main

import (
	"errors"
	"fmt"
	"strings"
)

// --- Client Registry ---
func getClientList() []string {
	// TODO: Return a slice with the clients
	var clients []string
	clients = append(clients, "SVC-AUTH")
	clients = append(clients, "SVC-USERS")
	clients = append(clients, "SVC-BILLING")
	return clients
}

func assignTier(clients []string, index int) string {
	// TODO: Return clients[index % len(clients)]
	return clients[index%len(clients)]
}

// --- Input Cleanup
// normalizeClientID trims whitespace and converts to uppercase.
func normalizeClientID(id string) string {
	// TODO: Implement
	var result string
	result = strings.TrimSpace(id)
	result = strings.ToUpper(result)
	return result
}

// isValidClientID returns true if the ID is non-empty and within maxLength.
func isValidClientID(id string, maxLength int) bool {
	// TODO: Implement
	if len(id) > 0 && len(id) <= maxLength {
		return true
	}

	return false
}

// --- Request Gatekeeper
// isRegisteredClient returns true if clientID is found in the clients slice.
func isRegisteredClient(clientID string, clients []string) bool {
	// TODO: Implement
	for i := 0; i < len(clients); i++ {
		if clientID == clients[i] {
			return true
		}
	}

	return false
}

// validateRequest checks format first, then registration.
// Returns nil if valid, or an error describing the problem.
func validateRequest(clientID string, maxLength int, clients []string) error {
	// TODO: Implement
	if isValidClientID(clientID, maxLength) == false {
		return errors.New("invalid client ID format")
	} else if isRegisteredClient(clientID, clients) == false {
		return errors.New("unregistered client")
	}

	return nil
}

// RequestResult holds the rate limit decision for a single request.
type RequestResult struct {
	ClientID string
	Status   string // "allowed" or "denied"
}

// checkRequests processes a batch of requests against a rate limit.
// priorCounts holds existing counts from earlier in the window.
func checkRequests(requests []string, priorCounts map[string]int, limit int) []RequestResult {
	// TODO: Implement
	//counts := make(map[string]int, len(priorCounts))

	return nil
}

func main() {
	// Test Client Registry
	fmt.Println("Test: Client Registry")
	clients := getClientList()
	fmt.Printf("Registry: %d clients\n", len(clients))
	for i := 0; i < 3; i++ {
		fmt.Printf("Tier %d: %s\n", i, assignTier(clients, i))
	}
	fmt.Printf("Wrapped: %s\n", assignTier(clients, len(clients)+2))

	// Test Input Cleanup
	fmt.Println("\nTest: Input Cleanup")
	// Test normalization
	fmt.Printf("Normalize '  svc-auth  ': %q\n", normalizeClientID("  svc-auth  "))
	fmt.Printf("Normalize 'Svc-Users': %q\n", normalizeClientID("Svc-Users"))

	// Test validation
	fmt.Printf("Valid 'SVC-AUTH' (20): %v\n", isValidClientID("SVC-AUTH", 20))
	fmt.Printf("Valid '' (20): %v\n", isValidClientID("", 20))
	fmt.Printf("Valid 'TOOLONGCLIENTIDTHATEXCEEDSLIMIT' (20): %v\n", isValidClientID("TOOLONGCLIENTIDTHATEXCEEDSLIMIT", 20))

	//Test Request Gatekeeper
	fmt.Println("Test: Request Gatekeeper")
	clients2 := getClientList()

	fmt.Printf("SVC-AUTH registered: %v\n", isRegisteredClient("SVC-AUTH", clients2))
	fmt.Printf("SVC-UNKNOWN registered: %v\n", isRegisteredClient("SVC-UNKNOWN", clients2))

	fmt.Printf("SVC-AUTH: %v\n", validateRequest("SVC-AUTH", 20, clients2))
	fmt.Printf("empty: %v\n", validateRequest("", 20, clients2))
	fmt.Printf("SVC-UNKNOWN: %v\n", validateRequest("SVC-UNKNOWN", 20, clients2))

	//Test Counting Requests
	fmt.Println("Test: Counting Requests")
	priorCounts := map[string]int{"SVC-AUTH": 2}
	requests := []string{"SVC-AUTH", "SVC-USERS", "SVC-AUTH", "SVC-USERS", "SVC-USERS"}
	results := checkRequests(requests, priorCounts, 3)
	for _, r := range results {
		fmt.Printf("[%s:%s] ", r.ClientID, r.Status)
	}
	fmt.Println()

}
