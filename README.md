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

## Multihreading

Iterago can execute its functions with a defined number of go routines.
The number of go routine is defined by the environment variable **ITERAGO_THREADS**.

## Examples

### Filter

```go
slice := []uint{1, 2, 4, 8, 16, 32, 64, 128, 256, 512}
// Filters all values upper than 128 and get back a slice of uint
result := iterago.Filter(slice, func(x uint) bool { return x <= 128 })
// Should display [1 2 4 8 16 32 64 128]
fmt.Println(result)
```

### Map

```go
slice := []uint{1, 2, 4, 8, 16, 32, 64, 128, 256, 512}
// Transforms all unsigned int values into a string 
result := iterago.Map(slice, func(x uint) string { return fmt.Sprintf("%d", x) })
// Should display [1 2 4 8 16 32 64 128 256 512]
fmt.Println(result)
```

### Reduce

```go
slice := []uint{1, 2, 4, 8, 16, 32, 64, 128, 256, 512}
// Add up all the unsigned int values 
result := iterago.Reduce(slice, 0, func(x, y uint) uint { return x + y })
// Should display 1023
fmt.Println(result)
```

## Licence

This project is under [MIT](https://opensource.org/licenses/mit-license.php) licence.
