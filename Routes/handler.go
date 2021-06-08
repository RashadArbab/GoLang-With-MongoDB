package Routes

import (
	"context"
	"fmt"
	"log"

	"github.com/RashadArbab/goServer/Database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var ctx = context.TODO()

type Note struct {
	Title     string `json: title, bson: title`
	Body      string `json: body, bson: body`
	Completed bool   `json: completed, bson: completed`
}
type User struct {
	Email string `json: email, bson: email`
	Name  string `json: name, bson: name`
	Pass  string `json: pass, bson: pass`
	Notes []Note `json: notes, bson: notes`
}

func GetAll(c *fiber.Ctx) error {

	fmt.Println("Checkpoint 1")

	var users []Database.User
	cursor, errorindb := Database.DB.Find(ctx, bson.M{})
	if errorindb != nil {
		log.Fatal("there was an error")
	}

	fmt.Println("Checkpoint 2")

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/

	defer cursor.Close(context.TODO())
	fmt.Println("Checkpoint 3")

	for cursor.Next(context.TODO()) {
		fmt.Println(cursor.Current.Values())
		// create a value into which the single document can be decoded
		var user Database.User
		// & character returns the memory address of the following variable.
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal("this is the error")
			log.Fatal(err)
		}
		fmt.Println(cursor)

		// add item our array
		users = append(users, user)
	}

	fmt.Println("Checkpoint 4")

	return c.JSON(users)

}

func CreateUser(c *fiber.Ctx) error {

	var body User
	c.BodyParser(&body)
	fmt.Println(body)

	cursor, err := Database.DB.InsertOne(context.Background(), bson.M{
		"email": body.Email,
		"pass":  body.Pass,
		"Notes": body.Notes})
	if err != nil {
		return c.Status(500).SendString("Sorry we ran into an error")
	}

	return c.JSON(cursor)
}

func updateNote(c *fiber.Ctx) error {
	type request struct {
		Email string `json: email`
		Pass  string `json: pass`
		Title string `json: title`
	}

	var body request

	c.BodyParser(&body)

	query := bson.M{
		"email":       body.Email,
		"pass":        body.Pass,
		"Notes.title": "this is a note"}
	update := bson.M{
		"$set": bson.M{
			"Notes.$.completed": true}}

	cursor, err := Database.DB.UpdateOne(context.Background(), query, update)

	if err != nil {
		return c.SendStatus(500)
	} else {
		fmt.Println(cursor)
	}

	return c.Status(200).JSON(cursor)
}

func GetSingle(c *fiber.Ctx) error {
	var document User
	fmt.Println(c.Params("id"))
	filter := bson.M{"email": c.Params("id")}
	fmt.Print(filter)
	err := Database.DB.FindOne(context.Background(), filter).Decode(&document)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).JSON(document)
}
