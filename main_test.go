package main

import (
	"testing"
)

func TestAddGame(t *testing.T) {
	err := AddGame("test_name", "test_genre")
	if err != nil {
		t.Error(err)
	}
}

func TestViewAll(t *testing.T) {
	err := ViewAll()
	if err != nil {
		t.Error(err)
	}
}

func TestUpdate(t *testing.T) {
	err := Update(1, "test_updated", "test_updated")
	if err != nil {
		t.Error(err)
	}
}
