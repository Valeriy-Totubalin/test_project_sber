package rate_limiter

import "runtime"

type Worker struct {
	limiter           RateLimiterInterface
	queue             []func()
	simultaneousTasks int
}

func NewWorker(
	limiter RateLimiterInterface,
	simultaneousTasks int,
) *Worker {
	if simultaneousTasks < 1 {
		simultaneousTasks = 1
	}
	return &Worker{
		limiter:           limiter,
		simultaneousTasks: simultaneousTasks,
	}
}

func (w *Worker) DoWork(c <-chan []func()) {
	runtime.GOMAXPROCS(w.simultaneousTasks)
	go w.work(c)
}

func (w *Worker) work(c <-chan []func()) {
	for {
		for _, fun := range w.queue {
			if !w.limiter.CanDoWork() {
				continue
			}
			go fun()
			w.queue = w.queue[1:len(w.queue)]
		}

		w.queue = append(w.queue, <-c...)
	}
}
