package scheduler

import "crawler/distributedConcurrentQueueCrawler/engine"

type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueueScheduler) WorkChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueueScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueueScheduler) WorkReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueueScheduler) ConfigMasterWorkerChan(r chan engine.Request) {
}

func (s *QueueScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorkerRequest chan engine.Request
			if len(requestQ) > 0 &&
				len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorkerRequest = workerQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorkerRequest <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
