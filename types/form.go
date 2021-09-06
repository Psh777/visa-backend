package types

type Form struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Message    string `json:"message"`
	Captcha bool   `json:"captcha"`
	Code    string `json:"code"`
}
