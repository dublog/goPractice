package Echo

import(
	"context"
    "github.com/mfslog/goPractice/RPCX_Message/demo"
)

type Calc int


func(c *Calc) Add(ctx context.Context,request demo.CalcRequest, response demo.CalcResponse) error {
	response.C = request.A + request.B
	return nil
}