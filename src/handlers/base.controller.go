package handlers

const (
	BaseRoute = "/"
)

type ErrorRes struct {
	Error    string `json:"error"'`
	Code     string `json:"code"'`
	ErrorDis string `json:"error_description"'`
}

type BasicResponse struct {
	Message string `json:"message"`
}

type ProductReq struct {
	Name  string `json:"name"`
	Price string `json:"price"`
	Code  string `json:"code"`
	Image string `json:"image"`
}
