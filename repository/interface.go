package repository

import (
	"context"

	"github.com/moronimotta/wdd330-crud-app/model"
)

type UserRepository interface {
	GetUser(ctx context.Context, email string) (model.User, error)
	CreateUser(ctx context.Context, in model.User) (model.User, error)
	UpdateUser(ctx context.Context, in model.User) (model.User, error)
	DeleteUser(ctx context.Context, email string) error
	ListUsers(ctx context.Context) ([]model.User, error)
}

type MealPlanRepository interface {
	ListMealPlans(ctx context.Context) ([]model.MealPlan, error)
	GetMealPlan(ctx context.Context, id string) (model.MealPlan, error)
	CreateMealPlan(ctx context.Context, mealPlan model.MealPlan) (model.MealPlan, error)
	UpdateMealPlan(ctx context.Context, id string, updates model.MealPlan) (model.MealPlan, error)
	DeleteMealPlan(ctx context.Context, id string) error
}
