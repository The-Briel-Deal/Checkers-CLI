package main

type Checkerboard struct {
	Name string
}

func (c *Checkerboard) Init(name string) {
	c.Name = name
}