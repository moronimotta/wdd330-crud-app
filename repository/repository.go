package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/moronimotta/wdd330-crud-app/model"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type userRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{db: db}
}

func (r userRepository) GetUser(ctx context.Context, email, password string) (model.User, error) {
	var out user
	err := r.db.
		Collection("users").
		FindOne(ctx, bson.M{"email": email, "password": password}).
		Decode(&out)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.User{}, ErrUserNotFound
		}
		return model.User{}, err
	}
	return toModel(out), nil
}

func (r userRepository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	var out user
	err := r.db.
		Collection("users").
		FindOne(ctx, bson.M{"email": email}).
		Decode(&out)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.User{}, ErrUserNotFound
		}
		return model.User{}, err
	}
	return toModel(out), nil
}

func (r userRepository) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	out, err := r.db.
		Collection("users").
		InsertOne(ctx, fromModel(user))
	if err != nil {
		return model.User{}, err
	}
	user.ID = out.InsertedID.(primitive.ObjectID).String()
	return user, nil
}

func (r userRepository) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	in := bson.M{}
	if user.Name != "" {
		in["name"] = user.Name
	}
	if user.LastName != "" {
		in["last_name"] = user.LastName
	}
	if user.Password != "" {
		in["password"] = user.Password
	}

	if user.Height != 0 {
		in["height"] = user.Height
	}
	if user.Weight != 0 {
		in["weight"] = user.Weight
	}
	if user.Age != 0 {
		in["age"] = user.Age
	}
	if user.Gender != "" {
		in["gender"] = user.Gender
	}

	if user.Goal != "" {
		in["goal"] = user.Goal
	}
	if user.GoalMacroProteins != 0 {
		in["goal_macro_proteins"] = user.GoalMacroProteins
	}

	if user.GoalMacroCarbs != 0 {
		in["goal_macro_carbs"] = user.GoalMacroCarbs
	}
	if user.GoalMacroFats != 0 {
		in["goal_macro_fats"] = user.GoalMacroFats
	}
	if user.Notes != "" {
		in["notes"] = user.Notes
	}

	out, err := r.db.
		Collection("users").
		UpdateOne(ctx, bson.M{"email": user.Email}, bson.M{"$set": in})
	if err != nil {
		return model.User{}, err
	}
	if out.MatchedCount == 0 {
		return model.User{}, ErrUserNotFound
	}
	return user, nil
}

type user struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	LastName string             `bson:"last_name,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`

	Height float64 `bson:"height,omitempty"`
	Weight float64 `bson:"weight,omitempty"`
	Age    int     `bson:"age,omitempty"`
	Gender string  `bson:"gender,omitempty"`

	Goal string `bson:"goal,omitempty"`

	GoalMacroProteins float64 `bson:"goal_macro_proteins,omitempty"`
	GoalMacroCarbs    float64 `bson:"goal_macro_carbs,omitempty"`
	GoalMacroFats     float64 `bson:"goal_macro_fats,omitempty"`

	Notes string `bson:"notes,omitempty"`
}

func fromModel(in model.User) user {
	return user{
		Name:     in.Name,
		LastName: in.LastName,
		Email:    in.Email,
		Password: in.Password,

		Height: in.Height,
		Weight: in.Weight,
		Age:    in.Age,
		Gender: in.Gender,
		Goal:   in.Goal,

		GoalMacroProteins: in.GoalMacroProteins,
		GoalMacroCarbs:    in.GoalMacroCarbs,
		GoalMacroFats:     in.GoalMacroFats,

		Notes: in.Notes,
	}
}

func toModel(in user) model.User {
	return model.User{
		ID:                in.ID.String(),
		Name:              in.Name,
		Email:             in.Email,
		Password:          in.Password,
		Height:            in.Height,
		Weight:            in.Weight,
		Age:               in.Age,
		Gender:            in.Gender,
		Goal:              in.Goal,
		GoalMacroProteins: in.GoalMacroProteins,
		GoalMacroCarbs:    in.GoalMacroCarbs,
		GoalMacroFats:     in.GoalMacroFats,
		Notes:             in.Notes,
	}
}

var (
	ErrMealPlanNotFound = errors.New("meal plan not found")
)

type mealPlanRepository struct {
	db *mongo.Database
}

// NewMealPlanRepository creates a new instance of MealPlanRepository
func NewMealPlanRepository(db *mongo.Database) MealPlanRepository {
	return &mealPlanRepository{db: db}
}

func (r mealPlanRepository) GetMealPlanByUserID(ctx context.Context, userID string) (model.MealPlan, error) {
	var out mealPlan
	err := r.db.Collection("mealplans").FindOne(ctx, bson.M{"user_id": userID}).Decode(&out)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.MealPlan{}, ErrMealPlanNotFound
		}
		return model.MealPlan{}, err
	}
	return toMealPlanModel(out), nil
}

func (r mealPlanRepository) GetMealPlan(ctx context.Context, id string) (model.MealPlan, error) {
	var out mealPlan
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.MealPlan{}, err
	}

	err = r.db.Collection("mealplans").FindOne(ctx, bson.M{"_id": objectID}).Decode(&out)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.MealPlan{}, ErrMealPlanNotFound
		}
		return model.MealPlan{}, err
	}
	return toMealPlanModel(out), nil
}

func (r mealPlanRepository) CreateMealPlan(ctx context.Context, mealPlan model.MealPlan) (model.MealPlan, error) {
	insertResult, err := r.db.Collection("mealplans").InsertOne(ctx, fromMealPlanModel(mealPlan))
	if err != nil {
		return model.MealPlan{}, err
	}
	mealPlan.ID = insertResult.InsertedID.(primitive.ObjectID).Hex()
	return mealPlan, nil
}

func (r mealPlanRepository) UpdateMealPlan(ctx context.Context, id string, updates model.MealPlan) (model.MealPlan, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.MealPlan{}, err
	}

	updateDoc := bson.M{
		"user_id":   updates.UserID,
		"monday":    updates.Monday,
		"tuesday":   updates.Tuesday,
		"wednesday": updates.Wednesday,
		"thursday":  updates.Thursday,
		"friday":    updates.Friday,
		"saturday":  updates.Saturday,
		"sunday":    updates.Sunday,
	}

	updateResult, err := r.db.Collection("mealplans").UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": updateDoc})
	if err != nil {
		return model.MealPlan{}, err
	}
	if updateResult.MatchedCount == 0 {
		return model.MealPlan{}, ErrMealPlanNotFound
	}

	return r.GetMealPlan(ctx, id)
}

type mealPlan struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    string             `bson:"user_id"`
	Monday    []model.Meal       `bson:"monday,omitempty"`
	Tuesday   []model.Meal       `bson:"tuesday,omitempty"`
	Wednesday []model.Meal       `bson:"wednesday,omitempty"`
	Thursday  []model.Meal       `bson:"thursday,omitempty"`
	Friday    []model.Meal       `bson:"friday,omitempty"`
	Saturday  []model.Meal       `bson:"saturday,omitempty"`
	Sunday    []model.Meal       `bson:"sunday,omitempty"`
}

func fromMealPlanModel(in model.MealPlan) mealPlan {
	return mealPlan{
		UserID:    in.UserID,
		Monday:    in.Monday,
		Tuesday:   in.Tuesday,
		Wednesday: in.Wednesday,
		Thursday:  in.Thursday,
		Friday:    in.Friday,
		Saturday:  in.Saturday,
		Sunday:    in.Sunday,
	}
}

func toMealPlanModel(in mealPlan) model.MealPlan {
	return model.MealPlan{
		ID:        in.ID.Hex(),
		UserID:    in.UserID,
		Monday:    in.Monday,
		Tuesday:   in.Tuesday,
		Wednesday: in.Wednesday,
		Thursday:  in.Thursday,
		Friday:    in.Friday,
		Saturday:  in.Saturday,
		Sunday:    in.Sunday,
	}
}
