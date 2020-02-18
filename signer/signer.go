package main

import (
	"sort"
	"strconv"
)

func SingleHash(in, out chan interface{}) {
	for {
		data, ok := <-in
		if !ok {
			break
		}
		out <- DataSignerCrc32(data.(string)) + "~" + DataSignerCrc32(DataSignerMd5(data.(string)))
	}
	close(out)
}

func MultiHash(in, out chan interface{}) {
	for {
		var result string
		data, ok := <-in

		if !ok {
			break
		}

		for th := 0; th < 6; th++ {
			result += DataSignerCrc32(strconv.Itoa(th) + data.(string))
		}

		out <- result
	}
	close(out)
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
	close(out)
}

func ExecutePipeline(jobs ...job) {
	in := make(chan interface{})
	out := make(chan interface{})
	for _, job := range jobs {
		go job(in, out)
	}
}
