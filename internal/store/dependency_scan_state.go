package store

import "errors"

type DependencyScanStream struct{}

func NewDependencyScanStream() *DependencyScanStream { return &DependencyScanStream{} }

func (s *DependencyScanStream) Start(items []string, failAt int) (<-chan string, <-chan error) {
	out := make(chan string)
	errs := make(chan error, 1)
	go func() {
		// defer 按 LIFO 执行：先关 out（让 Collect 的 range 收尾），再关 errs。
		// 出错时直接 return，未发送的 item 不会在已关闭的 out 上发送。
		defer close(errs)
		defer close(out)
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
