package engine

import (
	"crawler/concurrentCrawler/utils"
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkReady(w chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkCount; i++ {
		createWorker(e.Scheduler.WorkChan(), out)
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
			log.Printf("获取的数据: %v\n", item)
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			//去除重复的请求
			if utils.IsDuplicate(request.Url) {
				continue
			}

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
