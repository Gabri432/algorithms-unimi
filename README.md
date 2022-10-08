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
BenchmarkLanternFishes18-2       1876839               634.0 ns/op
BenchmarkLanternFishes80-2          7695            137720 ns/op
PASS
ok      github.com/Gabri432/algorithms-unimi/problem_solving    3.895s
```

## Project structure
- (main)
  - license
  - README.md
  - go.mod
  - first_week (folder)
    - ex3.go
    - ex4.go
    - lantern_fishes.go
    - first_week.go

## Links
[Problem Solving Exercises](https://lonati.di.unimi.it/algolab-go/22-23/materiale/settimana01/02-problemSolvingEOsservazioniDichiarative.pdf)

[How to make benchmark testing in go](https://dev.to/mcaci/introduction-to-benchmarks-in-go-3cii)