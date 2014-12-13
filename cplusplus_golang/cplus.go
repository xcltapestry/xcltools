//author:xiong chanliang
//date: 2014-12-13
package main

import (
	"errors"
	"fmt"
)

func main() {
	
path := `c:\a\b\c\GG再也不用烦转义符了`
mulln := `"(C++/Golang
	 aa'aaa\Cplusplus/gogogo
	 author"xiongchuanliang
	`	
fmt.Println(path)
fmt.Println(mulln)

	init_demo();
	lambda_demo();
	iota_demo();

	ti,ts := tuple_demo();
	fmt.Printf("ti:%d ts:%s \n",ti,ts)

	mapfind_demo();

	
	//可变参数
	args_demo(5,6,7,8);
	mArr := []int{ 5, 6, 7,8}
	args_demo(mArr[0],mArr[1:]...);

	fmt.Println("fmtPrintln(): ", 1, 2.0, "C++11", "Golang");


	callFunc("回调就是你调我，我调它，大家一起玩。",printFunc)

    v1 := 13
	v2 := 53
	ret, err := compare(v1, v2)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch ret {
	case -1:
		fmt.Println("v1 < v2")
	case 0:
		fmt.Println("v1 == v2")
	case 1:
		fmt.Println("v1 > v2")
	default:
		fmt.Println("defualt")
	}

	slice_demo();
}

func init_demo(){
	var k int
	fmt.Println(k)

	mArr := []int{ 1, 2, 3}
	var mMap = map[string]int {"a":1,"b":2}
	fmt.Printf("array:%v\nmap:%v \n",mArr,mMap)

	i := 10  
	pi := &i  //*i
	ppi := &pi  //**int
	fmt.Println(i,*pi,**ppi)
}

func lambda_demo(){

	mArr := []int{ 1, 2, 3}
	fun := func(i,v int){
		fmt.Printf("idx:%d value:%d \n",i,v)
	}
	for idx,val := range mArr{
		fun(idx,val)
	}
}

func iota_demo(){
	const (
		Red = iota
		Green
		Blue
	)
	fmt.Println("Red:",Red," Gree:",Green," Blue:",Blue);
}

func tuple_demo()(int,string){

	a,b := 1,2
	fmt.Println("a:",a," b:",b);

	c,d := b,a
	fmt.Println("c:",c," d:",d);

	return 168, "函数返回的字符串"
}

func mapfind_demo(){
	var mMap = map[string]int {"a":1,"b":2,"c":3}
	val,found := mMap["b"]
	if found {
		fmt.Println("found :",val);
	}else{
		fmt.Println("not found");
	}
}

func args_demo(first int,args ...int){
	fmt.Println("ars_demo first:",first)
	for _,i := range args {
		fmt.Println("args:",i)
	}
}

func compare(v1, v2 interface{}) (int, error) {

	switch v1.(type) {
	case int:
		if v1.(int) < v2.(int) {
			return -1, nil
		} else if v1.(int) == v2.(int) {
			return 0, nil
		}
	case int8:
		if v1.(int8) < v2.(int8) {
			return -1, nil
		} else if v1.(int8) == v2.(int8) {
			return 0, nil
		}
	case int32:
		if v1.(int32) < v2.(int32) {
			return -1, nil
		} else if v1.(int32) == v2.(int32) {
			return 0, nil
		}
	case int64:
		if v1.(int64) < v2.(int64) {
			return -1, nil
		} else if v1.(int64) == v2.(int64) {
			return 0, nil
		}
	case float32:
		if v1.(float32) < v2.(float32) {
			return -1, nil
		} else if v1.(float32) == v2.(float32) {
			return 0, nil
		}
	case float64:
		if v1.(float64) < v2.(float64) {
			return -1, nil
		} else if v1.(float64) == v2.(float64) {
			return 0, nil
		}
	default:
		return -2, errors.New("未能处理的数据类型.")
	}
	return 1, nil
}

type funcType func(string)

func printFunc(str string){
	fmt.Println( "callFunc() -> printFunc():",str)
}

func callFunc(arg string,f funcType){
	f(arg)
}

//callFunc("回调就是你调我，我调它，大家一起玩。",printFunc)

func slice_demo(){
    a := [5]int{ 1, 2, 3, 4, 5 }
    b := a[:3]
    c := a[1:2]
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
}