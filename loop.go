package main

import "time"
import "fmt"

func main(){
	const MS_PER_UPDATE = 1000 / 60
	var prevTime = time.Now()
	var lag float64 = 0.0
	fmt.Println("Begin loop...")
	for{
		var current = time.Now()
		var elapsed = time.Since(prevTime)
		lag += elapsed.Seconds() * 1000
		prevTime = current
		input()
		for lag >= MS_PER_UPDATE{
			update()
			lag -= MS_PER_UPDATE
		}
		draw()
	}
}

func input(){
}

func update(){
}

func draw(){

}
