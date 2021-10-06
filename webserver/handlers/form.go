package handlers

import "github.com/Psh777/visa-backend/class/form"
import "github.com/gin-gonic/gin"
import "github.com/Psh777/visa-backend/types"
import "github.com/Psh777/gin_handler_answer"

func FormHandler(w *gin.Context) {
	i := types.Form{}
	err := w.ShouldBind(&i)
	if err != nil {
		answer.HandlerError(w, "7")
		return
	}
	form.FormSubmit(w, i)
}
