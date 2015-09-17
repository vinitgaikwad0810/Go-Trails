package main

import (
	"testing"
	"fmt"
	)

func Test_Geometry1(t *testing.T){

float_num := rectangle_Perimeter(1,2,3,4)

fmt.Println("ACTUAL RESULT IS ",float_num)
fmt.Println("EXPECTED RESULT IS 8")

if float_num != 8 {
	t.Error("FAILED")
}
}
func Test_Geometry1_Circle(t *testing.T){

float_num := circle_Perimeter(1,2,4)

fmt.Println("ACTUAL RESULT IS ",float_num)
fmt.Println("EXPECTED RESULT IS 25.132741228718345")

if float_num != 25.132741228718345 {
	t.Error("FAILED")
}


}

func Test_Geometry2_FAIL(t *testing.T){

float_num := rectangle_Perimeter(1,2,3,4)

fmt.Println("ACTUAL RESULT IS ",float_num)
fmt.Println("EXPECTED RESULT IS 4 (This is a failure scenario since actual perimeter is 8");
if float_num != 4 {
	t.Error("FAILED")
}
}

func Test_Geometry2_FAIL_Circle(t *testing.T){

float_num := circle_Perimeter(1,2,4)

fmt.Println("ACTUAL RESULT IS ",float_num)
fmt.Println("EXPECTED RESULT IS 25.132741228718345. This test case should fail as we are not checking for the EXPECTED value.")

if float_num != 26.132741228718345 {
	t.Error("FAILED")
}
}

func Test_Geometry3_FLOAT(t *testing.T){

float_num := rectangle_Perimeter(100.34,234.12,323.67,423.12)

fmt.Println("ACTUAL RESULT IS ",float_num)
fmt.Println("EXPECTED RESULT IS 824.6600000000001");
if float_num != 824.6600000000001 {
	t.Error("FAILED")
}
}

func Test_Geometry4_INVALIDINPUT(t *testing.T){

float_num := rectangle_Perimeter(0,0,1,0)

fmt.Println("ACTUAL RESULT IS ",float_num)
fmt.Println("EXPECTED RESULT IS -1 which signifies invalid input as Length/Breadth of a rectangle can not be 0");
if float_num != -1 {
	t.Error("FAILED")

}
}


func Test_Geometry5_PERFORMANCE(t *testing.T){

var sum uint

for i := 0.0; i < 10.0; i = i+1.0 {
		float_num := rectangle_Perimeter(1,2,3+i,4+i)
		fmt.Println("ACTUAL RESULT FOR i =",i ,"IS ",float_num)
	sum += 1
	}

fmt.Println("EXPECTED RESULT IS -1 which signifies invalid input as Length/Breadth of a rectangle can not be 0");
if sum != 10 {
	t.Error("FAILED")

}
}
