package datasource

import "github.com/betterde/ects/internal/models"

var Users =  map[string]models.User {
	"9527": {
		ID: "9527",
		Name: "George",
		Email: "george@betterde.com",
		Password: "George@1994",
		Avatar: "",
		GroupId: 1,
		CreatedAt: "2019-01-01 23:00:00",
		UpdatedAt: "2019-01-01 23:00:00",
	},
}