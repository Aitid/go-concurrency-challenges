package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type ExternalAPI interface {
	Send(batch []Event) error
}

type Event struct {
	Name      string
	Timestamp time.Time
	Payload   map[string]string
}

type AnalyticsCollector struct {
	numWorks  int
	interval  time.Duration
	sizeBatch int
	eventCh   chan Event
	jobCh     chan []Event
	api       ExternalAPI
	wg        sync.WaitGroup
	stopCh    chan struct{}
}

func NewAnalyticsCollector(numWorks int, interval time.Duration, sizeBatch int, api ExternalAPI) *AnalyticsCollector {
	return &AnalyticsCollector{
		numWorks:  numWorks,
		interval:  interval,
		sizeBatch: sizeBatch,
		eventCh:   make(chan Event, sizeBatch),
		jobCh:     make(chan []Event, 10),
		stopCh:    make(chan struct{}),
		api:       api,
	}
}

func (c *AnalyticsCollector) Track(event Event) error {
	select {
	case <-c.stopCh:
		return nil
	case c.eventCh <- event:
	default:
		return errors.New("Service cannot handle events so backprassure occurs")
	}
	return nil
}

func (c *AnalyticsCollector) Start() {
	for range c.numWorks {
		c.wg.Add(1)
		go func() {
			defer c.wg.Done()
			for job := range c.jobCh {
				err := c.api.Send(job)
				if err != nil {
					fmt.Println("need to implement retry")
				}
			}
		}()
	}
	batch := []Event{}

	flush := func() {
		if len(batch) > 0 {
			c.jobCh <- batch
		}
		batch = []Event{}
	}

	go func() {
		for {
			select {
			case event, ok := <-c.eventCh:
				if !ok {
					return
				}
				batch = append(batch, event)
				if len(batch) == c.sizeBatch {
					flush()
				}
			case <-c.stopCh:
				flush()
				close(c.jobCh)
				return
			}
		}
	}()

	timer := time.NewTicker(c.interval)
	defer timer.Stop()

	for {
		select {
		case <-c.stopCh:
			return
		case <-timer.C:
			flush()
		}
	}
}

func (c *AnalyticsCollector) Stop() error {
	close(c.stopCh)

	c.wg.Wait()
	return nil
}

func main() {
}
