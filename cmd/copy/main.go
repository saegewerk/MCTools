package main

import "github.com/saegewerk/MCTools/pkg/geometry"

func main() {
	cube := geometry.Cube{
		Start: geometry.IntVector3{
			X: -1451,
			Y: 142,
			Z: -613,
		},
		End: geometry.IntVector3{
			X: -1409,
			Y: 159,
			Z: -571, //-1446,-574,,, -1479,-448.......-33, 126
		},
	}
	arrC, arrV := cube.SplitByLimitAndMove(32, geometry.IntVector3{
		X: -33,
		Y: 0,
		Z: 126,
	})
	for i, cube := range arrC {
		print("/clone ")
		print(cube.ToString() + "   ")
		print(arrV[i].ToString())
		print(" replace move")
		println()
	}
}
