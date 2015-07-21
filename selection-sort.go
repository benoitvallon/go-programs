package main

import (
  "fmt"
)

func main() {
  array := []int{9, 6, 7, 3, 2, 4, 5, 1}

  fmt.Println("array before")
  fmt.Println(array)

  selection(array)
}

func selection(array []int) {
  fmt.Println("\nSelection basic")
  lenght := len(array)
  iterations := 0


  for i := 0; i < lenght - 1; i++ {
    // set the smallest as the first for now
    minimum := i

    // find the smallest amoung others array values
    for j := i + 1; j < lenght; j++ {
      iterations++
      if array[j] < array[minimum] {
        minimum = j;
      }
    }

    // smallest found so swap the values
    if minimum != i {
      array[minimum], array[i] = array[i], array[minimum]
    }
  }

  fmt.Println(array)
  fmt.Printf("iterations: %d\n", iterations)
}
