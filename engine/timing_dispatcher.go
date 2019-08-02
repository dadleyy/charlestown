package engine

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

	for {
		last := time.Now()

		select {
		case <-kill:
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
}
