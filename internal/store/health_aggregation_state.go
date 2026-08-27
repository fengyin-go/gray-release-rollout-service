package store

import "errors"

type HealthAggregationStream struct{}

func NewHealthAggregationStream() *HealthAggregationStream { return &HealthAggregationStream{} }

func (s *HealthAggregationStream) Start(items []string, failAt int) (<-chan string, <-chan error) {
	out := make(chan string)
	errs := make(chan error, 1)
	go func() {
		// 关键：生产方退出时必须关闭 out，否则消费方 Collect 的 range out
		// 在收到第一份结果后会永远阻塞在等待下一个值上，既无法收尾，
		// 也读不到已经入队的 errs，已收集的数据随之丢失。
		// 出错分支同样依赖此 defer：发送错误后 return，由 defer 关闭 out，
		// 让本轮已发送的数据被消费完后结束 range，Collect 随即读取错误
		// 马上收尾并返回已保留的数据。
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
