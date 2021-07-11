package main

import (
	"./InsSort"
	"testing"
	"reflect"
	"github.com/stretchr/testify/assert"
)

func TestInsSort(t *testing.T) {
	arrToSort := []float32{36,24,10,6,12}
	arrExpected := []float32{36,24,10,6,12}
	got := InsSort.InsertionSort(&arrToSort)
	want := &arrExpected

	if !reflect.DeepEqual(got,want){
		t.Errorf("got %v want %v",arrToSort,arrExpected)
	}
	//Table method
	/*table := [] struct {
		arg []float32
		want []float32
	}{
		{[]float32{36,24,10,6,12},[]float32{36,24,10,6,12}},
		{[]float32{32,76,21,45,12},[]float32{76,45,32,21,12}},
	}
	for _, entry := range table {
		got := InsSort.InsertionSort(entry.arg)
		if !reflect.DeepEqual(got,want) {
			t.Errorf("got %v want %v",entry.arg,entry.want)
		}
	}*/
	//Testify
	/*table := [] struct {
		arg []float32
		want []float32
	}{
		{[]float32{36,24,10,6,12},[]float32{36,24,10,6,12}},
		{[]float32{32,76,21,45,12},[]float32{76,45,32,21,12}},
	}
	for _, entry := range table {
		assert.Equal(t,InsertionSort(entry.arg),entry.want,"SortFailed" )
		}
	}*/
}
