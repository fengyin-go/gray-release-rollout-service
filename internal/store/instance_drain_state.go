package store

import "errors"

type InstanceDrainStream struct{}

func NewInstanceDrainStream() *InstanceDrainStream { return &InstanceDrainStream{} }

func (s *InstanceDrainStream) Start(items []string, failAt int) (<-chan string, <-chan error) {
	out := make(chan string)
	errs := make(chan error, 1)
	go func() {
		// 无论正常结束还是中途失败，都必须关闭 out，否则消费者在
		// range 循环里拿不到通道关闭信号会永久阻塞——这正是“拿到一条
		// 成功结果后永久等待”的根因。错误分支同样需要走到这里。
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
