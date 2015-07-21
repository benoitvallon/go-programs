package main

import (
  "fmt"
)

func main() {
  array := []int{9, 6, 7, 3, 2, 4, 5, 1}

  fmt.Println("array before")
  fmt.Println(array)

  fmt.Println("\nQuick basic")
  quick(array, 0, len(array))

  fmt.Println(array)
}

func quick(array []int, left int, right int) {
  if left < right {
    pivot := (right + left) / 2
    pivot = partition(array, left, right, pivot)
    quick(array, left, pivot)
    quick(array, pivot + 1, right)
  }
}

// left is the index of the leftmost element of the subarray
// right is the index of the rightmost element of the subarray (inclusive)
func partition(array []int, left int, right int, pivot int) (int) {
  pivotValue := array[pivot]
  right--

  // put the chosen pivot at array[right]
  array[pivot], array[right] = array[right], array[pivot]

  // Compare remaining array elements against pivotValue = array[right]
  for i := left; i <= right - 1; i++  {
    if array[i] < pivotValue {
      array[left], array[i] = array[i], array[left]
      left = left + 1
    }
  }

  // Move pivot to its final place
  array[left], array[right] = array[right], array[left]
  return left
}
