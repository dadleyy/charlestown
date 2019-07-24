package engine

import "fmt"
import "time"
import "unicode/utf8"
import "github.com/gdamore/tcell"

type engine struct {
	*engineLogger
	screen tcell.Screen
	config Configuration
}

func (instance *engine) pollKeys(quit chan<- struct{}) {
	instance.Printf("polling screen keyboard")

	for {
		event := instance.screen.PollEvent()

		switch kind := event.(type) {

		case *tcell.EventKey:
			instance.Printf("received keyboard event")

			switch kind.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				instance.Printf("escape key pressed")
				quit <- struct{}{}
				return
			}
		}
	}
}

func (instance *engine) run() error {
	instance.Printf("starting screen")
	quit := make(chan struct{})

	if e := instance.screen.Init(); e != nil {
		return e
	}

	defer instance.screen.Fini()

	go instance.pollKeys(quit)

	messages := make(chan string, 6)

	go func() {
		<-time.After(time.Second)
		instance.Printf("sending message")
		messages <- "this is a test message"
		<-time.After(time.Second)
		close(messages)
	}()

	for {
		select {
		case message, ok := <-messages:
			instance.Printf("recieved message '%s'", message)

			if !ok || len(message) == 0 {
				instance.Printf("messages is now closed, exiting")
				return nil
			}

			instance.screen.Clear()

			character, size := utf8.DecodeRuneInString(message)
			instance.Printf("filling screen with '%v' (size %d)", character, size)
			instance.screen.Fill(character, tcell.StyleDefault)
			instance.screen.Show()
		case <-quit:
			instance.Printf("quit signal recieved")
			return fmt.Errorf("exit")
		}
	}

	return nil
}
