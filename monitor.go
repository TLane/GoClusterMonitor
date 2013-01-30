package main

// termbox-go documentation http://godoc.org/github.com/nsf/termbox-go
// Concurrency in go http://talks.golang.org/2012/concurrency.slide#1
import (
        "github.com/nsf/termbox-go"
        "time"
        //"io/ioutil"
        "container/list"
       )
       
func processInput(key termbox.Key, win int) {
  clear()
  displayCurrentWindow(win)
  flush()
}

func displayCurrentWindow(win int) {
  if win == 0 {
    mainMenu()
  } else {
    mainMenu()
  }
}


func write(x int, y int, mess string) {
  for _, c := range mess {
		termbox.SetCell(x, y, c, termbox.ColorDefault, termbox.ColorDefault)
		x++
	}
}

func mainMenu() {
	box("|","-")
  write(1,1, "Cluster Monitor - Main Menu")
  
  displayPings()
  
}


func displayPings() {
  
}

func loadNodes() *list.List {
  nodes := list.New()
  nodes.PushFront("node1")
  return nodes
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

func box(v string, h string) {
  x,y := termbox.Size()
  for i :=0; i < x-1; i++ {
    write(i,0,h)
    write(i,y-1,h)
  }
  for i :=0; i < y; i++ {
    write(0,i,v)
    write(x-1,i,v)
  }
  write(x-1,y-1,"·")
  write(0,0,"·")
  write(0,y-1,"·")
  write(x-1,0,"·")
}
        
func main() {
  err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
  c_window := 0;
  displayCurrentWindow(c_window);
	box("|","-")
	flush()
  loop:
	  for {
		  switch event := termbox.PollEvent(); event.Type {
		  case termbox.EventKey:
		    if event.Key == termbox.KeyCtrlQ {
		      clear()
		      centerMessage("Exiting...")
		      flush()
		      time.Sleep(1* time.Second)
				  break loop
			  } else if event.Key == termbox.KeyCtrlC {
			    clear()
			    flush()
			    break loop
			  } else {
			    processInput(event.Key, c_window)
			  }
		  case termbox.EventResize:
			  termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			  displayCurrentWindow(c_window)
		  case termbox.EventError:
			  panic(event.Err)
		}
  }
}
