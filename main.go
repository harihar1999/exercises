package main

import "fmt"

type matrix struct {
	row int
	col int
	element[][] int
}


func (m matrix ) getrow() int{
	return m.row
}
func (m matrix ) getcol() int{
	return m.col
}

func ( m matrix ) setvalue (i,j,val int) int{
	m.element[i][j]=val
	return val
}

func (m matrix) addmatrix ( m1[][] int ) [][] int{
	for i:=0; i<m.getrow(); i++ {
		for j:=0; j<m.getcol(); j++ {
			m1[i][j]+=m.element[i][j]
		}
	}
    return m1

}


func main() {
	m := matrix {3,3,[][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}}
	fmt.Println(m.getrow())
	fmt.Println(m.getcol())
	fmt.Println(m.setvalue(0,2,5))
	fmt.Println(m.addmatrix( [][]int{ {1, 2, 3},{4, 5, 6},{7, 8, 9} } )



}