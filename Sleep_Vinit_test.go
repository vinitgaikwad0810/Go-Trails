package main

import (
	"testing"
	"fmt"
	"time"
	)

func Test_Sleep1_SUCCESS_SCENARIO(t *testing.T){

	
	Before := time.Now().Second()
	fmt.Println("Before putting the golang program to sleep  .....nth Second at this instance is ",Before)

	Sleep_Vinit(5)
	After := time.Now().Second()
	fmt.Println("After program wakes up  .....nth Second at this instance \n",After)

	if(After - Before) !=5 {
		t.Error("FAILED")
	}

}

func Test_Sleep1_FAILURE_SCENARIO(t *testing.T){

	
	Before := time.Now().Second()
	fmt.Println("Before putting the golang program to sleep  .....nth Second at this instance is ",Before)

	Sleep_Vinit(5)
	After := time.Now().Second()
	fmt.Println("After program wakes up  .....nth Second at this instance ",After)

	if(After - Before) !=6{
		t.Error("FAILED")
	}

}