# Iterago

[![push_master](https://github.com/ulphidius/iterago/actions/workflows/push_master.yml/badge.svg)](https://github.com/ulphidius/iterago/actions/workflows/push_master.yml)
[![codecov](https://codecov.io/gh/ulphidius/iterago/branch/master/graph/badge.svg?token=AG5HTRG4W4)](https://codecov.io/gh/ulphidius/iterago)
[![Go Report Card](https://goreportcard.com/badge/github.com/ulphidius/iterago)](https://goreportcard.com/report/github.com/ulphidius/iterago)
[![Go Reference](https://pkg.go.dev/badge/github.com/ulphidius/iterago.svg)](https://pkg.go.dev/github.com/ulphidius/iterago)

`iterago` is an iterator library.

## Installation

```
go get github.com/ulphidius/iterago
```

## Examples

### Filter

```go
slice := []uint{1, 2, 4, 8, 16, 32, 64, 128, 256, 512}
// Converts the slice into an option of iterator
potentialIter := iterago.SliceIntoIter(slice)
// Get the iterator and ignores the potential error
iterator, _ := potentialIter.Unwrap()
// Filters all values upper than 128 and get back a slice of uint
result := iterator.Filter(func(x uint) bool { return x <= 128 }).
    Collect()
// Should display [1 2 4 8 16 32 64 128]
fmt.Println(result)
```

### Map

```go
slice := []uint{1, 2, 4, 8, 16, 32, 64, 128, 256, 512}
// Converts the slice into an option of iterator
potentialIter := iterago.SliceIntoIter(slice)
// Get the iterator and ignores the potential error
iterator, _ := potentialIter.Unwrap()
// Transforms all unsigned int values into a string 
result := iterator.Map(func(x uint) any { return fmt.Sprintf("%d", x) }).
    Collect()
// Should display [1 2 4 8 16 32 64 128 256 512]
fmt.Println(result)
```

### Reduce

```go
slice := []uint{1, 2, 4, 8, 16, 32, 64, 128, 256, 512}
// Converts the slice into an option of iterator
potentialIter := iterago.SliceIntoIter(slice)
// Get the iterator and ignores the potential error
iterator, _ := potentialIter.Unwrap()
// Add up all the unsigned int values 
result := iterator.Reduce(func(x, y uint) uint { return x + y })
// Should display 1023
fmt.Println(result)
```

### Chained function

```go
slice := []uint{1, 2, 4, 8, 16, 32, 64, 128, 256, 512}
// Converts the slice into an option of iterator
potentialIter := iterago.SliceIntoIter(slice)
// Get the iterator and ignores the potential error
iterator, _ := potentialIter.Unwrap()
// Filters all values lower or equal than 16 and upper than 128 and add up all the values
result := iterator.Filter(func(x uint) bool { return x <= 128 }).
    Filter(func(x uint) bool { return x >= 16 }).
    Reduce(func(x, y uint) uint { return x + y })
// Should display 240
fmt.Println(result)
```

## Licence

This project is under [MIT](https://opensource.org/licenses/mit-license.php) licence.
