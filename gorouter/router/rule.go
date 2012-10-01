// Package  provides ...
package router

// parse rule interface, any parse need to implement it
type Rule interface {
    Parse(string) (string, string)
}

