# Simple Deterministic Automaton implementation in go

## Intro

This automaton detects if a binary string ends with the
sequence `01`.

So the sequences: `001`, `10010101`, `01001`, `1111101` are validated by it.

## State Machine

![](./images/01.png)


|   | 0 | 1 |
| ------- | ----- | ---|
|-> q1| q2 | q1|
| q2 | q2 | q3 |
|* q3 | q2 | q1 |


## Demo

Running against a valid string

![](images/02.png)


Running againt an invalid string.  

![](images/03.png)

