package main

import "fmt"

// Config keeps config for a board
type Config struct {
	shape    int
	baseSide float64
	width    float64
	height   float64
	proto    []string
}

// Examples is a type that collects all example boards
type Examples map[string]*Config

func (ex *Examples) addExamples(newEx ...*Examples) {
	for _, exmpl := range newEx {
		for k, v := range *exmpl {
			(*ex)[k] = v
		}
	}
}

func (ex *Examples) print() {
	for k, v := range *ex {
		fmt.Println(k, "->", v.shape, v.width, v.height, v.baseSide)
	}
}
func loadExamples() *Examples {
	ex := Examples{}
	ex.addExamples(
		examplesSmallEasy,
		examplesSmallNormal,
		examplesSmallHard,
		examplesMediumEasy,
		examplesMediumNormal,
		examplesLargeHard,
	)
	return &ex
}
