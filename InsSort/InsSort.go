package InsSort

import "fmt"
//Задание 4
func InsSort()  {
	var arrToSort = make([]float32, 1)
	arrToSort = append(arrToSort,36,24,10,6,12)
	InsertionSort(&arrToSort)
}

func InsertionSort (numsToSort *[]float32)(*[]float32)  {
	var finished,moreToSearch bool
	var  current,startIndex,count int

	for  count = 0; count < len(*numsToSort); count++  {
		startIndex = 0
		finished = false
		current = count
		moreToSearch = current !=startIndex
		for {
			if !(moreToSearch && !finished) {
				break
			}
				if (*numsToSort)[current] < (*numsToSort)[current-1] {
					var tempVar float32
					tempVar = (*numsToSort)[current]
					(*numsToSort)[current] = (*numsToSort)[current-1]
					(*numsToSort)[current-1] = tempVar
					current--
					moreToSearch = current != startIndex
				}else {
					finished = true
				}
		}
	}
	fmt.Printf("%v", *numsToSort)
	return *numsToSort
}