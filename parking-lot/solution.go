//go:build solution
// +build solution

package parkinglot

import (
	"context"
	"time"
)

type Parking struct {
	ch       chan struct{}
	capacity int
}

func NewParking(capacity int) *Parking {
	ch := make(chan struct{}, capacity)
	for range capacity {
		ch <- struct{}{}
	}
	return &Parking{
		capacity: capacity,
		ch:       ch,
	}
}

func (p *Parking) Park(ctx context.Context) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(3*time.Second))
	defer cancel()
	select {
	case <-ctxTimeout.Done():
		return ctxTimeout.Err()
	case <-ctx.Done():
		return ctx.Err()
	case <-p.ch:
		return nil
	}
}

func (p *Parking) TryPark() bool {
	select {
	case <-p.ch:
		return true
	default:
		return false
	}
}

func (p *Parking) Leave() {
	select {
	case p.ch <- struct{}{}:
	default:
	}
}
