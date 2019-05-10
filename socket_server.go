package main
//note that the code source is from ://note that the code source is from :https://blog.csdn.net/caileigood/article/details/84661746
//the code is only used for learning pursose, not for bussinesses

import (
	"fmt"
	"log"
	"net"
	"time"
)

func startServer(){
	listener, err := net.Listen("tcp", "127.0.0.1:12031")	//use the same ip address to do the handshake
	defer listener.Close()			//defer must be used for a following function call, enable to auto close the file after reading and working on the files, can also be used in mutex lock
	if err != nil{			//print error if error occurs
		fmt.Println(err)
	}
	for {								  //blocks until connected with client 	
		conn, err := listener.Accept()   //accept the connection, listening  to the client request
		if err != nil{
			fmt.Println(err)
			continue
		}
		go receiveAndSendMsg(conn)		//working properly so that can send messages
		//when one goroutine ends, it will not affect other goroutine's functionality
		//time.Sleep(1*time.Second)	//to sleep in order to separte user connecting
	}
}

func receiveAndSendMsg(conn net.Conn){
	defer conn.Close()
	log.Println(conn.RemoteAddr().String())
	buffer := make([]byte, 50)	//can connect with maximum of 50 users
	
	for{
		conn.Write([]byte("hi, client, this is server:"+ time.Now().Format("10:04:05")))
		_, err := conn.Read(buffer)
		if err!= nil{		//have error, have to break the connection
			return
		}
		log.Printf("msg is : %v\n", string(buffer))			//read stringfied byte message
		//time.Sleep(1 * time.Second)
	
	}
	
}

func main() {
   startServer()
}