package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/moronimotta/wdd330-crud-app/model"
	"github.com/moronimotta/wdd330-crud-app/repository"
)

type Server struct {
	userRepo repository.UserRepository
	mealRepo repository.MealPlanRepository
}

func NewServer(userRepo repository.UserRepository, mealRepo repository.MealPlanRepository) *Server {
	return &Server{userRepo: userRepo,
		mealRepo: mealRepo}
}

func (s Server) GetUser(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument email"})
		return
	}
	user, err := s.userRepo.GetUser(ctx, email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (s Server) ListUsers(ctx *gin.Context) {
	users, err := s.userRepo.ListUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (s Server) CreateUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	user, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (s Server) UpdateUser(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument email"})
		return
	}
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	user.Email = email
	user, err := s.userRepo.UpdateUser(ctx, user)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (s Server) DeleteUser(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument email"})
		return
	}
	if err := s.userRepo.DeleteUser(ctx, email); err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

// Meal Functions
func (s Server) CreateMeal(ctx *gin.Context) {
	var meal model.MealPlan
	if err := ctx.ShouldBindJSON(&meal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	meal, err := s.mealRepo.CreateMealPlan(ctx, meal)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"meal": meal})
}

func (s Server) GetMeal(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument id"})
		return
	}
	meal, err := s.mealRepo.GetMealPlan(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrMealPlanNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"meal": meal})
}

func (s Server) ListMeals(ctx *gin.Context) {
	meals, err := s.mealRepo.ListMealPlans(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"meals": meals})
}
func (s Server) UpdateMeal(ctx *gin.Context) {
	var meal model.MealPlan
	if err := ctx.ShouldBindJSON(&meal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument id"})
		return
	}

	updatedMeal, err := s.mealRepo.UpdateMealPlan(ctx, id, meal)
	if err != nil {
		if errors.Is(err, repository.ErrMealPlanNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"meal": updatedMeal})
}

func (s Server) DeleteMeal(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument id"})
		return
	}
	if err := s.mealRepo.DeleteMealPlan(ctx, id); err != nil {
		if errors.Is(err, repository.ErrMealPlanNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}
