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

func (r userRepository) ListUsers(ctx context.Context) ([]model.User, error) {
	cursor, err := r.db.
		Collection("users").
		Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var out []user
	if err := cursor.All(ctx, &out); err != nil {
		return nil, err
	}

	users := make([]model.User, len(out))
	for i, u := range out {
		users[i] = toModel(u)
	}
	return users, nil
}

func (r userRepository) GetUser(ctx context.Context, email string) (model.User, error) {
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

func (r userRepository) DeleteUser(ctx context.Context, email string) error {
	out, err := r.db.
		Collection("users").
		DeleteOne(ctx, bson.M{"email": email})
	if err != nil {
		return err
	}
	if out.DeletedCount == 0 {
		return ErrUserNotFound
	}
	return nil
}

type user struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`

	Height float64 `bson:"height,omitempty"`
	Weight float64 `bson:"weight,omitempty"`
	Age    int     `bson:"age,omitempty"`
	Gender string  `bson:"gender,omitempty"`
}

func fromModel(in model.User) user {
	return user{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,

		Height: in.Height,
		Weight: in.Weight,
		Age:    in.Age,
		Gender: in.Gender,
	}
}

func toModel(in user) model.User {
	return model.User{
		ID:       in.ID.String(),
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
		Height:   in.Height,
		Weight:   in.Weight,
		Age:      in.Age,
		Gender:   in.Gender,
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

func (r mealPlanRepository) ListMealPlans(ctx context.Context) ([]model.MealPlan, error) {
	cursor, err := r.db.Collection("mealplans").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var out []mealPlan
	if err := cursor.All(ctx, &out); err != nil {
		return nil, err
	}

	mealPlans := make([]model.MealPlan, len(out))
	for i, mp := range out {
		mealPlans[i] = toMealPlanModel(mp)
	}
	return mealPlans, nil
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

	updateDoc := bson.M{}
	if updates.UserID != "" {
		updateDoc["user_id"] = updates.UserID
	}
	if updates.Monday != nil {
		updateDoc["monday"] = updates.Monday
	}
	if updates.Tuesday != nil {
		updateDoc["tuesday"] = updates.Tuesday
	}
	if updates.Wednesday != nil {
		updateDoc["wednesday"] = updates.Wednesday
	}
	if updates.Thursday != nil {
		updateDoc["thursday"] = updates.Thursday
	}
	if updates.Friday != nil {
		updateDoc["friday"] = updates.Friday
	}
	if updates.Saturday != nil {
		updateDoc["saturday"] = updates.Saturday
	}
	if updates.Sunday != nil {
		updateDoc["sunday"] = updates.Sunday
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

func (r mealPlanRepository) DeleteMealPlan(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	deleteResult, err := r.db.Collection("mealplans").DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	if deleteResult.DeletedCount == 0 {
		return ErrMealPlanNotFound
	}

	return nil
}

// --- Helper Functions ---

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
