package test

import (
	"testing"

	"github.com/yaaaaashiki/cstack/domain/repository"
	"github.com/yaaaaashiki/cstack/helper"
)

const Rand = 100

//FindByEmailOrNil function
//Success test case
//We give email address in DB, then this function should return correct user name
func TestReturnCorrectUserName(t *testing.T) {
	db := startUpDatabase()
	defer db.Close()
	userRepository := repository.NewUserRepository(db)

	var cases = []struct {
		mail string
		name string
	}{
		{"user1@example.com", "user1"},
		{"user2@example.com", "user2"},
		{"user3@example.com", "user3"},
		{"user4@example.com", "user4"},
		{"user5@example.com", "user5"},
		{"user8@example.com", "user8"},
	}
	for _, c := range cases {
		res, err := userRepository.FindByEmailOrNil(c.mail)
		checkErr(t, err)
		if res.Name != c.name {
			t.Errorf("When mail = %v, should return %v for the user name, but actually this function return %v", c.mail, c.name, res.Name)
		}
	}
}

//FindByEmailOrNil function
//Success test case
//We give missing email address, then this function should return nil
func TestReturnNilForUser(t *testing.T) {
	db := startUpDatabase()
	defer db.Close()
	userRepository := repository.NewUserRepository(db)

	var cases = []struct {
		mail string
	}{
		{"user100@example.com"},
		{"user500@example.com"},
		{"user800@example.com"},
	}
	for _, c := range cases {
		res, err := userRepository.FindByEmailOrNil(c.mail)
		if err != nil {
			t.Errorf("This function should return nil, but actually return the err (%v)", err)
		}
		if res != nil {
			t.Errorf("This function should return nil, but actually return the something (%v)", res)
		}
	}
}

//RegisterUser function
//Success test case
//We give some arguments, then should be created user data
//If you look at the gorm's panic error, please run 'make testenv'
func TestRegisterUser(t *testing.T) {
	db := startUpDatabase()
	defer db.Close()
	userRepository := repository.NewUserRepository(db)

	var cases = []struct {
		name     string
		email    string
		password string
		image    string
	}{
		{"user9", "user9@example.com", "password", "test"},
		{"user10", "user10@example.com", "hoge", "test"},
	}
	for _, c := range cases {
		salt := helper.Salt(Rand)
		salted := helper.Stretch(c.password, salt)
		act, err := userRepository.RegisterUser(c.name, c.email, salt, salted, c.image)
		if err != nil {
			t.Errorf(err.Error())
		}
		if act.Name != c.name && act.Email != c.email && act.Salt != salt && act.Salted != salted && act.IconImage != c.image {
			t.Errorf("Cannot create correct user data")

			if act.Name != c.name {
				t.Errorf("input name is %v, but actual return data is %v", c.name, act.Name)
			}
			if act.Email != c.email {
				t.Errorf("input email is %v, but actual return data is %v", c.email, act.Email)
			}
			if act.Salt != salt {
				t.Errorf("salt is %v, but actual return data is %v", salt, act.Salt)
			}
			if act.Salted != salted {
				t.Errorf("salted is %v, but actual return data is %v", salted, act.Salted)
			}
			if act.IconImage != c.image {
				t.Errorf("input image is %v, but actual return data is %v", c.image, act.IconImage)
			}
		}
	}
}

//RegisterUser function
//Failed test case
//We give some arguments, then should return err (unique)
//If you look at the gorm's panic error, please run 'make testenv'
/*
func TestReturnErrorForUser(t *testing.T) {
	db := startUpDatabase()
	defer db.Close()
	userRepository := repository.NewUserRepository(db)

	var cases = []struct {
		fbID     uint
		name     string
		email    string
		image    string
		biograpy string
		sex      int
	}{
		{10000, "a", "b", "c", "d", 1},
		{10000, "a", "b", "c", "d", 1},
	}
	for i, c := range cases {
		_, err := userRepository.RegisterUser(c.fbID, c.name, c.email, c.image, c.biograpy, c.sex)

		if i == 2 && err == nil {
			t.Errorf("This function Should return err (duplicate datas), but this time return nil")
		}
	}
}
*/
