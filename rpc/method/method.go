package method

import "errors"

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

// 1. 方法必需满足func (t *T) MethodName(argType T1, replyType *T2) error 签名
// 2. 方法必需是导出的
// 3. 方法接受两个参数
// 4. 方法必需返回error类型 .
type Arith struct{}

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by 0")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
