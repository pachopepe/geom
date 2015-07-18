/*
Package geom holds geometry objects that can be encoded, decoded,
and operated on by the other packages in this repository.
*/
package geom

type T interface {
	Bounds(*Bounds) *Bounds
}

type Linear interface {
	T
}

type Polygonal interface {
	T
}

type PointLike interface {
	T
}
