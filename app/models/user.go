package models

import (
	"fmt"
	"os"
	"strings"

	"github.com/twinj/uuid" //jwt-best-practices
	"golang.org/x/crypto/bcrypt"

	//"github.com/gofrs/uuid" //offersapp
	"time"

	"gorm.io/gorm"
)

var (
	tokenSecret = []byte(os.Getenv("TOKEN_SECRET"))
)

//https://gorm.io/docs/conventions.html
//type Tabler interface {
//TableName() string
//}

// TableName overrides the table name used by Empleado to `employee`
func (User) TableName() string {
	return "user_account"
}

// BeforeCreate will set a UUID rather than numeric ID. https://gorm.io/docs/create.html
func (tab *User) BeforeCreate(*gorm.DB) error {
	//uuidx := uuid.NewV4()
	tab.Id = uuid.NewV4().String()
	return nil
}

type User struct {
	Id string `gorm:"primary_key;column:id" json:"id"` //json:"id,omitempty"
	//ID              uuid.UUID `json:"id"`
	CreatedAt       time.Time `json:"_"`
	UpdatedAt       time.Time `json:"_"`
	Email           string    `gorm:"column:email" json:"email"`
	PasswordHash    string    `json:"-"`
	Password        string    `gorm:"-" json:"password"`
	PasswordConfirm string    `gorm:"-" json:"password_confirm"`
}

func (u *User) Register(conn *gorm.DB) error {

	if len(u.Password) < 4 || len(u.PasswordConfirm) < 4 {
		return fmt.Errorf("Password must be at least 4 characters long.")
	}

	if u.Password != u.PasswordConfirm {
		return fmt.Errorf("Passwords do not match.")
	}

	if len(u.Email) < 4 {
		return fmt.Errorf("Email must be at least 4 characters long.")
	}

	u.Email = strings.ToLower(u.Email)
	var userLookup User
	var err error
	if err = conn.First(&userLookup, "email = ?", u.Email).Error; err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		//return fmt.Errorf("Error p" + err.Error())
	}

	//row := conn.QueryRow(context.Background(), "SELECT id from user_account WHERE email = $1", u.Email)
	//userLookup := User{}
	//err := row.Scan(&userLookup)
	if u.Email == strings.ToLower(userLookup.Email) {
		fmt.Println("found user")
		fmt.Println(userLookup.Email)
		return fmt.Errorf("A user with that email already exists")
	}

	pwdHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("There was an error creating your account.")
	}
	u.PasswordHash = string(pwdHash)

	//now := time.Now()
	//_, err = conn.Exec(context.Background(), "INSERT INTO user_account (created_at, updated_at, email, password_hash) VALUES($1, $2, $3, $4)", now, now, u.Email, u.PasswordHash)

	//row2 := conn.QueryRow(context.Background(), "SELECT id, password_hash from user_account WHERE email = $1", u.Email)
	//row2.Scan(&u.ID, &u.PasswordHash)
	//if err := c.BindJSON(&userLookup); err != nil {

	//	return fmt.Errorf(err.Error())
	//}

	//u.Password = ""
	//u.PasswordConfirm = ""
	conn.Create(&u)
	return err // ya te asigna el ID del user
}

// IsAuthenticated checks to make sure password is correct and user is active
func (u *User) IsAuthenticated(conn *gorm.DB) error {
	//row := conn.QueryRow(context.Background(), "SELECT id, password_hash from user_account WHERE email = $1", u.Email)
	//err := row.Scan(&u.ID, &u.PasswordHash)
	//if err == pgx.ErrNoRows {
	//	fmt.Println("User with email not found")
	//	return fmt.Errorf("Invalid login credentials")
	//}

	u.Email = strings.ToLower(u.Email)
	var userLookup User
	var err error
	if err = conn.First(&userLookup, "email = ?", u.Email).Error; err != nil {
		//return fmt.Errorf("Error p" + err.Error())
		fmt.Println("User with email not found")
		return fmt.Errorf("Invalid login credentials email")
	}
	//if u.Email == strings.ToLower(userLookup.Email) {
	//	fmt.Println("found user")
	//	fmt.Println(userLookup.Email)
	//	return fmt.Errorf("A user with that email already exists")
	//}

	//pwdHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	//if err != nil {
	//	return fmt.Errorf("There was an error creating your account.")
	//}
	u.PasswordHash = string(userLookup.PasswordHash)
	u.Id = userLookup.Id

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(u.Password))
	if err != nil {
		//fmt.Println("pss: " + err.Error())
		return fmt.Errorf("Invalid login credentials pass")
	}

	return nil
}

func (u *User) UpdatePassword(conn *gorm.DB) error {

	if len(u.Password) < 4 || len(u.PasswordConfirm) < 4 {
		return fmt.Errorf("Password must be at least 4 characters long.")
	}

	if u.Password != u.PasswordConfirm {
		return fmt.Errorf("Passwords do not match.")
	}

	if len(u.Email) < 4 {
		return fmt.Errorf("Email must be at least 4 characters long.")
	}

	u.Email = strings.ToLower(u.Email)
	var userLookup User
	var err error
	if err = conn.First(&userLookup, "email = ?", u.Email).Error; err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		//return fmt.Errorf("Error p" + err.Error())
	}

	//row := conn.QueryRow(context.Background(), "SELECT id from user_account WHERE email = $1", u.Email)
	//userLookup := User{}
	//err := row.Scan(&userLookup)
	if u.Email != strings.ToLower(userLookup.Email) {
		fmt.Println("found user")
		fmt.Println(userLookup.Email)
		return fmt.Errorf("A user with that email not exists")
	}

	pwdHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("There was an error creating your account.")
	}
	u.PasswordHash = string(pwdHash)

	//now := time.Now()
	//_, err = conn.Exec(context.Background(), "INSERT INTO user_account (created_at, updated_at, email, password_hash) VALUES($1, $2, $3, $4)", now, now, u.Email, u.PasswordHash)

	//row2 := conn.QueryRow(context.Background(), "SELECT id, password_hash from user_account WHERE email = $1", u.Email)
	//row2.Scan(&u.ID, &u.PasswordHash)
	//if err := c.BindJSON(&userLookup); err != nil {

	//	return fmt.Errorf(err.Error())
	//}

	//u.Password = ""
	//u.PasswordConfirm = ""
	conn.Save(&u)
	return err // ya te asigna el ID del user
}
