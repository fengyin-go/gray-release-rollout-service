package store

import "errors"

type EventExportStream struct{}

func NewEventExportStream() *EventExportStream { return &EventExportStream{} }

func (s *EventExportStream) Start(items []string, failAt int) (<-chan string, <-chan error) {
	out := make(chan string)
	errs := make(chan error, 1)
	go func() {
		// 错误路径与正常结束路径都必须关闭结果流，
		// 否则消费端 range out 会永久阻塞，无法读取已交付的错误。
		defer close(out)
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
