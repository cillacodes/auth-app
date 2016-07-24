package controllers

import "github.com/revel/revel"

type Api struct {
	*revel.Controller
}

type ApiError struct {
	Type    string ` json:"type" xml:"type" `
	Message string ` json:"message" xml:"message" `
}

type ApiSuccess struct {
	Value string ` json:"value" xml:"value" `
}

type ApiResponse struct {
	Error ApiError   ` json:"error" xml:"error" `
	Data  ApiSuccess ` json:"data" xml:"data" `
}

func (c Api) SignIn() revel.Result {
	response := make(map[string]interface{})
	response["error"] = nil
	response["data"] = ApiSuccess{Value: "Sign in successful"}
	return c.RenderJson(response)
}

func (c Api) SignOut() revel.Result {
	response := make(map[string]interface{})
	response["error"] = nil
	response["data"] = ApiSuccess{Value: "Sign out successful"}
	return c.RenderJson(response)
}

func (c Api) SignUp() revel.Result {
	response := make(map[string]interface{})
	response["error"] = nil
	response["data"] = ApiSuccess{Value: "Sign up successful"}
	return c.RenderJson(response)
}
