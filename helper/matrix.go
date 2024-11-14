package helper

import (
	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/gonum/mat"
)

func DataframeToMatrix(df dataframe.DataFrame) *mat.Dense {
	rowCount, colCount := df.Dims()

	dataMatrix := make([]float64, rowCount*colCount)

	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			dataMatrix[i*colCount+j] = df.Elem(i, j).Float()
		}
	}

	return mat.NewDense(rowCount, colCount, dataMatrix)
}

func GetReciprocalRowNorms(matrix *mat.Dense) []float64 {
	rows, _ := matrix.Dims()
	reciprocalNorms := make([]float64, rows)
	for i := 0; i < rows; i++ {
		row := matrix.RowView(i)
		norm := mat.Norm(row, 2) // Norma L2
		if norm != 0 {
			reciprocalNorms[i] = 1 / norm
		} else {
			reciprocalNorms[i] = 0
		}
	}
	return reciprocalNorms
}

func NewDiagMatrix(diagonals []float64) *mat.Dense {
	n := len(diagonals)
	diag := mat.NewDense(n, n, nil)
	for i := 0; i < n; i++ {
		diag.Set(i, i, diagonals[i])
	}
	return diag
}

func CalculateCosSimilarity(M *mat.Dense) *mat.Dense {
	rowDim, _ := M.Dims()

	reciprocalRowNorms := GetReciprocalRowNorms(M)

	Dn := NewDiagMatrix(reciprocalRowNorms)

	// Calculate dot product row i with row j
	MMT := mat.NewDense(rowDim, rowDim, nil)
	MMT.Mul(M, M.T())

	// Divide with norm of row i
	temp := mat.NewDense(rowDim, rowDim, nil)
	temp.Mul(Dn, MMT)

	// Divide with norm of row j
	simMatrix := mat.NewDense(rowDim, rowDim, nil)
	simMatrix.Mul(temp, Dn)

	return simMatrix
}
