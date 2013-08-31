package main
import (
  "fmt"
  "math/rand"
  "time"
  "strconv"
  "os"
)

/*
 to do in parallel, can't just pass pointers to slices, 
 need to pass separate lists to "go quicksort" calls,
 and have a master thread (main()) that reads from the channel
 and 'merges' them
*/

func quicksort(a []int) {
  if len(a) <= 1 {
    return
  }
  new_pivot_index := partition(a)
  quicksort(a[:new_pivot_index+1])
  quicksort(a[new_pivot_index+1:])
}

func partition(a []int) (new_pivot int) {
  left, right := 0, len(a) -1
  pivot_i := left + (right-left)/2
  p_val := a[pivot_i]
  swap_i := left
  a[pivot_i], a[right] = a[right], a[pivot_i]
  for i := left; i < right; i++ {
    if (a[i] < p_val) {
      a[swap_i], a[i] = a[i], a[swap_i]
      swap_i++
    }
  }
  a[swap_i], a[right] = a[right], a[swap_i]
  return swap_i
}

func main() {
  rand.Seed(time.Now().Unix());

  size, error := strconv.Atoi(os.Args[1])
  if error == nil {
    a := rand.Perm(size)

    start := time.Now()
    quicksort(a)
    elapsed_ns := time.Since(start)
    fmt.Println("sorted: ", a)

    fmt.Printf("took %f seconds \n", float64(elapsed_ns)/1000000000.0 )
  } else {
    fmt.Println("Error!: %s", error)
  }
}