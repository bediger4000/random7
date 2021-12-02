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
2/7\*5/7 of the time after 2 invocations,
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
2/7\*5/7 is approximately .20317,
2/7\*2/7\*5/7 is approximately .05830,
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

In sum, don't use this question in an interview.
Don't feel bad if you get this question in an interview then crash and burn.
Not your fault.

---

# Same problem in reverse

## Daily Coding Problem: Problem #1064 [Easy]

This problem was asked by Two Sigma.

Using a function rand5() that returns an integer from 1 to 5 (inclusive)
with uniform probability,
implement a function rand7() that returns an integer from 1 to 7 (inclusive).

### Analysis

Problem #1064 is the opposite of #761.
I found it considerably harder than #761.

My [first solution](r5.go) is based on the observation that the problem
statement says nothing about the distribution of integers from 1 to 7.
I doubt this is correct.

[Stack Exchange answer](https://stackoverflow.com/questions/137783/expand-a-random-range-from-1-5-to-1-7)

[Reddit answers](https://www.reddit.com/r/compsci/comments/2mhqdz/rand5_from_rand7/)

My [second solution](r5a.go) uses one of the Stack Exchange answers.
It uses 2 invocations of `rand5` as indexes into a 5x5 array.
Values of the array are 1 through 7, except for 4 zeros.
If the code indexes a zero array value, it tries again.

If `rand5` truly returns values with a uniform probability,
then the code will give back 1-7 uniformly.

```sh
$ go build r5a.go
$ ./r5a | sort -n | uniq -c
  71237 1
  71637 2
  71398 3
  71345 4
  71457 5
  71209 6
  71717 7
$
```

Just to put this problem to bed, I wrote a [third solution](rmn.go),
which will create a uniform output 1-M,
from a "rand N" function that returns values 0 through N - 1 uniformly.

```sh
$ go build rmn.go
$ ./rmn 500000 7 5 | sort -n | uniq -c
Should end up with about 71428 of each value
  71905 1
  71430 2
  71734 3
  71590 4
  71029 5
  70931 6
  71381 7
$
```

### Interview analysis

The "rand5 into rand7" question in and of itself is a bad question,
for the same reasons as the "rand7 into rand5" question:
the candidate has to recognize some fact in a small amount of time.
Once recognized, there's not a huge amount of coding.

The general "randN into randM, where M > N" could be a decent interview question,
as long as the interviewer gives some hints.
If the interviewer walks the candidate through two `randN()` values
as indexes into a NxN array with 1-M values,
each value appearing some times,
the candidate could exhibit a lot of coding skills.

* indirection in service of generalization while setting up the NxN array
* remembering to seed the PRNG
* array particulars, like indexing and initialization, off-by-1 recognition
* mastery of division and modulo operators

All of these would show up in setting up the NxN array.
Making one distribution from another merely motivates the work.
