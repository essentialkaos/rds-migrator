//go:build linux
// +build linux

package main

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	APP "github.com/essentialkaos/rds-migrator/app"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// gitrev is short hash of the latest git commit
var gitrev string

// ////////////////////////////////////////////////////////////////////////////////// //

func main() {
	APP.Run(gitrev)
}
