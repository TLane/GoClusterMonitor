package main

import (
        "github.com/nsf/termbox-go"
       )
        
func main() {
  err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

  loop:
	  for {
		  switch event := termbox.PollEvent(); event.Type {
		  case termbox.EventKey:
		    if event.Key == termbox.KeyCtrlQ {
				  break loop
			  }
			  termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			  termbox.Flush()
		  case termbox.EventResize:
			  termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			  termbox.Flush()
		  case termbox.EventError:
			  panic(event.Err)
		}
  }
}
