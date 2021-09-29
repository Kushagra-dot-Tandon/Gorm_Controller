package main

import (
	// "fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

type AppProcess struct {
	gorm.Model

	AppID  int    //['1', '2' , '3']
	Status string //['Failed' , 'Done' , 'Pending']
	User   string //Kushagra
	UserID int    //12
}

type UserDetails struct {
	gorm.Model

	User    string //first_name
	UserID  int    //12
	EmailID string //@gmail.com
	Session string //Active_Diasbled
}

type Billing struct {
	gorm.Model

	UserID         int    //UserId
	PaymentID      int64  //paymentID generated during pay
	PaymentAmount  int    //amount
	PaymentDetails string //['credit' , 'upi' , 'debit']
}

func main() {
	db, err := gorm.Open("postgres", "user=postgres password=kush dbname=gorm sslmode=disable")
	CheckError(err)

	// close the databse after the main function finishes
	defer db.Close()
	// Connect the database , initatiate
	database := db.DB()
	// Check if the connections is made or not
	err = database.Ping()
	CheckError(err)

	//Inserting the data onto the database
	db.AutoMigrate(&AppProcess{})
	// Filling the data onto the database
	var person1 = &AppProcess{AppID: 2, Status: "pending", User: "kushagra", UserID: 1}
	var person2 = &AppProcess{AppID: 1, Status: "done", User: "sushant", UserID: 2}
	var person3 = &AppProcess{AppID: 3, Status: "failed ", User: "thouqueer", UserID: 3}
	db.Create(person1)
	db.Create(person2)
	db.Create(person3)

	db.AutoMigrate(&UserDetails{})
	var userdetails1 = &UserDetails{User: "Kushagra", UserID: 1, EmailID: "kushagra.tandon@maplelabs.com", Session: "Active"}
	var userdetails2 = &UserDetails{User: "Sushant", UserID: 2, EmailID: "sushant.pandey@maplelabs.com", Session: "Offline"}
	var userdetails3 = &UserDetails{User: "thouqueer", UserID: 2, EmailID: "thouqueer.ahmed@maplelabs.com", Session: "Active"}

	db.Create(userdetails1)
	db.Create(userdetails2)
	db.Create(userdetails3)

	db.AutoMigrate(&Billing{})
	var billingDetails1 = &Billing{UserID: 1, PaymentID: 32092883, PaymentAmount: 2425, PaymentDetails: "Debit"}
	var billingDetails2 = &Billing{UserID: 2, PaymentID: 32052682, PaymentAmount: 25553, PaymentDetails: "UPI"}
	var billingDetails3 = &Billing{UserID: 3, PaymentID: 26095535, PaymentAmount: 5545, PaymentDetails: "Credit"}

	db.Create(billingDetails1)
	db.Create(billingDetails2)
	db.Create(billingDetails3)
	// updating the data onto the database
}
