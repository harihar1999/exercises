package main

import "fmt"

func main() {

	var letters = [5]string{ "quick", "brown", "fox", "lazy", "dog" }
    m := make(map[string]int)
	for _,c := range "abcdefghijklmnopqrstuvwxyz" {
		m[string(c)] =0
	}
		for i :=0; i<5 ;i++ {
           for j:=0; j< len(letters[i]); j++{

              m[string(letters[i][j])]++
             // fmt.Println(j)
		   }



	}
	for _,c := range "abcdefghijklmnopqrstuvwxyz" {
		fmt.Println( string(c), "=>" ,m[string(c)])
	}



}
