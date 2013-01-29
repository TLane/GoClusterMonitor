package main

import (
        "github.com/nsf/termbox-go"
       )
       
       
func processInput(key termbox.Key) {

}

func write(x int, y int, mess string) {
  for _, c := range mess {
		termbox.SetCell(x, y, c, termbox.ColorDefault, termbox.ColorDefault)
		x++
	}
}
func clear() {
  termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}
        
func main() {
  err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	write(1,1,"Hello, world")
	termbox.Flush()
  loop:
	  for {
		  switch event := termbox.PollEvent(); event.Type {
		  case termbox.EventKey:
		    if event.Key == termbox.KeyCtrlQ {
				  break loop
			  } else {
			    processInput(event.Key)
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
