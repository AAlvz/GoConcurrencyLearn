package main

import(
	"fmt"
	"time"
	"math/rand"
)

func main(){
     rand.Seed(time.Now().UnixNano())
     saluda(1)
     saludaOtra(2)
}

func saluda(id int){
     time.Sleep(rand.Intn(1000))
     fmt.Println("Hola, ", id)
}

func saludaOtra(id int){
     time.Sleep(rand.Intn(1000))
     fmt.Println("Que ondas ", id)
}