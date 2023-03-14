package main

import (
	"fmt"
	lib "github.com/lms/lms_parser"
	"runtime"
	"testing"
)

func Add() {
	library := lib.NewLibrary()
	for i := 0; i < 10000; i++ {
		name := fmt.Sprintf("some%v", i)
		book := lib.NewBook(name, name)
		library.Add(book)
	}

}

// Test for how many value can add to Library

func TestAdding(t *testing.T) {
	Add()
	runtime.GC()
}

// For bench adding and removing properly working.
// This also check memory overhead
func TestAddingAndRemoving(t *testing.T) {
	library := lib.NewLibrary()
	for i := 0; i < 10000; i++ {
		name := fmt.Sprintf("some%v", i)
		book := lib.NewBook(name, name)
		library.Add(book)
	}
	for i := 0; i < 10000; i++ {
		name := fmt.Sprintf("some%v", i)
		library.Get(name, name)
	}
}
