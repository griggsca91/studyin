# Merge Sort

## Concepts

* Divide and conquer
* Recursive
* Split the array into half until the array has 0-1 elements left
* Begin the Merge


## Pseudocode

* Declare left variable to 0 and right variable to n-1 
* Find mid by medium formula. mid = (left+right)/2
* Call merge sort on (left,mid)
* Call merge sort on (mid+1,rear)
* Continue till left is less than right
* Then call merge function to perform merge sort.


