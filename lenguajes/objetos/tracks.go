package main

type Interface interface {
	Len() int
	Less(i,j int) bool
	Swap(i, j int)
}

type StringSlice []string

func (p StringSlice) Len () int {return len(p)}
func (p StringSlice) Less (i,j int) bool {return p[i] < p[j]}
func (p StringSlice) Swap (i, j int) {return p[i], p[j] = p[j], p[i]}

