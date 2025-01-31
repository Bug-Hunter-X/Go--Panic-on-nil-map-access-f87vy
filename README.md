# Go: Panic on nil map access

This repository demonstrates a common error in Go: panicking due to accessing a nil map. The `bug.go` file shows the problematic code, while `bugSolution.go` presents a safe alternative.

## Problem

Attempting to access a Go map before initializing it results in a runtime panic.  This often happens when a map is declared but not initialized before use.  The program crashes unexpectedly.

## Solution

The solution involves checking if the map is nil before accessing it.  This can be done using either an explicit `if` statement or the more concise blank assignment check (explained in `bugSolution.go`).