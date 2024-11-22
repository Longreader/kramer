package main

import (
	"KramerSolve/internal/models"
	"KramerSolve/internal/processor"
)

func main() {

	x := []models.Value{
		models.Value{
			Content: "M_зат",
			Sign:    true,
		},
		models.Value{
			Content: "0",
			Sign:    true,
		},
		models.Value{
			Content: "0",
			Sign:    true,
		},
		models.Value{
			Content: "0",
			Sign:    true,
		},
	}
	y := []models.Value{
		models.Value{
			Content: "M_ств*(r/R)",
			Sign:    true,
		},
		models.Value{
			Content: "M_ств*(1-(r/R))",
			Sign:    true,
		},
		models.Value{
			Content: "1",
			Sign:    false,
		},
		models.Value{
			Content: "0",
			Sign:    true,
		},
	}
	z := []models.Value{
		{
			Content: "0",
			Sign:    true,
		},
		{
			Content: "M_кол",
			Sign:    true,
		},
		{
			Content: "1",
			Sign:    true,
		},
		{
			Content: "1",
			Sign:    false,
		},
	}
	k := []models.Value{
		{
			Content: "J_кол/R",
			Sign:    true,
		},
		{
			Content: "-(J_кол/R)",
			Sign:    true,
		},
		{
			Content: "r",
			Sign:    false,
		},
		{
			Content: "R",
			Sign:    true,
		},
	}
	res := []models.Value{
		{
			Content: "P-П_зат_отк(х_1)",
			Sign:    true,
		},
		{
			Content: "П_ств_отк(х_1,x_2)",
			Sign:    false,
		},
		{
			Content: "П_кол_отк(x_2)",
			Sign:    false,
		},
		{
			Content: "0",
			Sign:    true,
		},
	}

	_ = processor.KramerRangFour(x, y, z, k, res)

}
