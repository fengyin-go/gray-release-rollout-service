package store

import "errors"

type InstanceDrainStream struct{}

func NewInstanceDrainStream() *InstanceDrainStream { return &InstanceDrainStream{} }

func (s *InstanceDrainStream) Start(items []string, failAt int) (<-chan string, <-chan error) {
	out := make(chan string)
	errs := make(chan error, 1)
	go func() {

		defer close(errs)
		for index, item := range items {
			if index == failAt {
				errs <- errors.New("stream interrupted")
				return
			}
			out <- item
		}
	}()
	return out, errs
}
