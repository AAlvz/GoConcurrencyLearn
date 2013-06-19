package main

import (
       "fmt"
       "net/http"
       "time"
)

var urls = []string {
    //"http://rubyconf.com",
    "http://golang.org", 
    "http://matt.aimonetti.net/",
    "http://101project.co",
}

type HttpResponse struct {
     url string
     response *http.Response
     err error
}

func Fetcher(url string, ch chan *HttpResponse){
    fmt.Printf("Fetching %s \n", url)
    resp, err := http.Get(url)
    ch <- &HttpResponse{url, resp, err}
}

func asyncHttpGets(urls []string) (responses []*HttpResponse) {
     ch := make(chan *HttpResponse)
     //response := []*HttpResponse //Sustituida por el parametro
     for _, uri := range urls {  //el _. descarta el primer valor que regresa ya que no interesa el index para este caso
     	 //creando la gorutine
	 go Fetcher(uri, ch)
     }
     
     for {
     	 select{
	    case r := <-ch:
	       fmt.Printf("%s was fetched \n", r.url)
	       responses = append(responses, r)
	       if len(responses) == len(urls){
	          return responses
	       }
	       
	    case <- time.After(50* time.Millisecond):
	       fmt.Print(".")
	 }
     }
     //es para que pueda compilar y no haya problemas con un metodo sin return . en realidad nunca se ejecuta.
     return responses
}

func main(){
     results := asyncHttpGets(urls)
     for _, result := range results {
        fmt.Printf("%s status: %s \n", result.url, result.response.Status)
     }
}