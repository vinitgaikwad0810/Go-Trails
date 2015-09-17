package main

import (
	"testing"
	"fmt"
	)

func Test_fibonacci1(t *testing.T){
//0, 1, 1, 2, 3, 5, 8, 13, 21, 34

int_num := Fibo(6)

fmt.Println("ACTUAL RESULT IS ",int_num)
fmt.Println("EXPECTED RESULT IS 8")
if int_num != 8 {
  t.Error("FAILED")
} 
}

func Test_fibonacci2_FAILURE_SCENARIO(t *testing.T){
//0, 1, 1, 2, 3, 5, 8, 13, 21, 34

int_num := Fibo(6)

fmt.Println("ACTUAL RESULT IS ",int_num)
fmt.Println("EXPECTED RESULT IS 9 (FAILURE SCENARIO)")

if int_num != 9 {
  t.Error("FAILED")
} 
}

func Test_fibonacci3_INVALID_INPUT(t *testing.T){
//0, 1, 1, 2, 3, 5, 8, 13, 21, 34

int_num := Fibo(-6)

fmt.Println("ACTUAL RESULT IS ",int_num)
fmt.Println("EXPECTED RESULT IS -1 ( -1 MEANS INVALID INPUT)")

if int_num != -1 {
  t.Error("FAILED")
} 
}

func Test_fibonacci4_PERFOMANCE_TESTING(t *testing.T){
//0, 1, 1, 2, 3, 5, 8, 13, 21, 34
var i int64
fmt.Println("THIS TESTCASE CHECKS IF PROGRAM AND HARDWARE CAN HANDLE HEAVY OPERATIONS")
for i = 0; i < 10; i = i+1 {
int_num := Fibo(i)
fmt.Println("ACTUAL RESULT IS ",int_num)
}

if i != 10 {
  t.Error("FAILED")
} 

}
