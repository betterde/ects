package organization

import (
	"github.com/kataras/iris"
	"log"
)

type TeamController struct {

}

func (instance * TeamController) Get(ctx iris.Context)  {
	if _, err := ctx.Text("This is Team"); err != nil {
		log.Println(err)
	}
}