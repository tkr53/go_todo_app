package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tkr53/go_todo_app/config"
)

func TestNewMux(t *testing.T) {
	t.Skip("スキップする")
	cfg, _ := config.New()
	ctx, cancel := context.WithCancel(context.Background())
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	sut, creanup, err := NewMux(ctx, cfg)
	if err != nil {
		t.Errorf("failed create mux: %v", err)
	}
	defer creanup()
	sut.ServeHTTP(w, r)
	resp := w.Result()
	t.Cleanup(func() { _ = resp.Body.Close() })

	if resp.StatusCode != http.StatusOK {
		t.Error("want status code 200, but", resp.StatusCode)
	}
	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	want := `{"status": "ok"}`
	if string(got) != want {
		t.Errorf("want %q, but goy %q", want, got)
	}
	cancel()
}
