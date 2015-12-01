package taobao

import (
	"testing"
	"fmt"
	"github.com/smartwalle/taobao/params"
)

func TestOpenSmsSendMsgParam(t *testing.T) {
	var p = params.OpenSmsSendMsgParam{}
	p.Mobile = "13558858128"
	p.TemplateId = "1025"
	p.AddParam("code", "1234511116")
	fmt.Println(Send("23274732", "8eed038bab11210aa8099d8e450cd631", p))
}