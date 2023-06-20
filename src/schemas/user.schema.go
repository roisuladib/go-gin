package schemas

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ? Register struct
type Register struct {
	Name				 string    `json:"name" bson:"name" binding:"required"`
	Email				 string    `json:"email" bson:"email" binding:"required"`
	Password			 string    `json:"password" bson:"password" binding:"required,min=8"`
	PasswordConfirm string    `json:"passwordConfirm" bson:"passwordConfirm,omitempty" binding:"required"`
	Role				 string	  `json:"role" bson:"role"`
	Verified			 bool      `json:"verified" bson:"verified"`
	CreatedAt		 time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt		 time.Time `json:"updated_at" bson:"updated_at"`
}

// ? Login struct
type Login struct {
	Email				 string    `json:"email" bson:"email" binding:"required"`
	Password			 string    `json:"password" bson:"password" binding:"required"`
}

// ? DB struct
type DB struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	Name				 string    			  `json:"name" bson:"name"`
	Email				 string    			  `json:"email" bson:"email"`
	Password			 string    			  `json:"password" bson:"password"`
	PasswordConfirm string    			  `json:"passwordConfirm,omitempty" bson:"passwordConfirm,omitempty"`
	Role				 string	  			  `json:"role" bson:"role"`
	Verified			 bool      			  `json:"verified" bson:"verified"`
	CreatedAt		 time.Time 			  `json:"created_at" bson:"created_at"`
	UpdatedAt		 time.Time 			  `json:"updated_at" bson:"updated_at"`
}

// ? User struct
type User struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name				 string    			  `json:"name,omitempty" bson:"name,omitempty"`
	Email				 string    			  `json:"email,omitempty" bson:"email,omitempty"`
	Role				 string	  			  `json:"role,omitempty" bson:"role,omitempty"`
	Verified			 bool      			  `json:"verified,omitempty" bson:"verified,omitempty"`
	CreatedAt		 time.Time 			  `json:"created_at" bson:"created_at"`
	UpdatedAt		 time.Time 			  `json:"updated_at" bson:"updated_at"`
}

func Formatted(user *DB) User {
	return User{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Role: user.Role,
		Verified: user.Verified,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
