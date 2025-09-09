// This code shows how to use local packages in a module.
//
// `go mod init <package URL>` is a must, since in go
// all packages are referenced using that url.
//
// Here, we use "packages" as a URL but regardless,
// it works. See here for more:
//   - https://go.dev/doc/code#ImportingLocal
//   - https://stackoverflow.com/questions/52026284/accessing-local-packages-within-a-go-module-go-1-11
package main

import (
	u "packages/utils"
	n "packages/utils/nested"
)

func main() {
	// This code is implemented in other package
	println(u.MagicNumber())
	println(n.MagicString())
}
