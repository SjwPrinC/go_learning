package main

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {

	ErrGroup()
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	// wg := &sync.WaitGroup{}

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	Job1(ctx)
	// }()

	// wg.Wait()
}

func Job1(ctx context.Context) error {

	ctx1, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	select {
	case <-ctx1.Done():
		fmt.Print("cancelled")
	}

	return nil
}

func ChangeSliceValue(s []int) {
	for i := 1; i < 5; i++ {
		s = append(s, i)
	}

	fmt.Println(s)
	fmt.Printf("inner s: %p\n", s)
}

func GoRoutineTest() {
	// g := &errgroup.Group{}

	var i int32

	result := runtime.GOMAXPROCS(1)

	fmt.Println(result)

	result1 := runtime.GOMAXPROCS(3)
	fmt.Println(result1)

	fmt.Println(runtime.GOROOT())
	go func() {
		time.Sleep(time.Second)
		for {
			fmt.Println("i is: ", i)
			time.Sleep(time.Second)
		}
	}()

	var ii int32 = 0
	go func() {
		for {
			if i > 10 {
				break
			}
			i += 1
			ii = i
		}
	}()

	fmt.Println("last ii is: ", ii)
	fmt.Println("last i is: ", i)

	// runtime.Caller(1)

	fmt.Println(runtime.NumGoroutine())

	// m:=make(map[int]string)

	select {}
}

func ErrGroup() {
	g, _ := errgroup.WithContext(context.Background())

	sayHello := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}

	sayGoodBye := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("good bye"))
	}
	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/bye", sayGoodBye)

	g.Go(func() error {
		err := http.ListenAndServe("localhost:8080", nil)

		if err != nil {
			return err
		}
		fmt.Println("wef")
		return nil
	})

	g.Go(func() error {
		err := http.ListenAndServe("localhost:8081", nil)

		if err != nil {
			return err
		}
		fmt.Println("wef11")
		return nil
	})

	// http.ListenAndServe("127.0.0.1:8081", nil)
	g.Wait()
	fmt.Println("wef")
}
