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
