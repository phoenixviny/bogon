
package main

import "fmt"
var	bet =[10]byte{'a','b','c','f','g'}
var sl1,sl2 []byte

func main() {
	a	:= [3]int{1,2,3}
	b	:= [5]int{1,2,3}
	c	:= [...]int{4,5,6}
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	fmt.Printf("The first element is %d\n", a[2]) 
	fmt.Printf("The first element is %d\n", b[4]) 
	fmt.Printf("The first element is %d\n", c[0]) 
	fmt.Printf("The first element is %d\n", doubleArray[0][3]) 
	fmt.Printf("The first element is %d\n",	doubleArray[1][3])
	fmt.Println("Hello World!")
	sl1 = bet[2:5]
	sl2 = bet[0:3]
	fmt.Printf("The first element is %d\n",sl2[0])
	fmt.Printf("The first element is %d\n",sl1[2])
	sum := 0
	for index:=0 ;index <=10; index ++{
		sum += index
	}
	fmt.Printf("The first element is %d\n",sum)
	for plex:=10;plex >0;plex --{
		if plex==5 {
			break
		}
	fmt.Printf("The first element is %d\n",plex)
	}
	integer := 6
	switch integer {
	case 4:
		fmt.Println("The integer was <= 4")
		fallthrough
	case 5:
		fmt.Println("The integer was <= 5")
		fallthrough
	case 6:
		fmt.Println("The integer was <= 6")
		fallthrough
	case 7:
		fmt.Println("The integer was <= 7")
		fallthrough
	case 8:
		fmt.Println("The integer was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
	



}