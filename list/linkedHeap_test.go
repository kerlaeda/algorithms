package list

import (
	"testing"
	"reflect"
	"fmt"
	"algorithms/heap"
	"sort"
	"math/rand"
)

func Test_likedHeap(t *testing.T) {
	arrSize := rand.Intn(100) + 50
	arr := make([]int, arrSize, arrSize)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	sortedArr := make([]int, 0, 0)
	lh := new(LinkedHeap)
	lh.Init()
	h := heap.Heap{lh}
	for _, v := range arr {
		h.Append(v)
	}
	for h.Len() > 0 {
		sortedArr = append(sortedArr, h.Pop().(int))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	if !reflect.DeepEqual(sortedArr, arr) {
		t.Log(fmt.Sprintf("expect:%v", arr) + fmt.Sprintf("but get:%v", sortedArr))
		t.Fail()
	}
}

func Benchmark_likedHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrSize := 10000
		arr := make([]int, arrSize, arrSize)
		for i := range arr {
			arr[i] = rand.Intn(arrSize)
		}
		b.StartTimer()

		sortedArr := make([]int, 0, 0)
		lh := new(LinkedHeap)
		lh.Init()
		h := heap.Heap{lh}
		for _, v := range arr {
			h.Append(v)
		}
		for h.Len() > 0 {
			sortedArr = append(sortedArr, h.Pop().(int))
		}
	}

}
