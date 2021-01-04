# Daily Coding Problem: Problem #761 [Easy]

This problem was asked by Two Sigma.

Using a function rand7() that returns an integer from 1 to 7 (inclusive)
with uniform probability,
implement a function rand5() that returns an
integer from 1 to 5 (inclusive).

## Build and Run

```sh
$ go build r7.go
$ ./r7 | sort | uniq -c
 100338 0
  99750 1
 100066 2
  99705 3
 100141 4
$
```

`r7` invokes `rand5` 500,000 times.
That's as close to uniform as you could expect.

## Analysis

[Code](r7.go)

I think solving the problem hinges on recognizing that
`rand7` returns uniformly-distributed numbers.
The `rand5` function should call `rand7`
while that return value is greater than 4.
It should return whatever it gets that's the first
value 4 or under.

If `rand7` returns uniformly distributed numbers,
so will such a `rand5` function,
it just "cuts off" the uniformly-distributed-values
that don't fit the 0-4 range.

My algorithm doesn't have a bounded run time:
it's entirely possible to get a series of values
of 5 or 6 from the `rand7` function,
which would cause my function to never return.
It's certainly possible to calculate that `rand5`
returns 5/7 of the time after 1 invocation of `rand7`,
2/7*5/7 of the time after 2 invocations,
but I did an empirical investigation by writing
a [version](r7a.go) that outputs the number of invocations
of `rand7` it did to get to a 0-4 valued output of `rand5`

```sh
$ go build r7a.go
$ ./r7a | sort -k1.1n | uniq -c | awk '{printf "%4d  %.5f\n", $2, $1/500000.}'
```

|rand7 invocations|Proportion of rand5 invocations|Running proportion total|
|------:|---------:|--------:|
|   1|0.71435|0.71435|
|   2|0.20300|0.91736|
|   3|0.05899|0.97634|
|   4|0.01681|0.99315|
|   5|0.00492|0.99807|
|   6|0.00138|0.99945|
|   7|0.00038|0.99983|
|   8|0.00014|0.99997|
|   9|0.00003|0.99999|
|  10|0.00000|1.00000|
|  11|0.00000|1.00000|

5/7 equates to approximately .71428,
2/7*5/7 is approximately .20317,
2/7*2/7*5/7 is approximately .05830,
and so forth.
It checks out.
Seems like there's a vanishingly 
small number of invocations of `rand5` that will invoke `rand7`
more than, say, 15 times.


## Interview Analysis


Asking a problem that hinges on recognizing some specific
aspect of a non-programming question seems absurd.
The coding involved is very easy,
seeing it done won't get the interviewer closer to figuring
out if the candidate is any good.

Similarly, the candidate can't fault themself for not figuring
out some obscure factoid while in a very stressful situation.

In sum, don't use this question in an interview,
and don't feel bad if you get this question in an interview
then crash and burn. Not your fault.
