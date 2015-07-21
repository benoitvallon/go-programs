package main

import (
  "fmt"
)

func main() {
  array := []int{9, 6, 7, 3, 2, 4, 5, 1}

  fmt.Println("array before")
  fmt.Println(array)

  bubble(array)

  array1 := []int{9, 6, 7, 3, 2, 4, 5, 1}
  bubble_opti1(array1)

  array2 := []int{9, 6, 7, 3, 2, 4, 5, 1}
  bubble_opti2(array2)
}

func bubble(array []int) {
  fmt.Println("\nBubble basic")
  lenght := len(array)
  iterations := 0

  for j := 0; j < lenght; j++ {
    swapped := false

    // swap values between previous and current if needed
    for i := 1; i < lenght; i++ {
      iterations++
      if array[i - 1] > array[i] {
        array[i - 1], array[i] = array[i], array[i - 1]
        swapped = true
      }
    }

    if swapped == false {
      break
    }
  }

  fmt.Println(array)
  fmt.Printf("iterations: %d\n", iterations)
}

func bubble_opti1(array []int) {
  fmt.Println("\nBubble opti1")
  lenght := len(array)
  iterations := 0

  swapped := true

  for swapped == true {
    swapped = false
    for i := 1; i < lenght; i++ {
      iterations++
      if array[i - 1] > array[i] {
        array[i - 1], array[i] = array[i], array[i - 1]
        swapped = true
      }
    }
    lenght = lenght - 1
  }

  fmt.Println(array)
  fmt.Printf("iterations: %d\n", iterations)
}

func bubble_opti2(array []int) {
  fmt.Println("\nBubble opti2")
  lenght := len(array)
  iterations := 0

  for lenght > 0 {
    newLenght := 0
    for i := 1; i < lenght; i++ {
      iterations++
      if array[i - 1] > array[i] {
        array[i - 1], array[i] = array[i], array[i - 1]
        newLenght = i
      }
    }
    lenght = newLenght
  }

  fmt.Println(array)
  fmt.Printf("iterations: %d\n", iterations)
}
