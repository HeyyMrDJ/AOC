# Notes
I had issues in part 2 with replacement order not being consistant. One of the reddit solutions was to replace written digits with the first char, number, last char instead of just the number which is what I was attempting. This provided a consistant replacement pattern

```go
    words := map[string]string{
        "one":   "o1e",
        "two":   "t2o",
        "three": "t3e",
        "four":  "f4r",
        "five":  "f5e",
        "six":   "s6x",
        "seven": "s7n",
        "eight": "e8t",
        "nine":  "n9e",
    }
```
