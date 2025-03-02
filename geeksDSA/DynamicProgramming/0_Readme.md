Its an optimization over plain recursion

the idea is to reuse the solution of sub problems

1) Memoization (Top down) recursion
2) Tabulation (Bottom up) Iteration

Applications
1) Bellman Fors Algorithm
2) Floyd warshall Algorithm
3) Diff Utility
4) Search closed words (Edit distamnce)
5) Resource Allocation (0-1 knapsack)

// tabulation is generally faster than memoization.

// when to use DP is when the problem can be divided into subproblem and can be reused
//If you can think of an example where earlier decisions affect future decisions, then DP is applicable
//Problems that have optimal substructure and overlapping subproblems are good problems to be solved with dynamic programming.