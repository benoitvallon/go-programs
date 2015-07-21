package main

import (
  "fmt"
)

func main() {
  array := []int{9, 6, 7, 3, 2, 4, 5, 1}

  fmt.Println("array before")
  fmt.Println(array)

  fmt.Println("\nMerge basic")
  array = merge(array)

  fmt.Println(array)
}

func merge(array []int) ([]int) {
  // Base case. A list of zero or one elements is sorted, by definition.
  if len(array) <= 1 {
    return array
  }

  // Recursive case. First, *divide* the list into equal-sized sublists.
  middle := len(array) / 2

  // Recursively sort both sublists
  left := merge(array[:middle])
  right := merge(array[middle:])

  // Then merge the now-sorted sublists.
  return merge_unit(left, right)
}

func merge_unit(left []int, right []int) ([]int) {
  var result []int

  for len(left) > 0 && len(right) > 0 {
    if left[0] <= right[0] {
      result = append(result, left[0])
      left = left[1:]
    } else {
      result = append(result, right[0])
      right = right[1:]
    }
  }

  // either left or right may have elements left
  for len(left) > 0 {
    result = append(result, left[0])
    left = left[1:]
  }
  for len(right) > 0 {
    result = append(result, right[0])
    right = right[1:]
  }

  return result
}
