package controllers

type Context interface {
	BindJSON(obj interface{}) error
	JSON(code int, obj any)
	Param(key string) string
	Query(key string) string
}
