package main

import (
	"sort"
	"strconv"
	"sync"
)

func SingleHash(in, out chan interface{}) {
	wg := new(sync.WaitGroup)

	for {
		data, ok := <-in
		if !ok {
			break
		}

		val := strconv.Itoa(data.(int))
		local := make(chan interface{})

		wg.Add(1)
		go func(local chan interface{}, val string) {
			defer wg.Done()
			local <- DataSignerCrc32(val)
		}(local, val)

		md5 := DataSignerMd5(val)

		wg.Add(1)
		go func(local, out chan interface{}, val string) {
			defer wg.Done()
			val = DataSignerCrc32(val)
			out <- (<-local).(string) + "~" + val
		}(local, out, md5)
	}

	wg.Wait()
}

func MultiHash(in, out chan interface{}) {
	wg := new(sync.WaitGroup)

	for {
		data, ok := <-in
		if !ok {
			break
		}

		wg.Add(1)
		go func(data interface{}) {
			defer wg.Done()
			channels := []chan interface{}{
				make(chan interface{}),
				make(chan interface{}),
				make(chan interface{}),
				make(chan interface{}),
				make(chan interface{}),
				make(chan interface{}),
			}

			for th, ch := range channels {
				wg.Add(1)
				go func(th int, local chan interface{}) {
					defer wg.Done()
					res := DataSignerCrc32(strconv.Itoa(th) + data.(string))
					local <- res
				}(th, ch)
			}

			wg.Add(1)
			go func(channels []chan interface{}, out chan interface{}) {
				var result string
				defer wg.Done()

				for _, ch := range channels {
					result += (<-ch).(string)
				}

				out <- result
			}(channels, out)
		}(data)
	}

	wg.Wait()
}

func CombineResults(in, out chan interface{}) {
	var result string
	var storage []string

	for {
		data, ok := <-in
		if !ok {
			break
		}

		storage = append(storage, data.(string))
	}

	sort.Strings(storage)

	for _, str := range storage {
		if result != "" {
			result += "_"
		}
		result += str
	}

	out <- result
}

func ExecutePipeline(jobs ...job) {
	wg := new(sync.WaitGroup)
	in := make(chan interface{})

	for _, j := range jobs {
		out := make(chan interface{})

		wg.Add(1)

		go func(j job, in, out chan interface{}) {
			defer func() {
				close(out)
				wg.Done()
			}()
			j(in, out)
		}(j, in, out)

		in = out
	}

	wg.Wait()
}
