package main

import "fmt"
//note that the code source is :https://blog.csdn.net/caileigood/article/details/84766188
//please don't use the code for purposes other than self-educating

func testGoroutine(ch chan string, param string){	//means that the channel pass strings
	fmt.Println("testing")
	ch <- param 		//passing the parameter into channel

}


func main() {
  ch := make(chan string, 3)		//making a string that can hold 3 elements at once
  go testGoroutine(ch, "1")
  go testGoroutine(ch, "2")
  go testGoroutine(ch, "3")			//for some reasons, channel will be printing from the back? read channel more
  for i :=0; i< cap(ch); i++{		//will be deadlocked if not filling out all the blank spaces	
	fmt.Println(<-ch)				//print out info from channel
  }
}