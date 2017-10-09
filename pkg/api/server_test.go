package api

import "testing"

func TestParseUserIDPositive(t *testing.T) {
	params := make(map[string]string)
	params["userID"] = "123"

	userID, err := parseUserID(params)

	if userID != 123 {
		t.Fatalf("Expected %d, Got %d", 123, userID)
	}

	if err != nil {
		t.Fatalf("Expected no error, got %s", err.Error())
	}
}

func TestParseUserIDNegative(t *testing.T) {
	params := make(map[string]string)

	userID, err := parseUserID(params)

	if userID != 0 {
		t.Fatalf("Expected %d, Got %d", 0, userID)
	}

	if err == nil {
		t.Fatal("Expected no error, got no error")
	}
}
