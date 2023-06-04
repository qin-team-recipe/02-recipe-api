package controllers

type H struct {
	Message string
	Object  interface{}
}

func NewH(message string, obj interface{}) *H {
	h := new(H)

	h.Message = message
	h.Object = obj

	return h
}
