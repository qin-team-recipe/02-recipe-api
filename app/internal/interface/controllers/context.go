package controllers

type Context interface {
	BindJSON(obj interface{}) error
	JSON(code int, obj any)
	GetHeader(key string) string
	MustGet(key string) any
	Param(key string) string
	PostForm(key string) (value string)
	Query(key string) string
	Value(key any) any
}
