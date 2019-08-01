package engine

import "log"
import "time"
import "sync"
import "github.com/dadleyy/charlestown/engine/mutations"
import "github.com/dadleyy/charlestown/engine/constants"

type economyManager struct {
	*log.Logger
	updates chan<- mutations.Mutation
}

func (manager *economyManager) tick(stop <-chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	timer := time.Tick(constants.IncomeDelay)

	for {
		next := mutations.Income(time.Now())

		select {
		case <-stop:
			manager.Printf("timer received stop signal, breaking loop")
			return
		case <-timer:
			manager.Printf("[econ] pushing income")
			manager.updates <- next
		}
	}

}
