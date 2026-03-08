package sdk

import (
	"testing"
	"time"
)

func TestDefaultConfig(t *testing.T) {
	c := NewClient("https://api.com")

	if c.timeout != 30*time.Second {
		t.Fatalf("expected timeout 30s, got %v", c.timeout)
	}

	if c.retries != 3 {
		t.Fatalf("expected retries 3, got %d", c.retries)
	}

	if c.debug != false {
		t.Fatalf("expected debug false, got %v", c.debug)
	}
}

func TestCustomConfig(t *testing.T) {
	c := NewClient(
		"https://api.com",
		WithTimeout(5*time.Second),
		WithRetries(10),
		WithDebug(true),
	)

	if c.timeout != 5*time.Second {
		t.Fatalf("expected timeout 5s, got %v", c.timeout)
	}

	if c.retries != 10 {
		t.Fatalf("expected retries 10, got %d", c.retries)
	}

	if c.debug != true {
		t.Fatalf("expected debug true, got %v", c.debug)
	}
}

func TestOptionApplicationOrder(t *testing.T) {
	c := NewClient(
		"https://api.com",
		WithTimeout(10*time.Second),
		WithTimeout(5*time.Second),
	)

	if c.timeout != 5*time.Second {
		t.Fatalf("expected timeout 5s because last option wins, got %v", c.timeout)
	}
}
