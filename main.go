package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	arrInts := StrToArr(string(content))
	avg := av(arrInts)
	varc := vars(arrInts, avg)

	fmt.Printf("Average: %d\n", int(math.Round(avg)))
	fmt.Printf("Median: %d\n", int(math.Round(med(arrInts))))
	fmt.Printf("Variance: %d\n", int(math.Round(math.Round(vars(arrInts, avg)))))
	fmt.Printf("Standard Deviation: %d\n", int(math.Round(std(varc))))
}

func StrToArr(s string) []int {
	arr := strings.Fields(s)
	var nums []int
	for i := range arr {
		num, err := strconv.Atoi(arr[i])
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	return nums
}

func av(nums []int) float64 {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	av := float64(sum) / float64(len(nums))
	return av
}

func med(nums []int) float64 {
	sort.Ints(nums)
	med := 0.0
	if len(nums)%2 != 0 {
		med = float64(nums[len(nums)/2])
	} else {
		med = av([]int{nums[len(nums)/2-1], nums[len(nums)/2]})
	}
	return med
}

func vars(nums []int, avg float64) float64 {
	sum := 0.0
	for _, v := range nums {
		sum += math.Pow((float64(v) - avg), 2.0)
	}
	return (sum / float64(len(nums)))
}

func std(nums float64) float64 {
	return math.Sqrt(nums)
}
