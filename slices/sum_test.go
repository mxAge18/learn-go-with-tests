package slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, want, got int, nums []int) {
		t.Helper()
		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, nums)
		}
	}
	t.Run("collection of non-limited element of slices", func(t *testing.T) {
		nums := []int{6,8,10}
		got := Sum(nums)
		want := 24
		assertCorrectMessage(t, want, got, nums)
	})
}


func TestSumAll(t *testing.T) {
	got := SumAll([]int{1,2}, []int{1,2,3,3})
	want := []int{3, 9}
	// want := "invokerx" // 编译仍通过哦
	
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v want %v ", got, want)
	}
}


// 把每个切片的尾部元素相加（尾部的意思就是除去第一个元素以外的其他元素）。
func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    }
	t.Run("make the sums of some slices", func(t *testing.T) {
        got := SumAllTails([]int{2,3}, []int{1, 3, 4, 5})
        want :=[]int{3, 12}

        checkSums(t, got, want)
    })
	t.Run("safely sum empty slices", func(t *testing.T) {
        got := SumAllTails([]int{}, []int{32, 4, 5})
        want :=[]int{0, 9}

        checkSums(t, got, want)
    })
}
