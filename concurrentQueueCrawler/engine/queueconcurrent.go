package engine

import "crawler/concurrentQueueCrawler/utils"

type QueueConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
	ItemChan  chan Item
}

func (e *QueueConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkCount; i++ {
		createWorkerQueue(e.Scheduler.WorkChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if utils.IsDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
		}

		for _, request := range result.Requests {
			if utils.IsDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func createWorkerQueue(in chan Request,
	out chan ParseResult, s Scheduler) {
	go func() {
		for {
			s.WorkReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			go func() {
				out <- result
			}()
		}
	}()
}
