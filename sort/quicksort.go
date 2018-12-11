package main
import (
	"fmt"
	"math/rand"
  "time"
)

func main() {

    slice := generateSlice(20)
    fmt.Println("\n--- Unsorted --- \n\n", slice)
    qsort(slice)
    fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

    slice := make([]int, size, size)
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < size; i++ {
        slice[i] = rand.Intn(999) - rand.Intn(999)
    }
    return slice
}

func qsort(a []int) []int {
  if len(a) < 2 { return a }

  left, right := 0, len(a) - 1

  // Pick a pivot
  pivotIndex := rand.Int() % len(a)

  // Move the pivot to the right
  a[pivotIndex], a[right] = a[right], a[pivotIndex]

  // Pile elements smaller than the pivot on the left
  for i := range a {
    if a[i] < a[right] {
      a[i], a[left] = a[left], a[i]
      left++
    }
  }

  // Place the pivot after the last smaller element
  a[left], a[right] = a[right], a[left]

  // Go down the rabbit hole
  qsort(a[:left])
  qsort(a[left + 1:])


  return a
}
