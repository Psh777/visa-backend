package answer

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Headers(c *gin.Context) *gin.Context {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "content-type,X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin, Authorization, Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	return c
}

func HeaderAccess(w *gin.Context) *gin.Context {
	w.Writer.Header().Set("Content-Type", "application/pdf")
	w.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	w.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	return w
}

func HandlerError(w *gin.Context, codeError string) {
	w = Headers(w)
	//textError := lang.CodeError(codeError)
	out := Error{false, codeError, codeError, nil}
	w.AsciiJSON(http.StatusBadRequest, out)
}

func HandlerSuccess(w *gin.Context, successText string) {
	w = Headers(w)
	out := Success{true, successText}
	w.AsciiJSON(http.StatusOK, out)
}

func HandlerInterface(w *gin.Context, data interface{}) {
	w = Headers(w)
	out := Success{
		Success: true,
		Result:  data,
	}
	w.AsciiJSON(http.StatusOK, out)
}

func HandlerPrint(w *gin.Context, data string) {
	w = Headers(w)
	w.Writer.Header().Set("Content-Type", "Content-Type: text/html; charset=utf-8")
	w.String(http.StatusOK, data)
}

func HandlerSuccessFile(w *gin.Context, file string) {
	w = HeaderAccess(w)
	w.File(file)
}

