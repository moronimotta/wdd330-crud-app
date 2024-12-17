package model

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`

	Height float64 `json:"height"`
	Weight float64 `json:"weight"`
	Age    int     `json:"age"`
	Gender string  `json:"gender"`
	Goal   string  `json:"goal"`

	ActivityFactor string `json:"activity_factor"`

	GoalMacroProteins float64 `json:"goal_macro_proteins"`
	GoalMacroCarbs    float64 `json:"goal_macro_carbs"`
	GoalMacroFats     float64 `json:"goal_macro_fats"`

	Notes string `json:"notes"`
}
