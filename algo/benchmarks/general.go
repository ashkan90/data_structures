package benchmarks

import (
	"data_structures/algo"
	"data_structures/misc"
	"testing"
)

var On string

func BenchmarkBubbleSortMIXED(b *testing.B) {
	for i := 0; i < b.N; i++ {
		On = "MIXED"
		a := misc.GenerateSlice(misc.MIXED, 1000)
		algo.BubbleSort(a)
	}
}

func BenchmarkBubbleSortBACKSORTED(b *testing.B) {
	for i := 0; i < b.N; i++ {
		On = "BACK_SORTED"
		a := misc.GenerateSlice(misc.BACK_SORTED, 1000)
		algo.BubbleSort(a)
	}
}

func BenchmarkBubbleSortPRESORTED(b *testing.B) {
	for i := 0; i < b.N; i++ {
		On = "PRE_SORTED"
		a := misc.GenerateSlice(misc.PRE_SORTED, 1000)
		algo.BubbleSort(a)
	}
}

func BenchmarkQuickSortMIXED(b *testing.B) {
	for i := 0; i < b.N; i++ {
		On = "MIXED"
		sl := misc.GenerateSlice(misc.MIXED, 1000)
		algo.Qsort(sl)
	}
}

func BenchmarkQuickSortBACKSORTED(b *testing.B) {
	for i := 0; i < b.N; i++ {
		On = "BACK_SORTED"
		sl := misc.GenerateSlice(misc.BACK_SORTED, 1000)
		algo.Qsort(sl)
	}
}

func BenchmarkQuickSortPRESORTED(b *testing.B) {
	for i := 0; i < b.N; i++ {
		On = "PRE_SORTED"
		sl := misc.GenerateSlice(misc.PRE_SORTED, 1000)
		algo.Qsort(sl)
	}
}

func BenchmarkMergeSortMIXED(b *testing.B) {
	for i := 0; i < b.N; i++ {
		On = "MIXED"
		sl := misc.GenerateSlice(misc.MIXED, 1000)
		algo.MSort(sl)
	}
}

func BenchmarkMergeSortBACKSORTED(b *testing.B) {
	for i := 0; i < b.N; i++ {
		On = "BACK_SORTED"
		sl := misc.GenerateSlice(misc.BACK_SORTED, 1000)
		algo.MSort(sl)
	}
}

func BenchmarkMergeSortPRESORTED(b *testing.B) {
	for i := 0; i < b.N; i++ {
		On = "PRE_SORTED"
		sl := misc.GenerateSlice(misc.PRE_SORTED, 1000)
		algo.MSort(sl)
	}
}

func BenchmarkSelectionSortMIXED(b *testing.B) {
	for i := 0; i < b.N; i++ {
		On = "MIXED"
		sl := misc.GenerateSlice(misc.MIXED, 1000)
		algo.SelectionSort(sl)
	}
}

func BenchmarkSelectionSortBACKSORTED(b *testing.B) {
	for i := 0; i < b.N; i++ {
		On = "BACK_SORTED"
		sl := misc.GenerateSlice(misc.BACK_SORTED, 1000)
		algo.SelectionSort(sl)
	}
}

func BenchmarkSelectionSortPRESORTED(b *testing.B) {
	for i := 0; i < b.N; i++ {
		On = "PRE_SORTED"
		sl := misc.GenerateSlice(misc.PRE_SORTED, 1000)
		algo.SelectionSort(sl)
	}
}
