package main

import (
"fmt"
"container/list"
)

func main() {

    l := list.New()
    e4 := l.PushBack(4)
    e1 := l.PushFront(1)
    l.InsertBefore(3, e4)
    l.InsertAfter(2, e1)


    // should be 4.
    fmt.Println(l.Len())
    fmt.Println("-------\n")

    // Iterate through list and and print its contents.
    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
    }
    fmt.Println("-------\n")

    // the head of the list,in this case it should be 1
    fmt.Println("the head of list is " , l.Front().Value)
    fmt.Println("-------\n")

    // the tail of the list,in this case it should be 4
    // why the method name is Back but not Tail?
    fmt.Println("the tail of list is ", l.Back().Value)
    fmt.Println("-------\n")

    // remove the last element
    var b = l.Back()
    l.Remove(b)
    // the last element should now be 3
    fmt.Println("the tail of list now is ", l.Back().Value)

    // form this source code ,we can see that a list is a bi-direction linked table.
/**

    15  type Element struct {
    16      // Next and previous pointers in the doubly-linked list of elements.
    17      // To simplify the implementation, internally a list l is implemented
    18      // as a ring, such that &l.root is both the next element of the last
    19      // list element (l.Back()) and the previous element of the first list
    20      // element (l.Front()).
    21      next, prev *Element
    22  
    23      // The list to which this element belongs.
    24      list *List
    25  
    26      // The value stored with this element.
    27      Value interface{}
    28  }

    48   type List struct {
    49      root Element // sentinel list element, only &root, root.prev, and root.next are used
    50      len  int     // current list length excluding (this) sentinel element
    51  }
*/
}
