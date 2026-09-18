package spinner

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type Spinner struct {
	ctx       context.Context
	cancel    context.CancelFunc
	label     string
	frames    []string
	interval  time.Duration
	out       io.Writer
	done      chan struct{}
	started   chan struct{}
	startOnce sync.Once
	stopOnce  sync.Once
}

type Option func(*Spinner)

func New(ctx context.Context, label string, opts ...Option) *Spinner {
	ctx, cancel := context.WithCancel(ctx)

	s := &Spinner{
		ctx:    ctx,
		cancel: cancel,
		label:  label,
		frames: []string{
			//"|", "/", "-", "\\"
			"∙∙∙",
			"●∙∙",
			"∙●∙",
			"∙∙●",
			"∙∙∙",
		},
		interval: 125 * time.Millisecond,
		out:      os.Stdout,
		done:     make(chan struct{}),
		started:  make(chan struct{}),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func WithInterval(interval time.Duration) Option {
	return func(s *Spinner) {
		if interval > 0 {
			s.interval = interval
		}
	}
}

func WithFrames(frames []string) Option {
	return func(s *Spinner) {
		if len(frames) > 0 {
			s.frames = append([]string(nil), frames...)
		}
	}
}

func WithOutput(out io.Writer) Option {
	return func(s *Spinner) {
		if out != nil {
			s.out = out
		}
	}
}

func (s *Spinner) Start() {
	s.startOnce.Do(func() {
		close(s.started)
		go s.run()
	})
}

func (s *Spinner) Stop() {
	s.stopOnce.Do(func() {
		s.cancel()

		select {
		case <-s.started:
			<-s.done
		default:
		}

		// Clear the spinner line and move to the next line
		fmt.Fprintf(s.out, "\r\033[2K")
	})
}

func (s *Spinner) run() {
	defer close(s.done)

	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	i := 0
	for {
		fmt.Fprintf(s.out, "\r%s %s", s.label, s.frames[i])
		i = (i + 1) % len(s.frames)

		select {
		case <-s.ctx.Done():
			return
		case <-ticker.C:
		}
	}
}

func Run(ctx context.Context, label string, fn func(context.Context) error, opts ...Option) error {
	s := New(ctx, label, opts...)
	s.Start()
	defer s.Stop()

	return fn(ctx)
}
