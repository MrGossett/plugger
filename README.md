# plugger

This is a POC of using Golang's plugin system.

A simplistic implementation of a Branch-and-Bound solver is implemented in the `provider` package. Notice that `provider/main.go` defines `package main` but does not define a `func main()`; this is the signature of a Golang plugin.

The `provider` package exports a `func Solve` to be used by consumers.

The `consumer` package utilizes `provider.Solve` via a plugin. It defines a vanilla worker assignment problem.

`package shared` defines an interface that is shared between `package provider` and `package consumer`, and also provides a helper that knows how to load an `.so` as a Golang plugin and lookup `func Solve` in its symbols.

### Usage

`make run` from the project root will:

- compile the `provider` package into an `.so`
- compile the `consumer` package as a regular executable binary
- run the `consumer` binary, passing the location of the `.so`

### Output

```bash
$ make run
consumer/consumer -plugin consumer/provider.so
2020/03/04 12:55:43 total cost is 16
2020/03/04 12:55:43 assign worker 9 to job 3
2020/03/04 12:55:43 assign worker 8 to job 1
2020/03/04 12:55:43 assign worker 7 to job 8
2020/03/04 12:55:43 assign worker 6 to job 6
2020/03/04 12:55:43 assign worker 5 to job 5
2020/03/04 12:55:43 assign worker 4 to job 2
2020/03/04 12:55:43 assign worker 3 to job 9
2020/03/04 12:55:43 assign worker 2 to job 4
2020/03/04 12:55:43 assign worker 1 to job 7
2020/03/04 12:55:43 assign worker 0 to job 0
```
