package processor

import (
	"KramerSolve/internal/models"
	"log"
)

// KramerRangFour - function which count matrix equal by Kramer Method
// This is function which counts matrix 4x4
func KramerRangFour(x, y, z, k, res []models.Value) []models.Value {

	x1 := []models.Value{res[0], x[1], x[2], x[3]}
	x2 := []models.Value{x[0], res[1], x[2], x[3]}
	x3 := []models.Value{x[0], x[1], res[2], x[3]}
	x4 := []models.Value{x[0], x[1], x[2], res[3]}

	y1 := []models.Value{res[0], y[1], y[2], y[3]}
	y2 := []models.Value{y[0], res[1], y[2], y[3]}
	y3 := []models.Value{y[0], y[1], res[2], y[3]}
	y4 := []models.Value{y[0], y[1], y[2], res[3]}

	z1 := []models.Value{res[0], z[1], z[2], z[3]}
	z2 := []models.Value{z[0], res[1], z[2], z[3]}
	z3 := []models.Value{z[0], z[1], res[2], z[3]}
	z4 := []models.Value{z[0], z[1], z[2], res[3]}

	k1 := []models.Value{res[0], k[1], k[2], k[3]}
	k2 := []models.Value{k[0], res[1], k[2], k[3]}
	k3 := []models.Value{k[0], k[1], res[2], k[3]}
	k4 := []models.Value{k[0], k[1], k[2], res[3]}

	det := KramerDetFour(x, y, z, k)

	log.Printf("********************************")
	log.Printf("det: %v", det.StringFinal())
	log.Printf("********************************")

	x_res := KramerDetFour(x1, y1, z1, k1)

	log.Printf("********************************")
	log.Printf("x_res: %v", x_res.StringFinal())
	log.Printf("********************************")

	y_res := KramerDetFour(x2, y2, z2, k2)

	log.Printf("********************************")
	log.Printf("y_res: %v", y_res.StringFinal())
	log.Printf("********************************")

	z_res := KramerDetFour(x3, y3, z3, k3)

	log.Printf("********************************")
	log.Printf("z_res: %v", z_res.StringFinal())
	log.Printf("********************************")

	k_res := KramerDetFour(x4, y4, z4, k4)

	log.Printf("********************************")
	log.Printf("k_res: %v", k_res.StringFinal())
	log.Printf("********************************")

	return nil
}

func KramerDetFour(x, y, z, k []models.Value) models.Value {

	A1 := x[0]
	A2 := x[1]
	A3 := x[2]
	A4 := x[3]

	A1 = A1.Multiply(KramerDetThree(y[1:], z[1:], k[1:]))
	A2 = A2.Multiply(KramerDetThree(
		[]models.Value{y[0], y[2], y[3]},
		[]models.Value{z[0], z[2], z[3]},
		[]models.Value{k[0], k[2], k[3]}),
	)
	A3 = A3.Multiply(KramerDetThree(
		[]models.Value{y[0], y[1], y[3]},
		[]models.Value{z[0], z[1], z[3]},
		[]models.Value{k[0], k[1], k[3]}),
	)
	A4 = A4.Multiply(KramerDetThree(
		[]models.Value{y[0], y[1], y[2]},
		[]models.Value{z[0], z[1], z[2]},
		[]models.Value{k[0], k[1], k[2]}),
	)
	A1 = A1.Minus(A2)
	A1 = A1.Plus(A3)
	A1 = A1.Minus(A4)

	return A1
}

// KramerDetThree - function which count matrix deter by Kramer Method
// This is function which counts matrix 3x3
func KramerDetThree(x, y, z []models.Value) models.Value {
	A1 := x[0]
	A1 = A1.Multiply(y[1])
	A1 = A1.Multiply(z[2])

	A2 := x[1]
	A2 = A2.Multiply(y[2])
	A2 = A2.Multiply(z[0])

	A3 := x[2]
	A3 = A3.Multiply(y[0])
	A3 = A3.Multiply(z[1])

	A4 := x[2]
	A4 = A4.Multiply(y[1])
	A4 = A4.Multiply(z[0])

	A5 := x[1]
	A5 = A5.Multiply(y[0])
	A5 = A5.Multiply(z[2])

	A6 := x[0]
	A6 = A6.Multiply(y[2])
	A6 = A6.Multiply(z[1])

	Total := A1
	Total = Total.Plus(A2)
	Total = Total.Plus(A3)
	Total = Total.Minus(A4)
	Total = Total.Minus(A5)
	Total = Total.Minus(A6)
	return Total
}
