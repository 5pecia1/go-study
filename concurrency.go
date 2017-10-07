package go_study

import (
	"fmt"
	"time"
	"math/rand"
)

var (
	Web = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Result string

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func Google(query string) (results []Result)  {
	c := make(chan Result)
	go func() {c <- Web(query)}()
	go func() {c <- Image(query)}()
	go func() {c <- Video(query)}()
	//results = append(results, Web(query))
	//results = append(results, Image(query))
	//results = append(results, Video(query))
	//return results

	timeout := time.After(80 * time.Millisecond)

	for i:=0; i<3; i++ {
		select {
			case result := <-c:
				results = append(results, result)
			case <- timeout:
				fmt.Println("time out")
				return
		}
		//result := <-c
		//results = append(results, result)
	}

	return
}

func main() {
	result := Google("동시성")
	fmt.Println(result)
}
