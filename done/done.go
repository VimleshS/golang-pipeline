package done

import (
        "fmt"
        "sync"
        "time"
)

func Test() {
        done := make(chan struct{})

        //fannnedinchan := mergefanIn(done, sq(done, gen(done, 4, 8, 7, 4, 6, 8, 8, 8, 3343, 52, 4, 457, 48, 49, 04, 14, 34, 45, 2)), sq(done, gen(done, 1, 3, 5)))
        fannnedinchan := mergefanIn(done, sq(done, gen(4, 8, 7, 4)), sq(done, gen(1, 3, 5)))
        /*
                for v := range fannnedinchan {
                        fmt.Println(v)
                }
        */

        fmt.Println(<-fannnedinchan)
        fmt.Println(<-fannnedinchan)
        close(done)

        <-time.After(3 * time.Second)
}

/*
func gen(done <-chan struct{}, nums ...int) <-chan int {
        out := make(chan int)

        go func() {
                defer close(out)
                for _, i := range nums {
                        select {
                        case <-done:
                                fmt.Println("aborting gen")
                                return
                        default:
                                out <- i
                        }
                }
        }()
        return out
}
*/
func gen(nums ...int) <-chan int {
        out := make(chan int)

        go func() {
                for _, i := range nums {
                        out <- i
                }
                close(out)
        }()
        return out
}
func sq(done <-chan struct{}, in <-chan int) <-chan int {
        out := make(chan int)
        go func() {
                defer close(out)
                defer fmt.Println("sq completes")
                for n := range in {
                        select {
                        case <-done:
                                fmt.Println("aborting sq")
                                return
                        default:
                                out <- n * n
                        }
                }

        }()
        return out
}

func mergefanIn(done <-chan struct{}, ch ...<-chan int) <-chan int {
        out := make(chan int)
        wg := &sync.WaitGroup{}
        wg.Add(len(ch))

        output := func(ch <-chan int) {
                defer wg.Done()
                for n := range ch {
                        select {
                        case <-done:
                                fmt.Println("aborting mergefanin")
                                return
                        default:
                                out <- n

                        }
                }
        }

        for _, c := range ch {
                go output(c)
        }

        go func() {
                wg.Wait()
                close(out)
        }()
        return out
}
