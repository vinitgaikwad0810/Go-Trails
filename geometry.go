package main

import (
	"fmt"
	"math"
	)

type Shape interface{
area() float64
perimeter() float64
}

type Rectangle struct {
x1, y1, x2, y2 float64
}

type Circle struct {
x, y, r float64
}

func distance(x1, y1, x2, y2 float64) float64 {
var a,b float64
a = x2 - x1
b = y2 - y1
return math.Sqrt(a*a + b*b)
}

func (r Rectangle) area() float64 {
l := distance(r.x1, r.y1, r.x1, r.y2)
w := distance(r.x1, r.y1, r.x2, r.y1)
return l * w
}

func (r Rectangle) perimeter() float64 {
l := distance(r.x1, r.y1, r.x1, r.y2)
w := distance(r.x1, r.y1, r.x2, r.y1)

if l == 0 {
    return -1
}else if w == 0 {
    return -1
}else {
 return 2*(l + w)   
}


}



func (c Circle) area() float64 {
return math.Pi * c.r*c.r
}

func (c Circle) perimeter() float64 {
return 2*math.Pi * c.r
}

func calculate(g Shape){
    fmt.Println(g)
    fmt.Println("Area =",g.area())
    fmt.Println("Perimeter =",g.perimeter())
   
}

func rectangle_Perimeter(x1, y1, x2, y2 float64) float64{
    r := Rectangle{x1, y1, x2, y2}

    float_per := r.perimeter()
    return float_per;
}

func circle_Perimeter(x,y,r float64) float64{
    c := Circle{x, y, r}

    float_per := c.perimeter()
    return float_per;
}
/*
func main(){
    var x1,x2,y1,y2,x,y,rad float64
    fmt.Println("Enter co-ordinates of a rectangle x1,x2,y1,y2")
    fmt.Scanf("%f",&x1)
    fmt.Scanf("%f",&x2)
    fmt.Scanf("%f",&y1)
    fmt.Scanf("%f",&y2)
    fmt.Println("Co-ordniates of the rectangle : X1 =",x1,"\nX2 =",x2,"\nY1 =",y1,"\nY2 =",y2)
    r := Rectangle{x1, y1, x2, y2}
    
    fmt.Println("Enter the co-ordinate and radius of a circle (X,Y,R)")
    fmt.Scanf("%f",&x)
    fmt.Scanf("%f",&y)
    fmt.Scanf("%f",&rad)

    c := Circle{x, y, rad}

    calculate(r)
    calculate(c)
*/


