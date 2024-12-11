package model

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`

	Height float64 `json:"height"`
	Weight float64 `json:"weight"`
	Age    int     `json:"age"`
	Gender string  `json:"gender"`
}
