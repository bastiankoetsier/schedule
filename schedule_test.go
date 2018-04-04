package schedule

import (
	"context"
	"testing"
)

func TestShutdown(t *testing.T) {
	s := &Schedule{}
	ctx, cancel := context.WithCancel(context.Background())
	done := s.Start(ctx)
	cancel()
	<-done
}

func TestCommandKeepsTrackOfEntry(t *testing.T) {
	s := &Schedule{}
	f := RunFunc(func(ctx context.Context) {})
	e := s.Command(f)
	if len(s.jobs) != 1 {
		t.Fatalf("expected the len of jobs to be 1, got %d", len(s.jobs))
	}
	if s.jobs[0] != e {
		t.Error("expected the returned entry to be the same pointer")
	}
}

func TestRunFuncCallsWrappedFunc(t *testing.T) {
	called := false
	f := RunFunc(func(ctx context.Context) {
		called = true
	})
	f.Run(context.Background())
	if called != true {
		t.Error("expected the wrapped func to be called")
	}
}
