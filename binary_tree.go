package main

import (
  "fmt"
  "math/rand"
  "os"
  "strconv"
)

type Tree struct {
  left *Tree
  value int
  right *Tree
}

// due to the sorted nature of binary trees,
// this depth-first walk will send the integer node values
// to the channel in least-to-greatest order
func DepthFirstWalk(t *Tree, ch chan int) {
  // base case
  if t == nil {
    return;
  }
   // walk down the left subtree
  DepthFirstWalk(t.left, ch)
  // send the value to the channel
  ch <- t.value
  // walk down right subtree
  DepthFirstWalk(t.right, ch)
}

// a function that walks down a tree and returns 
// the resulting channel, which may or may
// not be closed at the time of the return
func WalkedTreeChannel(tree *Tree) <- chan int {
  channel := make(chan int)

  // anonymous go routine to fire off a parallel
  // depth-first walk of the tree that fills the channel
  // and then closes it.
  go func() {
    DepthFirstWalk(tree, channel)
    close(channel)
  }()
  return channel
}


func EqualTreeChannels(tree_1 *Tree, tree_2 *Tree) bool {
  t1_channel := WalkedTreeChannel(tree_1)
  t2_channel := WalkedTreeChannel(tree_2)
  for {

    // read in the values from the tree channels
    t1_value, read_ok_t1 := <-t1_channel
    t2_value, read_ok_t2 := <-t2_channel

    // if we've read all the way to the end,
    // compare values. Final case of false/false means
    if !read_ok_t1 || !read_ok_t2 {
      return t1_value == t2_value
    }

    // if we come across something that differs, we're done
    // looking. and we can break out and return false
    if t1_value != t2_value {
      break
    }
  }
  return false
}


// New returns a new, random binary tree
// holding the values 1, 2, ..., n, inserted in random order
func NewTree(n int) *Tree {
  var tree *Tree
  for _, to_insert := range rand.Perm(n) {
    tree = Insert(tree, to_insert)
  }
  return tree
}

// inserts to_insert into the tree
func Insert(tree *Tree, to_insert int) *Tree {
  if tree == nil {
    return &Tree{nil, to_insert, nil}
  } else {
    if to_insert < tree.value {
      tree.left = Insert(tree.left, to_insert)
    } else {
      tree.right = Insert(tree.right, to_insert)
    }
  }
  return tree
}


func main() {
  size, error := strconv.Atoi(os.Args[1])
  if error == nil {
    fmt.Println("Equal?:", EqualTreeChannels(NewTree(size), NewTree(size)))
  } else {
    fmt.Println("Error!: %s", error)
  }
  
}