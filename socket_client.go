package main
//note that the code source is from ://note that the code source is from :https://blog.csdn.net/caileigood/article/details/84661746
//the code is only used for learning pursose, not for bussinesses
import (
	"fmt"
	"net"
	"os"
	"time"

)

func startClient(){
	conn,err := net.Dial("tcp", "127.0.0.1:12031")			//the conn here is the connection, tcp three way handshake
	defer conn.Close()
	if err != nil{
		fmt.Println(err)
	}
	receiveAndSendMsg(conn)
}

func receiveAndSendMsg(conn net.Conn){
	for{
		buffer := make([]byte, 50)
		_, err := conn.Read(buffer)			//because the server is writing first, client side read first
		if err != nil{		//stop the program
			os.Exit(-1)
		}
		fmt.Println(string(buffer))
		
		str := "hello, i am client:"+time.Now().Format("10:04:05")
		conn.Write([]byte(str))
		time.Sleep(1 * time.Second)
	}
}


func main() {
	startClient()
}