package main

import (
  "fmt"
)


func main() {
  array := []int{9, 6, 7, 3, 2, 4, 5, 1}

  fmt.Println("array before")
  fmt.Println(array)

  insertion(array)
}

func insertion(array []int) {
  fmt.Println("\nInsertion basic")
  lenght := len(array)
  iterations := 0

  for i := 1; i < lenght; i++ {
    j := i
    for j > 0 && array[j - 1] > array[j] {
      iterations++
      array[j - 1], array[j] = array[j], array[j - 1]
      j = j - 1
    }
  }

  fmt.Println(array)
  fmt.Printf("iterations: %d\n", iterations)
}
