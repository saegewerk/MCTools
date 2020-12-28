package geometry

import (
	"fmt"
)

type IntVector3 struct {
	X int
	Y int
	Z int
}
type Cube struct {
	Start IntVector3
	End   IntVector3
}

func (cube *Cube) GetLength() IntVector3 {
	return IntVector3{
		X: cube.End.X - cube.Start.X,
		Y: cube.End.Y - cube.Start.Y,
		Z: cube.End.Z - cube.Start.Z,
	}
}
func (cube *Cube) ToString() string {
	return cube.Start.ToString() + "   " + cube.End.ToString()
}
func (v *IntVector3) ToString() string {
	return fmt.Sprintf("%d %d %d", v.X, v.Y, v.Z)
}
func (cube *Cube) SplitByLimitAndMove(limit int, move IntVector3) ([]Cube, []IntVector3) {
	a := limit
	l := cube.GetLength()
	c := IntVector3{
		X: l.X/a + 1,
		Y: l.Y/a + 1,
		Z: l.Z/a + 1,
	}
	arrCube := make([]Cube, 0)
	arrVector := make([]IntVector3, 0)
	for z := 0; z < c.Z; z++ {
		for y := 0; y < c.Y; y++ {
			for x := 0; x < c.X; x++ {
				arrCube = append(arrCube, Cube{
					Start: IntVector3{
						X: x*a + cube.Start.X,
						Y: y*a + cube.Start.Y,
						Z: z*a + cube.Start.Z,
					},
					End: IntVector3{
						X: (1+x)*a - 1 + cube.Start.X,
						Y: (1+y)*a - 1 + cube.Start.Y,
						Z: (1+z)*a - 1 + cube.Start.Z,
					},
				})
				arrVector = append(arrVector,
					IntVector3{
						X: arrCube[len(arrCube)-1].Start.X + move.X,
						Y: arrCube[len(arrCube)-1].Start.Y + move.Y,
						Z: arrCube[len(arrCube)-1].Start.Z + move.Z,
					},
				)
				if x == c.X-1 {
					arrCube[len(arrCube)-1].End.X = cube.End.X
				}
				if y == c.Y-1 {
					arrCube[len(arrCube)-1].End.Y = cube.End.Y
				}
				if z == c.Z-1 {
					arrCube[len(arrCube)-1].End.Z = cube.End.Z
				}
			}
		}

	}
	return arrCube, arrVector
}
