# Notes


## Traversing

* usually uses recursion, but can use a stack in replacement of recursion if you want to.

### Depth First Search (DFS)

* Can use a stack to keep track of the nodes it's visited and go back up if needed (node doesn't have access to parent), useful for non-recursive implementation
  * though if done recursively, it naturally is using the call __stack__ ;)

* Goes down as much as possible before going to next sibling
  * in the name

1. If current node is empty, return
1. Check current node
1. Check Left node recursively
1. check right node recursively

#### Pre-Order, Post-Order, In-Order

* Pre-Order - used to create a copy of a tree, also can get the prefix expression of a tree
  * NLR - Current node, left, right
* Post-Order - used to delete the tree, can be used to get the postfix expression of a tree
  * LRN - Left, right, Current node
* In-Order - Can be used in balanced trees to get a non-decreasing order of nodes
  * LNR - Left, Current node, right

### Breadth First Search (BFS)

* Goes through the tree level by level
* uses a queue to keep track of the children that will be visited next, even though they are not connected
* loop through the children, put it's children in the queue, and go to the next child
* Space complexity is O(log(n))
  * worst case is there is only 1 level and all the elements are on that level


### References 

https://medium.com/basecs/breaking-down-breadth-first-search-cebe696709d9

