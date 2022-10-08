# Algorithms-unimi
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Gabri432/algorithms-unimi)

A repository containing some 'Algorithms and Data Structures' course exercises of professor Violetta Lonati, Unimi.

## How to run the tests
- Run `go test` or `go test -v` to make unit testing and see if the output is correct;
- Run `go test -bench .` to make benchmark testing and see the average execution time per iteration;

## Example
- go test -bench .
```
goos: windows
goarch: amd64
pkg: github.com/Gabri432/algorithms-unimi/problem_solving
cpu: AMD Athlon Silver 3050U with Radeon Graphics
BenchmarkLanternFishes-2         1837238               641.9 ns/op
PASS
ok      github.com/Gabri432/algorithms-unimi/problem_solving    2.166s
```

## Project structure
- (main)
  - license
  - README.md
  - go.mod
  - problem_solving (folder)
    - ex3.go
    - ex4.go
    - lantern_fishes.go
    - problem_solving_test.go

## Links
[Problem Solving Exercises](https://lonati.di.unimi.it/algolab-go/22-23/materiale/settimana01/02-problemSolvingEOsservazioniDichiarative.pdf)

[How to make benchmark testing in go](https://dev.to/mcaci/introduction-to-benchmarks-in-go-3cii)