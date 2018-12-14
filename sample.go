// +build omit

package main

import "ermao"

func main() {
	app := ermao.New()

	app.Get("/xx", ermao.Handler(func(ctx ermao.Ctx) {
		ctx.Out("nihao ")
		ctx.Next()
		ctx.Out("tony")
	}), ermao.Handler(func(ctx ermao.Ctx) {
		ctx.Out("middle")
	}))

	app.Listen(":8080")

}
