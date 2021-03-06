# Daily Coding Problem: Problem #598 [Easy]

## Problem

This problem was asked by Microsoft.

You have n fair coins and you flip them all at the same time. Any that come up
tails you set aside. The ones that come up heads you flip again. How many
rounds do you expect to play before only one coin remains?

Write a function that, given n, returns the number of rounds you'd expect to
play until one coin remains.

## Analysis

The problem statement emphasises "fair coins",
which appears to have a [precise technical meaning](https://en.wikipedia.org/wiki/Fair_coin).
What about unfair coins?
Looks like it's [surprisingly hard](https://izbicki.me/blog/how-to-create-an-unfair-coin-and-prove-it-with-math.html)
to actually make a biased coin, at least by bending the coin.
If you limit yourself to disk (short cylinder) coins,
it just might be [impossible](http://www.stat.columbia.edu/~gelman/research/published/diceRev2.pdf).

Based on eliminating half the coins at each mass flip,
you'd expect to make log<sub>2</sub>(n) rounds before only one 
coin remains.
This follows the same reasoning that lets us say
that binary search makes log<sub>2</sub>(n) comparisons before
finding the searched-for item.

You could also reason that (most times) half the coins turn up heads,
and thus get eliminated.
How many halvings does it take to get to 1?

I went on to write [a program](main.go) that tells you log<sub>2</sub>(n),
and then does 10,000 iterations of the rounds of flipping n coins
as described above.

```
% go build faircoins.go
% ./faircoins 256
256 fair coins, log2(256) = 8.000000
10000 iterations gives 7.879900 average attempts
```

The simulation always comes out lower than log<sub>2</sub>(n),
even if you do many more than 10,000 iterations of rounds of coin flipping.
I really don't know why.
The simulation can "overshoot" (more rounds than predicted) as well as "undershoot"
(fewer rounds than predicted), so it's not a systemic problem.
The only thing I can think of is that `rand.Intn(2)`
is very slightly biased.
So I wrote [a program ](intntest.go) to test that.
It's not biased that I can tell.

As far as interview questions go,
it won't test any programming knowledge.
The question is phrased in a way that the candidate gets
nudged towards maybe calculating log<sub>2</sub> of numbers.
Unless the position is for numerical analyst,
you don't want this.
Even if someone writes the same sort of simulation I did,
there's very little programming,
no room for the candidate to riff on test cases,
corner cases or potential optimizations.
It doesn't even demand much error handling.

If you used it as one question of several,
it might work,
but otherwise,
a silly question for an interview.
