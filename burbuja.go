package main

import "fmt"

func Burbuja(s []int64) []int64 {
	fmt.Println("Hola")
	for i:=1; i <len(s); i++{
		for j:=0; j<len(s)-i; j++{
            if(s[j] > s[j+1]){
				k := s[j+1]
                s[j+1] = s[j]
                s[j] = k;
			}
		}	
	}	
	return s    

}

func main()  {

	s := []int64{1, -10, 90, 14, -100, -2}
	fmt.Println(s)
	s= Burbuja(s)
	fmt.Println(s)






	
}