/*
Working with multiple packages and how to import 3rd party packages


//Download new package


download flect : go get github.com/gobuffalo/flect
doanload flect latest version : go get -u github.com/gobuffalo/flect

How to import downloaded package

import flect

how to change package name

import customName "package"

import KDFlect "flect"
*/

package main

import k "fmt"

func main() {
	k.Println("fmt name changed.now we can user k.print")
}
