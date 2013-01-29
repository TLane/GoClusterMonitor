package main

import (
        "github.com/nsf/termbox-go"
        "time"
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
func flush() {
  termbox.Flush()
}
func centerMessage(mess string) {
  x,y := termbox.Size()
  write( ((x/2)-(len(mess)/2)) , ((y-1)/2), mess)
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
		      clear()
		      centerMessage("Exiting...")
		      flush()
		      time.Sleep(2* time.Second)
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
