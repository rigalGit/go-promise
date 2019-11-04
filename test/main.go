package main

import (
	"errors"
	"promise-t/promise"
	"time"
)
import "fmt"

func main() {
	fmt.Printf("start \n");
	fmt.Printf("------------------------Testing for success return value----------------------- \n");
	testForSuccessReturnValue();
	time.Sleep(20 * time.Second);
	fmt.Printf("------------------------Testing for Error return value----------------------- \n");
	testForErrorReturnValue()
	time.Sleep(60 * time.Second);
}

func testForSuccessReturnValue() {
	f := func() (interface{}, error) {
		fmt.Println("Executing success return function ");
		time.Sleep(5 * time.Second);
		fmt.Printf("Exeuction done for success case");
		return 100, nil;
	}
	p1 := promise.NewPromise(f)
	p1.Then(func(data interface{}) {
		value := data.(int);
		fmt.Println("[Then-1] Got return value ", value);
		fmt.Println("[Then-1] Processing ", value);
		time.Sleep(5 * time.Second);
		fmt.Println("[Then-1] Processed done", value);
	}).Finally(func() {
		fmt.Println("[Then-Finally-1] finally called for promise p1 done");
	})
}

func testForErrorReturnValue() {
	errorFunc := func() (interface{}, error) {
		fmt.Printf("Executing Async Task  \n");
		time.Sleep(5 * time.Second);
		fmt.Printf("Task Done , will throw error \n");
		return nil, errors.New("Custom Error");
	}
	p2 := promise.NewPromise(errorFunc)
	onFulfilled := func(r interface{}) {
		fmt.Printf("[onFulfilled]  Got this from promise %d\n , now sleeping \n", r);
		time.Sleep(5 * time.Second);
		fmt.Printf("[onFulfilled] done \n");
	}
	onRejected := func(e error) {
		fmt.Println("[onRejected] Handling error ", e);
		time.Sleep(2 * time.Second);
		fmt.Println("[onRejected] Error handling done ");
	}
	p2.ThenWithErrorHandler(onFulfilled, onRejected).Catch(func(e error) {
		fmt.Println("[onRejected-ThenWithErrorHandler] Error in catch block2 done", e);
	}).Finally(func() {
		fmt.Println("[onRejected-ThenWithErrorHandler] Finally called2 done");
	})
}
