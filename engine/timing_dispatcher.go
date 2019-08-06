package engine

import "log"
import "sync"
import "time"
import "github.com/dadleyy/charlestown/engine/mutations"

type timingDispatcher struct {
	updates  chan<- mutations.Mutation
	mutators []mutator
}

func (dispatch *timingDispatcher) start(kill <-chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	timer := time.Tick(time.Millisecond * 100)

	log.Printf("entering loop")

	for {
		last := time.Now()

		select {
		case <-kill:
			log.Printf("kill signal received, exiting")
			return
		case t := <-timer:
			dt := t.Sub(last)

			for _, m := range dispatch.mutators {
				updates := m.tick(dt)

				for _, u := range updates {
					dispatch.updates <- u
				}
			}
		}
	}

	log.Printf("dispatcher shuttiung down")
}
