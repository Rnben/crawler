package scheduler

import (
	"crawler/engine"
)

type QueueSchduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueueSchduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueueSchduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueueSchduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueueSchduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activateRequest engine.Request
			var activateWorker chan engine.Request

			if len(requestQ) > 0 && len(workerQ) > 0 {
				activateRequest = requestQ[0]
				activateWorker = workerQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activateWorker <- activateRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}

	}()
}
