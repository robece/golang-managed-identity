package main

import (
	"fmt"
	"order/controllers/secret"
	"order/structs"

	"github.com/kataras/iris/v12"
)

const (
	serviceName = "order"
	servicePort = "80"
	reasonKey   = "reason"
)

func init() {
	fmt.Println("package: order.main - initialized")
}

func order() *iris.Application {

	app := iris.New()
	app.OnAnyErrorCode(handleErrors)

	v1 := app.Party("/api/v1")
	{
		topicsAPI := v1.Party("/order")
		{
			topicsAPI.Post("/secret", secret.PostAuthHandler)
		}
	}

	app.Get("/", websiteHandler)

	return app
}

func main() {
	app := order()
	app.Run(iris.Addr(fmt.Sprint(":", servicePort)), iris.WithoutServerError(iris.ErrServerClosed))
}

func handleErrors(ctx iris.Context) {
	err := structs.HTTPError{
		Code:   ctx.GetStatusCode(),
		Reason: ctx.Values().GetStringDefault(reasonKey, "unknown"),
	}

	ctx.JSON(err)
}

func websiteHandler(ctx iris.Context) {
	ctx.ContentType("text/html")
	ctx.Writef(`
	<html>
	<head><title>Service is up and running!</title><head>
	<body>Service is up and running!</body>
	</html>`)
}
