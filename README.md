# GoRefresher

# GoRefresher - Enhanced DSA Study Plan

# DSA Study Plan (Daily Timetable & Interview Guide)

## **📅 Week 1**
- **Mar 2** → Arrays & Strings
- **Mar 3** → Hashing
- **Mar 4** → Linked Lists

## **📅 Week 2**
- **Mar 9** → Stacks & Queues
- **Mar 10** → Recursion & Backtracking
- **Mar 11** → Binary Search

## **📅 Week 3**
- **Mar 16** → Trees
- **Mar 17** → Graphs
- **Mar 18** → Dynamic Programming (DP)

## **📅 Week 4**
- **Mar 23** → Bit Manipulation
- **Mar 24** → Tries
- **Mar 25** → Advanced Data Structures

---

# Data Structures & Algorithms (DSA) Interview Guide

## **1. Arrays & Strings**

| Technique | Description | Solves/Applies To | Example Problems |
|-----------|-------------|-------------------|------------------|
| Two-Pointer | Use two indices to process data efficiently | Finding pairs, reversing, sorting, merging. Efficiently iterating through arrays/strings. | Pair Sum, Dutch Flag |
| subsets | Usefult to find subsets| creating subset of n array 2^n | find sum of subsets equal to sum|
|Permutation|Find how to create permutation  n!|||
|[Lomuto Partition](#lomuto-partition)|used in quick sort|can be used to find top k elements||
|[Hoares Partition](#hoares-partition)|used in quick sort|can be used to find top k elements. It is some what better than lomuto partition||
| Sliding Window | Optimize subarray problems using a window | Finding subarrays/substrings with specific properties (max/min sum, unique chars). Efficient for continuous data. | Longest Substring w/o Repeats |
| Kadane’s Algorithm | Max sum of contiguous subarray | Finding the maximum sum of a contiguous subarray. Used in financial analysis, signal processing. Variation is maximum sum of circular sub array | Maximum Subarray Sum |
| Prefix Sum | Efficient range sum computation | Efficiently calculating sums of subarrays/submatrices. Used in image processing, data analysis. | Range Sum Query |
| Binary Search | Search in sorted or rotated arrays | Finding elements in sorted data, optimizing search time. Used in database systems, search engines. | First/Last Occurrence, Peak Element |
| KMP Algorithm | String pattern matching efficiently | Finding occurrences of a pattern in a string. Used in text editors, search engines, bioinformatics. | Substring Search |
| [**Dutch National Flag Problem** ](#dutch-flag)| Sort an array of 0s, 1s, and 2s in one pass. | Partitioning elements based on a pivot. Used in sorting algorithms, quicksort variations. | Sort Colors |

---

## **2. Hashing**

| Technique | Description | Solves/Applies To | Example Problems |
|-----------|-------------|-------------------|------------------|
| HashMap Usage | Fast lookups & frequency counting | Counting frequency of elements, checking for duplicates, grouping anagrams. Used in caching, database indexing. | Two Sum, Group Anagrams |
| Rolling Hash | Efficient substring search | Finding patterns in strings, plagiarism detection. Used in text search, bioinformatics. | Rabin-Karp Algorithm |
| Trie | Prefix tree for fast lookups | Auto-completion, spell checking, IP routing. Used in search engines, dictionaries. | AutoComplete, Word Search |
| **Bloom Filter** | Probabilistic data structure for set membership testing. | Checking if an element is present in a set with a possibility of false positives. Used in network applications, caching, databases. | Implement Bloom Filter |

---

## **3. Linked Lists**

| Technique | Description | Solves/Applies To | Example Problems |
|-----------|-------------|-------------------|------------------|
| [Slow & Fast Pointers](#fast-slow-pointer) | Detect cycles | Finding cycles in linked lists, finding the middle element. Used in memory management, garbage collection. | Floyd’s Cycle Detection |
| Merge & Reverse | Linked list manipulation | Merging sorted lists, reversing lists. Used in sorting algorithms, memory management. | Merge Sorted Lists, Reverse k-groups |
| LRU Cache | Implement caching using DLL+HashMap | Implementing a Least Recently Used (LRU) cache. Used in web servers, databases. | LRU Cache |
| **Intersection of Two Linked Lists** | Find the intersection node of two linked lists. | Finding common nodes between two lists. Used in database systems, memory management. | Intersection of Two Linked Lists |

---

## **4. Stacks & Queues**

| Technique | Description | Solves/Applies To | Example Problems |
|-----------|-------------|-------------------|------------------|
| Monotonic Stack | Find next greater/smaller elements | Finding the next greater/smaller element in an array, calculating histogram area. Used in stock market analysis, signal processing. | Largest Rectangle in Histogram |
| Sliding Window | Optimize max/min retrieval in a window | Finding the maximum/minimum element in a sliding window. Used in real-time data processing, signal analysis. | Sliding Window Maximum |
| Queue using Stacks | Implement Queue with Stack | Simulating queue behavior using stacks. Useful when stack operations are more readily available. | Two-Stack Queue |
| **Min Stack** | Implementing a stack that supports retrieving the minimum element in O(1) time. | Maintaining a minimum value in a stack. Used in real-time data analysis, system monitoring. | Min Stack |

---

## **5. Recursion & Backtracking**

| Technique | Description | Solves/Applies To | Example Problems |
|-----------|-------------|-------------------|------------------|
| Backtracking | Try all possibilities & prune invalid cases | Solving constraint satisfaction problems, finding all possible solutions. Used in game solving, puzzle solving. | N-Queens, Word Break |
| Subset & Permutation | Generate all subsets & permutations | Generating all possible combinations/arrangements. Used in cryptography, combinatorial problems. | Power Set, Letter Combinations |
| **Combination Sum** | Find all combinations that sum to a target. | Finding combinations of elements that satisfy a given condition. Used in combinatorial problems, optimization. | Combination Sum |

---

## **6. Binary Search**

| Technique | Description | Solves/Applies To | Example Problems |
|-----------|-------------|-------------------|------------------|
| Search in Rotated Array | Handle rotated sorted data | Finding elements in rotated sorted arrays. Used in search engines, database systems. | Find Minimum in Rotated Array |
| Binary Search on Answer | Optimize searching for solutions | Solving optimization problems by searching the solution space. Used in resource allocation, scheduling. | Aggressive Cows, Minimize Max Distance |
| **Find Peak Element II** | Find a peak element in a 2D grid. | Finding local maxima in a multidimensional array. Used in image processing, data analysis. | Find Peak Element II |

---

## **7. Trees & BST**

| Technique | Description | Solves/Applies To | Example Problems |
|-----------|-------------|-------------------|------------------|
| DFS & BFS Traversal | Tree traversal techniques | Traversing trees, finding paths, searching for nodes. Used in file systems, network routing. | Level Order, Inorder Traversal |
| Lowest Common Ancestor | Find common ancestor of nodes | Finding the lowest common ancestor of two nodes in a tree. Used in phylogenetic analysis, file systems. | LCA in BST |
| Serialize & Deserialize | Store & retrieve tree efficiently | Storing and retrieving trees from storage. Used in data serialization, network communication. | Binary Tree Serialization |
| **Validate Binary Search Tree** | Check if a binary tree is a valid BST. | Ensuring the properties of a BST are maintained. Used in data storage, search algorithms. | Validate Binary Search Tree |

---

## **8. Graphs**

| Technique | Description | Solves/Applies To | Example Problems |
|-----------|-------------|-------------------|------------------|
| DFS & BFS | Graph traversal techniques | Traversing graphs, finding connected components, detecting cycles. Used in social networks, web crawlers. | Cycle Detection |
| Shortest Path | Find shortest paths in weighted graphs | Finding the shortest path between two nodes. Used in navigation systems, network routing. | Dijkstra’s Algorithm |
| Topological Sorting | Order tasks with dependencies | Ordering tasks based on dependencies. Used in task scheduling, dependency resolution. | Kahn’s Algorithm |
| **Bellman-Ford Algorithm** | Finds shortest paths in graphs with negative weights. | Handling graphs with negative weight edges. Used in network routing, financial modeling. | Bellman-Ford Algorithm |

---

## **9. Dynamic Programming**

| Technique | Description | Solves/Applies To | Example Problems |
|-----------|-------------|-------------------|------------------|
| 1D DP | Solve problems using previous results | Solving problems that can be broken down into overlapping subproblems. Used in financial analysis, game theory. | Fibonacci, House Robber |
| 2D DP | Grid-based problems | Solving grid-based problems, finding optimal paths. Used in image processing, bioinformatics. | LCS, Edit Distance |
| Knapsack | Optimize resource allocation | Optimizing resource allocation with constraints. Used in operations research, financial planning. | 0/1 Knapsack |
| Matrix Chain Multiplication | Optimize multiplication order | Optimizing the order of matrix multiplications. Used in compiler design, scientific computing. | MCM Problem |
| Egg Dropping Puzzle | Min attempts to find safe floor | Finding the minimum number of attempts to find a threshold. Used in testing, optimization problems. | Egg Drop DP |
| Palindrome Partitioning | Min cuts for palindrome substrings | finding min cuts to make all partitions palindromic. Used in string manipulation, text processing

---

## **10. Bit Manipulation**

| Technique | Description | Example Problems |
|-----------|-------------|------------------|
| XOR Trick | Find missing/non-repeating elements | Single Number |
| Bit Counting | Optimize bit operations | Count Set Bits |

---

## **11. Tries (Prefix Trees)**

| Technique | Description | Example Problems |
|-----------|-------------|------------------|
| Insert & Search | Store words for fast retrieval | AutoComplete, Dictionary Lookup |
| Maximum XOR Pair | Optimize bitwise operations | Trie-based XOR Maximization |

---

## **12. Advanced Data Structures**

| Technique | Description | Example Problems |
|-----------|-------------|------------------|
| Segment Tree | Efficient range queries | Range Sum Query |
| Fenwick Tree | Optimize prefix sums | Binary Indexed Tree |
| Disjoint Set | Connected components in graphs | Union-Find Algorithm |

---

<a id="fast-slow-pointer"></a>**Fast Slow Pointer**

In fast slow pointer slow moves one position at a time and fast moves two position at a time. when both fast and slow meets then there is a cycle. Once slow and fast meets, reset the slow pointer to startig and move both slow and fast at speed of 1, now at whatever the element slow and fast meets is the point where the cycle happens.
With the above code we can use it to find loop,detect cycle, remove cycle etc..

---


<a id="lomuto-partition"></a> 

**Lomuto Partition**

Lomuto partition is used in quick sort, and depend upon the element we take for partition sort can be o(nlogn) in wrost case it could be o(n^2). Lomuto partition puts the partition element in its correct place. Always take partition element as last index and styarting from start,  if element < partition then swap(i,partitionIndex) partitionIndex++ else continue , once every thing is over.

---

<a id="hoares-partition"></a> 

**Hoares Partition**

Hoares Partition is some what more efficient than lomuto, where in hoares does not put the partition element in its correct position.

```
low = low-1 and high = high+1

for {

 do low++ while(arr[low) >p)

 do high-- while (arr[high] <p)
swap(low,high,arr)
if(low >=high){
return high;
}
}
```

---
<a id="dutch-flag"></a> **Dutch Flag Algorithm**

Dutch flag algorithm in used to partition three types of element.here low is low end and mid is mid end and high is high start. low to mid is range where we will have 1

```
low =0
mid =0

for mid <= high{

if(arr[mid] ==0){
 swap(low,mid)
low++
mid++
}
if(arr[mid]==1){
mid++
}

if(arr[mid] ==2{
swap(mid,high)
high--
}

}

```
.

---
