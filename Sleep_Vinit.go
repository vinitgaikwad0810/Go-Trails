package main

import (
	"time"
	)


func Sleep_Vinit(seconds_to_sleep int) {
  <-time.After(time.Second * time.Duration(seconds_to_sleep))
}

/*

func main() {
//var number int64

fmt.Println("Before putting the golang program to sleep  .....",time.Now().Second())

Sleep_Vinit(5)

fmt.Println("After putting the golang program to sleep  .....:",time.Now().Second())

}
*/	

