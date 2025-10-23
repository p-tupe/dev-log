---
modified: Thu Oct 27 11:56:20 EDT 2022
---
# Graphs

- Collection of nodes/vertices and connecting edges

- Can be (un)directed, (uni|bi)directional, weighted, (un)connected, (a)cyclic.

- Usually represented as a matrix, can also be constructed as an ADT.

- Path-finding (Traversal) Techniques:

  - Depth First Search (Pre-order)
  - Breadth First Search (Post-order)
  - Dijkstra's Algorithm

## DFS Pseudocode

```
current vertex = start vertex

While current vertex exists
  Add it to "visited" list
  Add to "revisit" stack
  If current vertex has edges to vertices not in "visited" list
    current vertex = unvisited vertex
  Else
    current vertex = pop from revisit stack

DFS = visited list
```

## BFS Pseudocode

```
current vertex = start vertex

While current vertex exists
  Add it to "visited" list
  Add to "revisit" queue
  If current vertex has edges to vertices not in "visited" list
    current vertex = unvisited vertex
  Else
    current vertex = dequeue from revisit queue

BFS = visited list
```

> Yes, apparently the only difference is that DFS uses a stack while BFS uses a queue!

## Dijkstra's Algorithm

> Compute shortest distance between any two vertices. Requires graph be weighted.

```
A "distances" dict/map that contains distances from start vertex to rest of the vertices.

Do a BFS from start vertex
  Sum weight to current vertex from start
  If it is less than that in "distances", update "distances" with new path/weight

```
