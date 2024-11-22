package main

import (
	"KramerSolve/internal/models"
	"KramerSolve/internal/processor"
	"fmt"
)

func main() {

	x := []models.Value{
		models.Value{
			Content: "H1",
			Sign:    false,
		},
		models.Value{
			Content: "H2",
			Sign:    true,
		},
		models.Value{
			Content: "0",
			Sign:    true,
		},
		models.Value{
			Content: "1",
			Sign:    false,
		},
	}
	y := []models.Value{
		models.Value{
			Content: "B1",
			Sign:    false,
		},
		models.Value{
			Content: "B2",
			Sign:    true,
		},
		models.Value{
			Content: "B3",
			Sign:    true,
		},
		models.Value{
			Content: "B4",
			Sign:    true,
		},
	}
	z := []models.Value{
		{
			Content: "F1",
			Sign:    true,
		},
		{
			Content: "F2",
			Sign:    true,
		},
		{
			Content: "F3",
			Sign:    true,
		},
		{
			Content: "F4",
			Sign:    true,
		},
	}
	k := []models.Value{
		{
			Content: "L1",
			Sign:    false,
		},
		{
			Content: "L2",
			Sign:    false,
		},
		{
			Content: "L3",
			Sign:    false,
		},
		{
			Content: "L4",
			Sign:    false,
		},
	}
	res := []models.Value{
		{
			Content: "P1-I1",
			Sign:    false,
		},
		{
			Content: "0",
			Sign:    true,
		},
		{
			Content: "P3+P3",
			Sign:    true,
		},
		{
			Content: "5",
			Sign:    true,
		},
	}

	result := processor.KramerRangFour(x, y, z, k, res)

	for _, value := range result {
		fmt.Println("***************")
		fmt.Println(value.String())
		fmt.Println("***************")
	}
}
