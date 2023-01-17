package main

func main() {
	//hertz := server.New()
	//hertz.POST("/login", func(c context.Context, ctx *app.RequestContext) {
	//	//logs.Infof("",ctx.Request.)
	//	log.Infof("Recevied RawRequest：%s", ctx.Request.RawRequest())
	//})

}

//
//func (ctx *RequestContext) Next() {
//	ctx.index++
//	for ctx.index < int8(len(ctx.handlers)) {
//		ctx.handlers[ctx.index]()
//		ctx.index++
//	}
//}
//func (ctx *RequestContext) Abort() {
//	ctx.index = IndexMax
//}

//type Conn interface {
//	net.Conn
//	Reader
//	Writer
//}
func index(b []byte, c byte) int {
	for i := 0; i < len(b); i++ {
		if b[i] == c {
			return i
		}
	}
	return -1
}
