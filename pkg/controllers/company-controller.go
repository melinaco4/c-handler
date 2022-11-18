package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Company struct {
	ID                string `json:"id,omitempty" bson:"_id,omitempty"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	AmountOfEmployees int    `json:"amount_of_employees"`
	Registered        bool   `json: "registered"`
	Type              string `json: "type"`
}

func GetCompanies(c *fiber.Ctx) error {

	query := bson.D{{}}

	cursor, err := mg.Db.Collection("companies").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var companies []Company = make([]Company, 0)

	if err := cursor.All(c.Context(), &companies); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(companies)

}

func CreateCompany(c *fiber.Ctx) error {
	collection := mg.Db.Collection("employees")

	company := new(Company)

	if err := c.BodyParser(company); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	company.ID = ""

	insertionResult, err := collection.InsertOne(c.Context(), company)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	createdCompany := &Company{}
	createdRecord.Decode(createdCompany)

	return c.Status(201).JSON(createdCompany)

}

func UpdateCompany(c *fiber.Ctx) error {
	idParam := c.Params("id")

	companyID, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		return c.SendStatus(400)
	}

	company := new(Company)

	if err := c.BodyParser(company); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	query := bson.D{{Key: "_id", Value: companyID}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "name", Value: company.Name},
				{Key: "description", Value: company.Description},
				{Key: "amount_of_employees", Value: company.AmountOfEmployees},
				{Key: "registered", Value: company.Registered},
				{Key: "type", Value: company.Type},
			},
		},
	}

	mg.Db.Collection("companies").FindOneAndUpdate(c.Context(), query, update).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(400)
		}

		return c.SendStatus(500)
	}

	company.ID = idParam

	return c.Status(200).JSON(company)
}

func DeleteCompany(c *fiber.Ctx) error {

	companyID, err := primitive.ObjectIDFromHex(c.Params("id"))

	if err != nil {
		return c.SendStatus(400)
	}

	query := bson.D{{Key: "_id", Value: companyID}}
	result, err := mg.Db.Collection("companies").DeleteOne(c.Context(), &query)

	if err != nil {
		return c.SendStatus(500)
	}

	if result.DeletedCount < 1 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON("company deleted")

}
