// Package schedule is a small library that helps you specify recurring tasks in a cronjob-like interface (without actually using cron's syntax).
package schedule

import (
	"context"
	"time"

	"github.com/gorhill/cronexpr"
)

// Schedule is the main type that "stores" the jobs to execute
type Schedule struct {
	jobs []*Entry
}

// Runner is the definition of a func to run on a given schedule
type Runner interface {
	Run(ctx context.Context)
}

// RunFunc is a wrapper for Runner to use plain functions instead of implementing a type for Runner
type RunFunc func(ctx context.Context)

// Run satisfies the Runner interface
func (r RunFunc) Run(ctx context.Context) {
	r(ctx)
}

// Command sets a new command in the scheduler. It returns an entry which can be used to set a time schedule.
func (s *Schedule) Command(r Runner) *Entry {
	e := &Entry{expression: "* * * * * *", Command: r}
	s.jobs = append(s.jobs, e)
	return e
}

// Start starts the scheduler to run. It checks every minute for "due" tasks
// to execute them given the same context as the scheduler runs on, to support potential graceful shutdown.
func (s *Schedule) Start(ctx context.Context) <-chan struct{} {
	interval := 1 * time.Minute
	t := time.NewTicker(interval)
	done := make(chan struct{}, 1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				t.Stop()
				done <- struct{}{}
				return
			case <-t.C:
				now := time.Now()
				lastRun := now.Add(-interval)
				for _, e := range s.jobs {
					nextRun := cronexpr.MustParse(e.String()).Next(lastRun)
					if nextRun.Before(now) {
						e.Command.Run(ctx)
						continue
					}
				}
			}
		}
	}()
	return done
}
