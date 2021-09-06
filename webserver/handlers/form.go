package handlers

import "../../class/form"
import "github.com/gin-gonic/gin"
import "../../types"
import "../answer"

func FormHandler(w *gin.Context) {
	i := types.Form{}
	err := w.ShouldBind(&i)
	if err != nil {
		answer.HandlerError(w, "7")
		return
	}
	form.FormSubmit(w, i)
}