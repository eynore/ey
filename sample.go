// +build omit

package main

import ey "github.com/eynore"

func main() {
	app := ey.New()

	app.Get("/xx", func(ctx ey.Ctx) {
		ctx.Out("nihao ")
		ctx.Next()
		ctx.Out("tony")
	}, func(ctx ey.Ctx) {
		ctx.Out("middle")
	})

	app.Listen(":8080")

}
