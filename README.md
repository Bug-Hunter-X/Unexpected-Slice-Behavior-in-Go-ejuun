# Unexpected Slice Behavior in Go

This repository demonstrates a common pitfall in Go related to slice behavior and how to avoid it.  The `bug.go` file showcases how modifying a slice after a function returns does not affect the returned slice.  The `bugSolution.go` file provides a corrected version that returns the modified data.

## Problem

In Go, slices are passed by value, but the value is a pointer to the underlying array data.  The `bug.go` file shows a function `myFunc` that doubles the values of a slice; however, changes to the original slice after calling `myFunc` do not affect the slice `myFunc` returned. This can be very confusing, because one may think the function is returning a reference to the original array, but it's just returning a copy of the slice header which points to the same data. The key here is that the slice itself is copied, not the underlying array.

## Solution

The `bugSolution.go` file demonstrates the correct way to handle this. Instead of modifying the original array, we create a new array and return that. This ensures that modifications to the original slice do not affect the results.
