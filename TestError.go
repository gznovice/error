package main

import(
	"fmt"
	"errors"
)

type MyError struct{
	errNo int
	errMsg string
}

func (myerr *MyError)Error() string{
	return fmt.Sprintf("%d --- %s", myerr.errNo, myerr.errMsg)
}

//func MyfuncWithErr(i int) (int, *MyError){
func MyfuncWithErr(i int) (int, error){
	if i >= 100 {
		return 0, &MyError{101, "Input can't be more than 100"}
	}else 	if i > 50 {
		return 0, errors.New("Input bigger than 50 is Not allowed Now")
	}
	return i*2, nil
}

func main(){
	var input = []int{3, 2, 50, 80, 120}
	
	for _, v := range(input){
		ret, err := MyfuncWithErr(v)
		if(err == nil){
			fmt.Println("We got the value", ret)
		}else{
			fmt.Println("We see the error", err.Error())
		}
	}
}