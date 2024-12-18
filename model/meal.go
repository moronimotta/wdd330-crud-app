package model

type MealEntry struct {
	Meal string `json:"meal"`
	Time string `json:"time"`
}

type MealPlan struct {
	ID        string      `json:"id,omitempty"`
	UserID    string      `json:"user_id"`
	Monday    []MealEntry `json:"monday"`
	Tuesday   []MealEntry `json:"tuesday"`
	Wednesday []MealEntry `json:"wednesday"`
	Thursday  []MealEntry `json:"thursday"`
	Friday    []MealEntry `json:"friday"`
	Saturday  []MealEntry `json:"saturday"`
	Sunday    []MealEntry `json:"sunday"`
}
