package promise

import (
	"fmt"
	"sync"
)

type promiseImpl struct {
	val          interface{};
	err          error;
	valueChannel chan interface{};
	errorChannel chan  error;
	catchWG      sync.WaitGroup;
	finallyWG    sync.WaitGroup;
}
func NewPromise(f func() (interface{},error)) (*promiseImpl){
	p := &promiseImpl{
		valueChannel: make(chan interface{},1),
		errorChannel: make(chan error,1),
	}
	p.catchWG.Add(1);
	p.finallyWG.Add(1);
	go func() {
		p.val,p.err = f();
		if(p.err != nil){
			p.errorChannel <- p.err;
			fmt.Println("Error sent");
		}else {
			p.valueChannel <- p.val;
			fmt.Println("Value sent");
		}
		p.catchWG.Done();
		p.finallyWG.Done();
	}();
	return p;
}

func (p *promiseImpl )Then(onFulfilled func(v interface{}))  Promise{
	go func() {
		select {
		case val := <- p.valueChannel:
			//fmt.Println("[Then] Calling onFulfilled with val ",val);
			onFulfilled(val);
		case  <- p.errorChannel:
			//fmt.Printf("\n [Then] Then block Error occured",err);
			fmt.Println();
		}
	}();
	return p;
}


func (p *promiseImpl )ThenWithErrorHandler(onFulfilled func(v interface{}), onRejected func(e error)) Promise {
	go func() {
		select {
		case val := <- p.valueChannel:
			//fmt.Println("[ThenWithErrorHandler] Calling onFulfilled with val ",val);
			onFulfilled(val);
		case err := <- p.errorChannel:
			//fmt.Printf("[ThenWithErrorHandler] ThenWithErrorHandler calling onRejected with error ",err);
			fmt.Println();
			onRejected(err);
		}
	}();
	return p;
}

func (p *promiseImpl )Catch(onRejected func(e error))  Promise{
	go func() {
		p.catchWG.Wait();
		if(p.err != nil){
			//fmt.Printf("[Catch] Catch on Rejected called occured ",p.err);
			fmt.Println();
			onRejected(p.err);
		}
	}();
	return p;
}
func (p *promiseImpl )Finally(onFinally func())  Promise{
	go func() {
		p.finallyWG.Wait();
		onFinally();
	}();
	return p;
}
