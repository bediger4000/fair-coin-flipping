# Daily Coding Problem: Problem #598 [Easy]

## Problem

This problem was asked by Microsoft.

You have n fair coins and you flip them all at the same time. Any that come up
tails you set aside. The ones that come up heads you flip again. How many
rounds do you expect to play before only one coin remains?

Write a function that, given n, returns the number of rounds you'd expect to
play until one coin remains.

## Analysis

Based on eliminating half the coins at each mass flip,
you'd expect to make log<sub>2</sub>(n) rounds before only one 
coin remains.
This follows the same reasoning that lets us say
that binary search makes log<sub>2</sub>(n) comparisons before
finding the searched-for item.

You could also reason that (most times) half the coins turn up heads,
and thus get eliminated.

I went on to write [a program](main.go) that tells you log<sub>2</sub>(n),
and goes on to do 10,000 iterations of the rounds of flipping n coins
as described above.

```
% ./faircoins 256
256 fair coins, log2(256) = 8.000000
10000 iterations gives 7.879900 average attempts
```

The simulation always comes out lower than log<sub>2</sub>(n),
even if you do many more than 10,000 iterations of rounds of coin flipping.
I really don't know why.
The simulation can "overshoot" (more rounds) as well as "undershoot"
(fewer rounds than predicted), so it's not a systemic problem.
The only thing I can think of is that `rand.Intn(2)`
is very slightly biased.
So I wrote [a program ](intntest.go) to test that.
It's not biased that I can tell.
