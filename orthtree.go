package main

type numeric interface {
	int | int8 | int16 | int32 | int64
	uint | uint8 | uint16 | uint32 | uint64
	float32 | float64
	complex64 | complex128
}

type Point[N numeric] interface {
	getDimension(d uint64) N
	getDimensions() int64
}

type node[N numeric] struct {
	center   Point[N]
	data     Point[N]
	parent   *node[N]
	children []node[N]
	level    int64
	isLeaf   bool
}

type Orthtree[N numeric] struct {
	size N
	root *node[N]
}

func Create[N numeric](size N) Orthtree[N] {
	tree := Orthtree[N]{size: size}
	return tree
}

func (t *Orthtree[N]) At() *Point[N] {

}

func (t *Orthtree[N]) Insert(ps ...*Point[N]) {

}

func (t *Orthtree[N]) Delete(ps ...*Point[N]) {

}

func (t *Orthtree[N]) Swap(p1 *Point[N], p2 *Point[N]) {

}

func (t *Orthtree[N]) And(ct *Orthtree[N]) Orthtree[N] {

}

func (t *Orthtree[N]) Or(ct *Orthtree[N]) Orthtree[N] {

}

func (t *Orthtree[N]) Not(ct *Orthtree[N]) Orthtree[N] {

}

func (t *Orthtree[N]) Equals(ct *Orthtree[N]) bool {

}

func (t *Orthtree[N]) NearestNeighbor(p *Point[N], radius float64) *Point[N] {

}

func (t *Orthtree[N]) Depth() int64 {

}

func (t *Orthtree[N]) NodeCount() int64 {

}
