package datasource

import "github.com/betterde/ects/internal/models"

var Users =  map[int64]models.User {
	1: {
		ID: 1,
		Name: "George",
		Email: "george@betterde.com",
		Password: "George@1994",
		Avatar: "",
		GroupId: 1,
		CreatedAt: "2019-01-01 23:00:00",
		UpdatedAt: "2019-01-01 23:00:00",
	},
}