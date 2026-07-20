package main

import (
	"context"
)

type Job struct {
	ID int
}

type Result struct {
	JobID int
	Value string
}

type WorkerPool struct {
	// your fields
}

func NewWorkerPool(workerCount int) *WorkerPool

func (p *WorkerPool) Submit(job Job) error

func (p *WorkerPool) Results() <-chan Result

func (p *WorkerPool) Start(ctx context.Context)

func (p *WorkerPool) Stop()
