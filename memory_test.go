package pool

import (
	"context"
	"fmt"
	"testing"
)

func TestMemoryPool(t *testing.T) {

	ctx := context.Background()
	
	uri := fmt.Sprintf("%s://", MEMORY_SCHEME)

	pl, err := NewPool(ctx, uri)

	if err != nil {
		t.Fatalf("Failed to create new pool for %s, %v", uri, err)
	}

	err = pl.Push(ctx, "a")

	if err != nil {
		t.Fatalf("Failed to add 'a', %v", err)
	}

	err = pl.Push(ctx, "b")

	if err != nil {
		t.Fatalf("Failed to add 'b', %v", err)
	}
	
	if pl.Length(ctx) != 2 {
		t.Fatalf("Expected length of 2, got %d", pl.Length(ctx))
	}

	v, ok := pl.Pop(ctx)

	if !ok {
		t.Fatalf("Failed to pop")
	}

	if v.(string) != "b" {
		t.Fatalf("Unexpected value %v (expected %s)", v, "b")
	}

	if pl.Length(ctx) != 1 {
		t.Fatalf("Expected length of 1, got %d", pl.Length(ctx))
	}

	v, ok = pl.Pop(ctx)

	if !ok {
		t.Fatalf("Failed to pop")
	}

	if v.(string) != "a" {
		t.Fatalf("Unexpected value %v (expected %s)", v, "a")
	}

	if pl.Length(ctx) != 0 {
		t.Fatalf("Expected length of 0, got %d", pl.Length(ctx))
	}
	
}
