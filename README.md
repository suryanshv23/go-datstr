[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) [![Go Report Card](https://goreportcard.com/badge/github.com/suryanshv23/datstr)](https://goreportcard.com/report/github.com/suryanshv23/datstr)
# datstr
The implementation of package is very simple and can be used with a little knowledge of the programming langage and data structure.

#### Install by running following command in your command prompt or bash :
    go get -u "github.com/suryanshv23/datstr" 

### It contains three sub packages dedicated to different branches of data Structure :

- Search : It contains different function for searching based on different searching algorithm which you can use to search for an element in your array. the functions will return the index of the element and bool value as true if present otherwise -1 and false. The different types of searches in this package are :
    
    - Linear Search
    - Binary Search
    - Interpolation Search
    
- Sort : It containns different types of sorting functions based on different sorting algorithms which you can use to sort an array. the functions will return the sorted array. The different types of sorting in this package are :
    
    - Bubble Sort
    - Insertion Sort
    - Quick Sort
    - Selection Sort
    - Shell Sort
    - Comb SOrt
    
- StorageDS : It contains different functions to implement different Data Structures and use them through different methods. The different types of storage data structure is :
    - Binary Tree
    - Graph (using Adjacency List)
    - Linked List
    - Doubly Linked List
    - Circular Queue (using Array)
    - Queue (using Array)
    - Queue (using Linked List)
    - Stack (using Array)
    - Stack (using Linked List)
    - list (Using Slice)
    
To use different sub packages you have to import these sub-packages in your .go file where code is written using the below path in your imports (Note - After installing datstr package using above command).

To import Search -

    "github.com/suryanshv23/datstr/Search"
    
To import Sort - 

    "github.com/suryanshv23/datstr/Sort"
    
To import StorageDS - 

    "github.com/suryanshv23/datstr/StorageDS"
    
### Note : Anyone who wants to raise issues or contribute to the project are invited.
