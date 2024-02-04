package services

import (
	"water/config"
	"water/models"

	_ "gorm.io/driver/postgres"
)

type Response struct {
	Status   int    `json:"status"`
	Message  string `json:"message"`
	Response string `json:"response"`
}

func ConvertToResponse(status int, message string, response string) Response {
	var quality models.Quality
	if err := config.DB.Where("name = ?", response).First(&quality).Error; err != nil {
		return Response{Status: 404, Message: "Not Found", Response: ""}
	}

	return Response{Status: status, Message: message, Response: response}

}

func GetAllQuality(quality *[]models.Quality) (err error) {
	if err = config.DB.Find(quality).Error; err != nil {
		return err
	}
	return nil
}
