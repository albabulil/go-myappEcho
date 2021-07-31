package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// User data structure
type User struct {
	ID *primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	Name            string     `bson:"name,omitempty" json:"name"`
	Email           string     `bson:"email,omitempty" json:"email" validate:"email"`
	IsEmailVerified bool       `bson:"emailVerification,omitempty" json:"emailVerification"`
	Phone           string     `bson:"phoneNumber,omitempty" json:"phone"`
	IsPhoneVerified bool       `bson:"phoneNumberVerification,omitempty" json:"phoneNumberVerification"`
	Password        string     `bson:"password,omitempty" json:"-"`
	Pwd             string     `bson:"-" json:"password,omitempty"`
	Picture         string     `bson:"picture,omitempty" json:"picture"`
	Roles           []UserRole `bson:"roles,omitempty" json:"-"`

	// legacy datas,
	// needed for backward compatibility
	Role          []string `bson:"role,omitempty" json:"-"`
	ApartmentID   string   `bson:"apartmentID,omitempty" json:"-"`
	ApartmentName string   `bson:"apartmentName,omitempty" json:"-"`
	Path          string   `bson:"path,omitempty" json:"-"`

	CreatedAt *time.Time `bson:"createdAt,omitempty" json:"-"`
	UpdatedAt *time.Time `bson:"updatedAt,omitempty" json:"-"`
	DeletedAt *time.Time `bson:"deletedAt" json:"-"`
}
type RequestChangePassword struct {
	UserID      string `json:"userID"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
type RequestResetPassword struct {
	Email       string `json:"email"`
	NewPassword string `json:"newPassword"`
}

// UserRole data structure
type UserRole struct {
	EntityID string `bson:"entityID,omitempty" json:"entityID,omitempty"`
	RoleID   string `bson:"roleID,omitempty" json:"roleID,omitempty"`
}

// Validate user
func (u User) Validate() error {
	if u.Email == "" && u.Phone == "" {
		return fmt.Errorf("User must have email or phone")
	}

	validate := validator.New()
	return validate.Struct(u)
}

//Validate request change password
func (r RequestChangePassword) Validate() error {
	if r.UserID == "" ||
		r.NewPassword == "" ||
		r.OldPassword == "" {
		return errors.New("UserID, old password, new password cannot be empty")
	}
	return nil
}

//Validate request Reset password
func (r RequestResetPassword) Validate() error {
	if r.Email == "" ||
		r.NewPassword == "" {
		return errors.New("email and new password cannot be empty")
	}
	return nil
}

// SetPassword set password
func (u *User) SetPassword() error {
	if u.Pwd == "" {
		return fmt.Errorf("Password is empty")
	}

	if len(u.Pwd) < 6 {
		return fmt.Errorf("Password must have at least 6 characters")
	}
	fmt.Printf("user: %+v\n", u)
	fmt.Println("pwd:", u.Pwd)
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Pwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("error generate hash:", err)
		return err
	}

	u.Password = string(hashed)
	u.Pwd = ""

	return nil
}
