package main

import (
	"net/url"
	"testing"
)

func TestGetFiltersFromKey(t *testing.T) {
	// Test for teamKey
	t.Run("team key filters", func(t *testing.T) {
		teamCache := &teamCache{}
		key := teamKey{
			name:          "Mumbai Indians",
			is_male:       true,
			playing_level: "domestic",
		}

		values := teamCache.getFiltersFromKey(key)

		expected := url.Values{
			"name":          []string{"Mumbai Indians"},
			"is_male":       []string{"true"},
			"playing_level": []string{"domestic"},
			"__page":        []string{"1"},
			"__limit":       []string{"1"},
		}

		// Check if all expected keys exist with correct values
		for k, expectedVal := range expected {
			if gotVal := values.Get(k); gotVal != expectedVal[0] {
				t.Errorf("for key %s, expected %s but got %s", k, expectedVal[0], gotVal)
			}
		}

		// Check if there aren't any extra keys
		if len(values) != len(expected) {
			t.Errorf("expected %d keys but got %d", len(expected), len(values))
		}
	})

	// Test for playerKey
	t.Run("player key filters", func(t *testing.T) {
		playerCache := &playerCache{}
		key := playerKey{
			cricsheet_id: "PLAYER123",
		}

		values := playerCache.getFiltersFromKey(key)

		expected := url.Values{
			"cricsheet_id": []string{"PLAYER123"},
			"__page":       []string{"1"},
			"__limit":      []string{"1"},
		}

		// Check if the expected key exists with correct value
		if got := values.Get("cricsheet_id"); got != expected.Get("cricsheet_id") {
			t.Errorf("expected cricsheet_id=PLAYER123 but got %s", got)
		}

		// Check if there aren't any extra keys
		if len(values) != len(expected) {
			t.Errorf("expected 1 key but got %d", len(values))
		}
	})

	// Test empty values
	t.Run("empty values", func(t *testing.T) {
		teamCache := &teamCache{}
		key := teamKey{
			name:          "",
			is_male:       false,
			playing_level: "",
		}

		values := teamCache.getFiltersFromKey(key)

		// Check if is_male is included even though it's false
		if got := values.Get("is_male"); got != "false" {
			t.Errorf("expected is_male=false but got %s", got)
		}

		// Check if empty strings are excluded
		if got := values.Get("name"); got != "" {
			t.Errorf("expected empty name to be excluded but got %s", got)
		}
		if got := values.Get("playing_level"); got != "" {
			t.Errorf("expected empty playing_level to be excluded but got %s", got)
		}
	})
}
