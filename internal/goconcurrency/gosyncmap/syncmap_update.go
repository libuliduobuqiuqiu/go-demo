package goconcurrency

import (
	"fmt"
	"sync"
)

var allCount sync.Map

type StoreCount struct {
	AddCount    int
	DeleteCount int
	UpdateCount int
}

func FirstStoreCount() {
	if v, ok := allCount.LoadOrStore("test", &StoreCount{}); ok {
		if tmp, ok := v.(*StoreCount); ok {
			fmt.Println("First StoreCount")
			fmt.Println(tmp.AddCount)
			tmp.AddCount += 1
			fmt.Println(tmp.DeleteCount)
			tmp.DeleteCount += 1
			fmt.Println(tmp.UpdateCount)
			tmp.UpdateCount += 1
		}
	}

}

func SecondStoreCount() {
	if v, ok := allCount.LoadOrStore("test", &StoreCount{}); ok {
		if tmp, ok := v.(*StoreCount); ok {
			fmt.Println("Second StoreCount")
			fmt.Println(tmp.AddCount)
			tmp.AddCount += 1
			fmt.Println(tmp.DeleteCount)
			tmp.DeleteCount += 1
			fmt.Println(tmp.UpdateCount)
			tmp.UpdateCount += 1
		}
	}
}

func ConcurrentSyncMap() {
	wg := &sync.WaitGroup{}
	allCount.Store("test", &StoreCount{})
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		FirstStoreCount()
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		SecondStoreCount()
	}(wg)

	wg.Wait()
}
