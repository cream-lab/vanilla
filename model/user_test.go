package model

import (
	"fmt"
	"time"

	"github.com/cream-lab/vanilla/model"
)

func ExampleSetUser() {
	user := model.User{
		Seq:           1,
		Email:         "luv4u525@gmail.com",
		Name:          "Teddy",
		ProfileImgUrl: "http://graph.facebook.com/686990364693854/picture?type=square",
		Bio:           "I'm a developer",
		IdProvider:    model.Facebook,
		RegisterDate:  time.Now(),
		UpdateDate:    time.Now(),
	}

	fmt.Println(user.IdProvider)

	// Output: Facebook
}
