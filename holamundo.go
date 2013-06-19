package main

import(
	"fmt"
	"time"
	"math/rand"
)

func main(){
     
     rand.Seed(time.Now().UnixNano())
     for i := 1; i < 10; i++ {
     	 go saluda(i)
     	 go saludaOtra(i)
     }
     time.Sleep(200000)
}

func saluda(id int){
     time.Sleep(time.Duration(rand.Intn(1000)))
     fmt.Println("Hola, ", id)
}

func saludaOtra(id int){
     time.Sleep(time.Duration(rand.Intn(1000)))
     fmt.Println("Que ondas ", id)
}