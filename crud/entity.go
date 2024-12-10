package wdd330crudapp

import "time"

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type Meal struct {
	Time     string  `bson:"time"`     // Meal time, e.g., "08:00 AM"
	Name     string  `bson:"name"`     // Meal name, e.g., "Breakfast"
	Proteins float64 `bson:"proteins"` // Protein content
	Carbs    float64 `bson:"carbs"`    // Carb content
	Fats     float64 `bson:"fats"`     // Fat content
}

type DayPlan map[int]Meal // Map of meals for the day (1, 2, 3...)

type MealPlan struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"` // MongoDB ID
	UserID    string             `bson:"user_id"`       // Reference to the user
	Monday    DayPlan            `bson:"monday"`        // Meals for Monday
	Tuesday   DayPlan            `bson:"tuesday"`       // Meals for Tuesday
	Wednesday DayPlan            `bson:"wednesday"`     // Meals for Wednesday
	Thursday  DayPlan            `bson:"thursday"`      // Meals for Thursday
	Friday    DayPlan            `bson:"friday"`        // Meals for Friday
	Saturday  DayPlan            `bson:"saturday"`      // Meals for Saturday
	Sunday    DayPlan            `bson:"sunday"`        // Meals for Sunday
	CreatedAt time.Time          `bson:"created_at"`    // Timestamp for creation
	UpdatedAt time.Time          `bson:"updated_at"`    // Timestamp for last update
}
