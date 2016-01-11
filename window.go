package main

import (
	"github.com/mattn/go-gtk/gtk"
)

func main(){
	gtk.Init(nil)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("Hello world!")
	window.SetIconName("Hello go!")
	window.Connect("destroy", gtk.MainQuit)
	window.SetSizeRequest(800,800)
	window.ShowAll()
	gtk.Main()
}
