package form

import (
	"fmt"
	"github.com/Psh777/gin_handler_answer"
	"github.com/Psh777/visa-backend/modules/telegram"
	"github.com/Psh777/visa-backend/types"
	"github.com/gin-gonic/gin"
)

func FormSubmit(w *gin.Context, form types.Form) {
	fmt.Printf("%+v", form)

	line := "\n\n"

	message := "**New message from Contact form:**" + line + form.Name + "\n" + form.Phone + "\n" + form.Email + "\n\n" + form.Message + line

	telegram.SendMsgBot(message)
	answer.HandlerInterface(w, "ok")
}
