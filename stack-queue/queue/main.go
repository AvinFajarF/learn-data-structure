package main

import "fmt"

// Queue
type Queue struct {
  items []int
}


// Enqueue
func (q *Queue) Enqueue(i int){
  q.items = append(q.items, i)
}

// Deqeue
func (q *Queue) Deqeue() int{
  toRemove := q.items[0]
  q.items = q.items[1:]
  return toRemove
}

func main(){
 myQueue := Queue{}
 myQueue.Enqueue(30)
 myQueue.Enqueue(31)
 myQueue.Enqueue(39)
 fmt.Println(myQueue)

  data := myQueue.Deqeue()
  fmt.Println(data)
  fmt.Println(myQueue)
}


