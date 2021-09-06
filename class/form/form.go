package form

import (
	"../../modules/telegram"
	"../../types"
	"../../webserver/answer"
	"fmt"
	"github.com/gin-gonic/gin"
)

func FormSubmit(w *gin.Context, form types.Form) {
	fmt.Printf("%+v", form)

	line := "\n\n"

	message := "**New message from Contact form:**" + line + form.Name + "\n" + form.Phone + "\n" + form.Email + "\n\n" + form.Message + line

	telegram.SendMsgBot(message)
	answer.HandlerInterface(w, "ok")
}