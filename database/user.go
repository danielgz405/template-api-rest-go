package database

import (
	"context"
	"fmt"

	"github.com/danielgz405/template-api-rest-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *MongoRepo) InsertUser(ctx context.Context, user *models.InsertUser) (profile *models.Profile, err error) {
	collection := repo.client.Database("[db-name]").Collection("users")
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, fmt.Errorf("email already exists: %v", err)
		}
		return nil, err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	profile, err = repo.GetUserById(ctx, oid.Hex())
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (repo *MongoRepo) GetUserById(ctx context.Context, id string) (*models.Profile, error) {
	collection := repo.client.Database("[db-name]").Collection("users")
	var user models.User
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	// Find one and populate company
	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		return nil, err
	}
	// Populate profile
	var profile = models.Profile{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Roles: user.Roles,
	}
	return &profile, nil
}

func (repo *MongoRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	collection := repo.client.Database("[db-name]").Collection("users")
	var user models.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *MongoRepo) ListUsers(ctx context.Context) ([]models.Profile, error) {
	collection := repo.client.Database("StockWiseWorks").Collection("users")
	cursor, err := collection.Find(ctx, bson.M{}, options.Find().SetSort(bson.M{"_id": -1}))
	if err != nil {
		return nil, err
	}
	var users []models.User
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}
	profiles := []models.Profile{}
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		// Populate profile
		var profile = models.Profile{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
			Roles: user.Roles,
		}
		profiles = append(profiles, profile)
	}
	return profiles, nil
}

func (repo *MongoRepo) UpdateUser(ctx context.Context, data models.UpdateUser) (*models.Profile, error) {
	collection := repo.client.Database("[db-name]").Collection("users")
	oid, err := primitive.ObjectIDFromHex(data.Id)
	if err != nil {
		return nil, err
	}
	update := bson.M{
		"$set": bson.M{},
	}
	iterableData := map[string]interface{}{
		"name":  data.Name,
		"email": data.Email,
		"roles": data.Roles,
	}
	for key, value := range iterableData {
		if value != "" || value == nil {
			update["$set"].(bson.M)[key] = value
		}
	}
	err = collection.FindOneAndUpdate(ctx, bson.M{"_id": oid}, update).Err()
	if err != nil {
		return nil, err
	}
	profile, err := repo.GetUserById(ctx, data.Id)
	if err != nil {
		return nil, err
	}
	return profile, nil
}
func (repo *MongoRepo) DeleteUser(ctx context.Context, id string) error {
	collection := repo.client.Database("[db-name]").Collection("users")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}

func (repo *MongoRepo) UpdateUserPassword(ctx context.Context, userId string, newPassword string) (profile *models.Profile, err error) {
	collection := repo.client.Database("[db-name]").Collection("users")

	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": oid}
	update := bson.M{"$set": bson.M{"password": newPassword}}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	profile, err = repo.GetUserById(ctx, oid.Hex())
	if err != nil {
		return nil, err
	}
	return profile, nil
}
