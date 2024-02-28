# Example of building a double linked list

This implementation is in Go and uses Generics to work with many types.  The nodes maintain a double linked list with ordering constraint most to least recently inserted starting from head and going to tail. The list can set to enforce a maximum number of nodes or unbounded.

``` text
    head -> [Node *] <--> [Node *] <--> [Node *] <--> [Node *] <--> [Node *] <- tail
            [Blob  ]      [Blob  ]      [Blob  ]      [Blob  ]      [Blob  ]
```

## Methods

* New
* PushTail
* PushHead
* PopTail
* PopHead
* Length
* Seek
* Print

## Complexity

* PushTail, PushHead, PopHead, PopTail, Length, New are O(1)
* Seek is O(n) but will search only half the list, from either head or tail whichever is closest
* Print is O(n) - obviously to traverse the structure

## Algorithm

Each node has two pointers, next and previous, as well as a pointer to a generic data node.  The list has a head and tail pointer as well as the current size and maximimum size.

Adding or removing nodes involves setting the pointers for head and/or tail as well as the node next/previous.

## Limitations

* Does not currently support multi-threaded access
* No option to sort the list in any other way
* Have not implemented an efficient search using hashes/maps/trees for intermediate nodes

## Testing

* 100% code coverage
* Unit tests for boundary conditions, unitialized list etc.
* Functional series of pops and pushes to both head and tail
* Benchmarks