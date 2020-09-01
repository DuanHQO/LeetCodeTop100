package kit

import (
	"math/rand"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	//将时间戳设置成种子数
	var arr []int
	rand.Seed(time.Now().UnixNano())
	//生成10个0-99之间的随机数
	for i := 0; i < 20; i++ {
		arr = append(arr, rand.Intn(100))
	}
	MergeSort(arr, 0, len(arr)-1)
}
