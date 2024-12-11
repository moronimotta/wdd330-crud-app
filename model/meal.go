package model

type Meal struct {
	ID       string  `json:"id"`
	Time     string  `json:"time"`
	Name     string  `json:"name"`
	Proteins float64 `json:"proteins"`
	Carbs    float64 `json:"carbs"`
	Fats     float64 `json:"fats"`
}

type MealPlan struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Monday    []Meal `json:"monday"`
	Tuesday   []Meal `json:"tuesday"`
	Wednesday []Meal `json:"wednesday"`
	Thursday  []Meal `json:"thursday"`
	Friday    []Meal `json:"friday"`
	Saturday  []Meal `json:"saturday"`
	Sunday    []Meal `json:"sunday"`
}
