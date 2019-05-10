//simplepipeline

package simple

import "fmt"
import "sync"

func Test() {
        fannnedinchan := mergefanIn(sq(gen(4, 8, 2)), sq(gen(1, 3, 5)))
        for v := range fannnedinchan {
                fmt.Println(v)
        }
}

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

func sq(in <-chan int) <-chan int {
        out := make(chan int)

        go func() {
                for n := range in {
                        out <- n * n
                }
                close(out)
        }()
        return out
}

func mergefanIn(ch ...<-chan int) <-chan int {
        out := make(chan int)
        wg := &sync.WaitGroup{}
        wg.Add(len(ch))

        output := func(ch <-chan int) {
                for n := range ch {
                        out <- n
                }
                wg.Done()
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
