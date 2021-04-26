package client

// Call represents an active RPC
type Call struct {
	Seq           uint64
	ServiceMethod string      // format "<service>.<method>"
	Args          interface{} // arguments to the function
	Reply         interface{} // reply from the function
	Error         error       // if error occurs, it will be complete
	Done          chan *Call  // Stores when call is complete
}

func (call *Call) done() {
	call.Done <- call
}

