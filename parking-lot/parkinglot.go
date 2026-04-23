//go:build !solution
// +build !solution

package parkinglot

import (
	"context"
)

type Parking struct {
	ch chan struct{}
}

func NewParking(capacity int) *Parking {
}

func (p *Parking) Park(ctx context.Context) error {
}

func (p *Parking) TryPark() bool {
}

func (p *Parking) Leave() {
}
