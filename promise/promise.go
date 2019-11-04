package promise

type Promise interface {
	Then( onFulfilled func(v interface{})) Promise;
	ThenWithErrorHandler( onFulfilled func(v interface{}),onRejected func(e error)) Promise;
	Catch(onRejected func(e error)) Promise;
	Finally(onFinally func()) Promise;
}
