package main

import(
	"fmt"
	"time"
	"math/rand"
)

func main(){
     keepCalm := make(chan int)
     
     rand.Seed(time.Now().UnixNano())
     for i := 1; i < 10; i++ {
     	 go saluda(i, keepCalm)
     	 go saludaMundo(keepCalm)
     }
     time.Sleep(200000)
}

func saluda(id int, canal chan int){
     time.Sleep(time.Duration(rand.Intn(1000)))
     fmt.Println("Hola, ", id)
     canal <- id
}

func saludaMundo(canal chan int){
     id := <- canal
     fmt.Println("Que ondas ", id)
}