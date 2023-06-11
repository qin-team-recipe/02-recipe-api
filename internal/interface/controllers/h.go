package controllers

type H struct {
	Message string      `json:"message"`
	Object  interface{} `json:"data"`
}

func NewH(message string, obj interface{}) *H {
	h := new(H)

	h.Message = message
	h.Object = obj

	return h
}
