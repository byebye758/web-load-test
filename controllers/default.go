package controllers

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego"
	"io"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl"
	go Xunhuan()
	c.Ctx.WriteString("nihao")

}

func Xunhuan() {
	for i := 0; i < 100000; i++ {
		fmt.Println(Hash("nihao"))

	}
}
func Hash(data string) string {

	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))

}
