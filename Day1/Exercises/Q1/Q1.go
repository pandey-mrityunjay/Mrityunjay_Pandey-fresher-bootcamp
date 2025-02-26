package main

import(
	"encoding/json"
	"fmt"
)

type Matrix struct{
	Rows int
	Cols int
	Elements [][]int
}

func(m Matrix) GetRows()int{
	return m.Rows
}

func(m Matrix) GetCols()int{
	return m.Cols
}

// Method to set an element at (i, j)
func (m *Matrix) SetElement(i, j, value int) {
	if i < 0 || i >= m.Rows || j < 0 || j >= m.Cols {
		fmt.Println("Invalid index")
		return
	}
	m.Elements[i][j] = value
}
// Method to add two matrices
func (m Matrix) AddMatrix(other Matrix) Matrix {
	result := Matrix{Rows: m.Rows, Cols: m.Cols, Elements: make([][]int, m.Rows)}
	for i := range result.Elements {
		result.Elements[i] = make([]int, m.Cols)
		for j := range result.Elements[i] {
			result.Elements[i][j] = m.Elements[i][j] + other.Elements[i][j]
		}
	}
	return result
}
// Method to print matrix as JSON
func (m Matrix) ToJSON() {
	jsonData, _ := json.MarshalIndent(m, "", "  ")
	fmt.Println(string(jsonData))
}
// Main function
func main() {
	// Creating a 2x2 matrix
	m1 := Matrix{
		Rows:    2,
		Cols:    2,
		Elements: [][]int{{1, 2}, {3, 4}},
	}
	m2 := Matrix{
		Rows:    2,
		Cols:    2,
		Elements: [][]int{{5, 6}, {7, 8}},
	}
	// Adding matrices
	m3 := m1.AddMatrix(m2)
	// Print matrix in JSON format
	m3.ToJSON()
}