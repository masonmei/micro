// Package sync is an interface for distributed synchronization
package sync

import (
	"context"
	"crypto/tls"
	"errors"
	"time"

	"github.com/smart-echo/micro/logger"
)

var (
	ErrLockTimeout = errors.New("lock timeout")
)

// Sync is an interface for distributed synchronization.
type Sync interface {
	// Initialize options
	Init(...Option) error
	// Return the options
	Options() Options
	// Elect a leader
	Leader(id string, opts ...LeaderOption) (Leader, error)
	// Lock acquires a lock
	Lock(id string, opts ...LockOption) error
	// Unlock releases a lock
	Unlock(id string) error
	// Sync implementation
	String() string
}

// Leader provides leadership election.
type Leader interface {
	// resign leadership
	Resign() error
	// status returns when leadership is lost
	Status() chan bool
}

type Options struct {
	Context   context.Context
	Logger    logger.Logger
	TLSConfig *tls.Config
	Prefix    string
	Nodes     []string
}

type Option func(o *Options)

type LeaderOptions struct{}

type LeaderOption func(o *LeaderOptions)

type LockOptions struct {
	TTL  time.Duration
	Wait time.Duration
}

type LockOption func(o *LockOptions)
