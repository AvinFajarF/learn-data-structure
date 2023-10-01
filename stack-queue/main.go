package main

import "fmt"

//stack represents stack that hold a slice
type Stack struct {
  items []int
}


// push
func (s *Stack) Push(i int){
  s.items = append(s.items, i)    
}


// pop
func (s *Stack) Pop() int {
  l := len(s.items) -1
  toRemove := s.items[l]
  s.items = s.items[:l]
  return toRemove
}


func main(){
  myStack := Stack{}
  fmt.Println(myStack)  
  myStack.Push(200)
  myStack.Push(50000000)
  myStack.Push(12022311230123120)
  fmt.Println(myStack)
  data := myStack.Pop()
  fmt.Println(data)
  fmt.Println(myStack)
}
