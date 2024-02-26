Example of building a double linked list

This implementation is in Go and uses a Generics to work with many types.  The nodes maintain a double linked list with ordering constraint most recently used to least recently inserted from head to tail. The list can set to enforce a maximum number of nodes or unbounded.

``` 
    head -> [Node *] <--> [Node *] <--> [Node *] <--> [Node *] <--> [Node *] <- tail
            [Blob  ]      [Blob  ]      [Blob  ]      [Blob  ]      [Blob  ]
```

**Methods**
* New
* Append
* Prepend
* PopHead
* PopTail
* Length
* GetIndex
* Print

**Complexity**
* Append, Prepend, PopHead, PopTail, Length, New are O(1)
* GetIndex is O(n) but will search only half the list, from either head or tail whichever is closest
* Print is O(n) - obviously to traverse the structure

**Algorithm**

Each node has two pointers, next and previous, as well as a pointer to a generic data node.  The list has a head and tail pointer as well as the current size and maximimum size.

Adding or removing nodes involves setting the pointers for head and/or tail as well as the node next/previous.

**Limitations**
* Does not currently support multi-threaded access

**Testing**
* 100% code coverage
* Unit tests for boundary conditions, unitialized list etc.
* Functional series of pops and pushes to both head and tail
* Benchmarks