package utils

type RPCError struct {
	state   int
	message string
	err     error
}

func (rpc RPCError) Error() string {
	if rpc.err != nil {
		return rpc.err.Error()
	}
	return ""

}

func (rpc RPCError) GetState() int {
	return rpc.state
}

func (rpc RPCError) GetMessage() string {
	return rpc.message
}

func NewRPCError(err error, state int) RPCError {
	return RPCError{
		state: state,
		err:   err,
	}
}

func NewSuccess(message string) RPCError {
	return RPCError{
		state:   1,
		message: message,
	}
}
