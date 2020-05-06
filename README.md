# Path Calculator

This is a kata for determining the optimal path between two points.

The problem and solution is described below.


## Problem Summary

Suppose you have two parallell paths that take you to the same destination.
The two paths also have crossroads between them.
Traversing each segment or crossroad has a unique fixed cost associated with it.

Assuming you are able to start on either path with no pre-existing cost, how can you determine the cheapest way to traverse to your destination?

### Example path

```
    50      5     40     10
A ------1------2------3------4
        |      |      |      |
      30|    20|    25|     0|
        |      |      |      |
B ------1------2------3------4
    10     90      2      8
```

## Solution

This problem can be broken down into a few chunks.

Start by identifying the cost to a the first point in the path, A1 above.
This can be done by crossing forward to A1 for a total cost of 50.
It can aslo be done by crossing forward to B1 (10) and then taking a crossroad (30) to A1 for a total cost of 40.
Keep the lower cost (40).

Then repeat these steps for B1 resulting in a cost of 10.


Now let's repeat everything above to calculate the cheapest path to A2 and B2.


Now let's repeat everything above to calculate the cheapest path to A2 and B2.
Let's do this for A2 first by just calulating the cheapest path assuming we are at A1 or B1.
The cost from A1 to A2 is 5.
The cost from B1 to A2 is 110 (90 + 20).

Before we decided which cost is lower, we must first also add the cost we have already calculated to get to A1 or B1.
The cost from our starting point  using A1 is 5 plus the cheapest path to A1 (40): 45.
The cost from our starting point  using B1 is 110 plus the cheapest path to B1 (10): 110.
Thus the cheapest cost to A2 is 45 which we keep.

Then repeat this for B2 to find the cheapest cost to be 100.

### Solution Summary

So if we were to take this problem all at once.
We could ask ourselves what is the shortest path to A4?

Written in pseudocode, where n is the number of segments in the paths, it would look something like this:

```
func optimal_path_to(An) {
    minimum_between {
        cost_between_An_and_A(n-1) + optimal_path_to(A(n-1))
        cost_between_Bn_and_B(n-1) + cost_of_crossroads_between_Bn_and_An + optimal_path_to(B(n-1))
    }
}
```

Then in order to find the shortest distance from start to destination,
you simply take the minimum of the optimal path to the destination on point path A and the optimal path to the destination on point path B.


Written in pseudocode, where n is the number of segments in the paths, it would look something like this:

```
func optimal_path_to_destination() {
    minimum_between {
        optimal_path_to(An)
        optimal_path_to(Bn)
    }
}
```

This solution can be implemented differently depending on language.
However, if you are thinking functionally, it essentially equates to using a [left fold](https://en.wikipedia.org/wiki/Fold_(higher-order_function)) to walk the path from left to right while accumulating the best path to An and Bn.


> Credit for the original idea for this exercise goes to the book [Learn You a Haskell for Great Good](http://learnyouahaskell.com/functionally-solving-problems#heathrow-to-london)
