package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID 				primitive.ObjectID 	`json:"id" bson:"_id,omitempty"`
	Username 		string			   	`json:"username" bson:"username,omitempty"`
	Email			string				`json:"email" bson:"email,omitempty"`
	Password		string				`json:"password" bson:"password,omitempty"`
	IsEmailVerified	bool				`json:"is_email_verified" bson:"is_email_verified,omitempty"`
	SocialProfile	string				`json:"social_profile" bson:"social_profile,omitempty"`
}

type UserLogin struct {
	Email		string `json:"email" bson:"email,omitempty"`
	Password	string `json:"password" bson:"password,omitempty"`
}

type UserPasswordChange struct {
	OldPassword	string				`json:"old_password"`
	NewPassword	string				`json:"new_password"`
}