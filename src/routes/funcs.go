package routes

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)


func Create(c *fiber.Ctx) error {
	/*
		Read request JSON body
	*/

	CreateReqBody := new(CreateReqBody)

	if err := c.BodyParser(CreateReqBody); err != nil {
		return err
	}

	/*
		Create new task
	*/

	newTask := new(Task)
	newTask.Id = uuid.New()
	newTask.Status = "not complete"
	newTask.TaskName = CreateReqBody.Message

	/*
		Read database JSON
	*/
	
	rawFile, err := ioutil.ReadFile("database.json")

	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, "Error occured while trying to access db!")
	}

	var data []Task

	_ = json.Unmarshal([]byte(rawFile), &data)

	/*
		Write database JSON
	*/

	updatedDatabaseRaw := append(data, *newTask)
	updatedDatabase, _ := json.MarshalIndent(updatedDatabaseRaw, "", " ")

	_ = ioutil.WriteFile("database.json", updatedDatabase, 0064)

	return c.JSON(updatedDatabaseRaw)
}

func Read(c *fiber.Ctx) error {
	rawFile, err := ioutil.ReadFile("database.json")

	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, "Error occured while trying to access db!")
	}

	var data []Task
	
	_ = json.Unmarshal([]byte(rawFile), &data)
	
	return c.JSON(data)
}

func Update(c  *fiber.Ctx) error {
	UpdateReqBody := new(UpdateReqBody)

	if err := c.BodyParser(UpdateReqBody); err != nil {
		return err
	}

	rawFile, err := ioutil.ReadFile("database.json")

	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, "Error occured while trying to access db!")
	}

	var data []Task

	_ = json.Unmarshal([]byte(rawFile), &data)

	for i := range(data){
		if data[i].Id == UpdateReqBody.Id {
			data[i].Status = UpdateReqBody.Status
		}
	}

	updatedDatabase, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("database.json", updatedDatabase, 0064)

	return c.JSON(data)
}

func Delete(c  *fiber.Ctx) error {
	DeleteReqBody := new(DeleteReqBody)

	if err := c.BodyParser(DeleteReqBody); err != nil {
		return err
	}

	rawFile, err := ioutil.ReadFile("database.json")

	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, "Error occured while trying to access db!")
	}

	var data []Task

	_ = json.Unmarshal([]byte(rawFile), &data)

	var newData []Task 

	for i := range(data){
		if data[i].Id == DeleteReqBody.Id {
			continue
		}
		newData = append(newData, data[i])
	}

	updatedDatabase, _ := json.MarshalIndent(newData, "", " ")

	if len(newData) == 0{
		updatedDatabase, _ = json.MarshalIndent([]int{}, "", " ")
	}

	_ = ioutil.WriteFile("database.json", updatedDatabase, 0064)
	return c.JSON(newData)
}
