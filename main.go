package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

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

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func initDatabase() *gorm.DB {
	db, err := gorm.Open("postgres", "user=postgres password=kush dbname=gorm sslmode=disable")
	CheckError(err)
	return db
}

func main() {
	//Initalization of GIN
	r := gin.Default()
	// Connect to Database
	db := initDatabase()
	//Close the Database after main is over
	defer db.Close()

	r.GET("/read_all_database", func(c *gin.Context) {

		// sqlStatement := `SELECT * FROM public.app_processes`
		// rows, err := db.DB().Query(sqlStatement)
		// CheckError(err)
		// defer rows.Close()
		// for rows.Next() {
		// 	err := rows.Scan(&AppID, &Status, &User, &UserID)
		// 	CheckError(err)
		// 	fmt.Println(AppID, UserID, Status, User)
		// }

		var data []AppProcess
		db.Find(&data)
		//  iterative onto the database and get all the fields or jobs having timespan less than 2 hours
		for _, u := range data {
			fmt.Println("User_Name", u.User, "App_ID", u.AppID, "Status", u.Status, "User_ID", u.UserID)
			time.Sleep(1 * time.Second)
		}

		var user []UserDetails
		db.Find(&user)
		//  iterative onto the database and get all the fields or jobs having timespan less than 2 hours
		for _, u := range user {
			fmt.Println("User_Name", u.User, "User_ID", u.UserID, "Email", u.EmailID, "Session", u.Session)
			time.Sleep(1 * time.Second)
		}

		var payment []Billing
		db.Find(&payment)
		//  iterative onto the database and get all the fields or jobs having timespan less than 2 hours
		for _, u := range payment {
			fmt.Println("User_ID", u.UserID, "Payment_ID", u.PaymentID, "Payment_Amount", u.PaymentAmount, "Payment_Details", u.PaymentDetails)
			time.Sleep(1 * time.Second)
		}

	})

	r.GET("/:id", func(c *gin.Context) {
		var data []AppProcess
		datasetID, _ := strconv.Atoi(c.Param("id"))
		db.Where("user_id =?", uint(datasetID)).Find(&data)
		fmt.Println(reflect.TypeOf(&data))
		for _, u := range data {
			fmt.Println("App_Id:", u.AppID, "Created_Time:", u.CreatedAt, "App_Status:", u.Status, "User_Name:", u.User)
		}
	})

	r.Run()
}
