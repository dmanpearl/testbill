Simple package main

Command line app processes a literal slice of 8 integers and prints the sum of
each group of 3 values.

# Run

    go run main.go

Result:

    [{6 true} {15 true} {15 false}]

Returns an array of groups including the sum of up to 3 values and a boolean
indicating if all 3 values were available.

# Test

    go test main/proc
