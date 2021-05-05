package models

// import ("fmt")

type ResponseError struct{
	Code int `json:"-"`
	Description string `json:"description"`
}
// func (e *ResponseError) Error() string{
// 	return fmt.Sprintf("%s", e.Description)
// }