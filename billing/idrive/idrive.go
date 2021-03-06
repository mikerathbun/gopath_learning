package idrive

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	costCategoryId           = 1
	cost                     = 5.00
	totalSize                = 1000
	tierBreakdown            = 1000
	customerName             = "NEC"
	invoiceChargeDescription = "Provide workgroup shares/home directory file storage(700.1.7.1)"
	expenseDate              = "8/8/2017"
	chargeType               = "Gigabyte(s)"
)

type Query struct {
	UserList []User `xml:"User"`
}
type User struct {
	Name string `xml:"Name,attr"`
	Size int    `xml:"Usage-MB"`
}

func (u User) String() string {
	return fmt.Sprintf("The name is: %s - With a size of: %d and cost of %.2f", u.Name, u.Size, u.Cost())
}
func (u User) Cost() float64 {
	var userCost = 0.0

	if u.Size-totalSize <= tierBreakdown {
		return userCost
	}
	var myTier = int(u.Size / tierBreakdown)
	return float64(myTier-1) * cost

}

func (u User) ChargeLine() string {
	return fmt.Sprintf("%s I drive size %d", u.Name, u.Size)

}

type ChargeLine struct {
	ChargeName        string
	ChargeDescription string
	ChargeDate        time.Time
	ChargeAmount      float64
}

func (u ChargeLine) String() string {
	return fmt.Sprintf("%s", u.ChargeDescription)

}
func (u ChargeLine) Cost() string {
	return fmt.Sprintf("$%.2f", u.ChargeAmount)
}
func GetCharges() []ChargeLine {
	var returnCharges []ChargeLine

	chargeDate := time.Date(2017, time.July, 28, 0, 0, 0, 0, time.UTC)
	db, err := sql.Open("mysql",
		"root:www.mike.com@tcp(127.0.0.1:3306)/billing?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT ChargeName, 
				ChargeDescription, ChargeDate, ChargeAmount 
				FROM charges 
				WHERE ChargeAmount > 0 
				AND CostCategoryID = ? 
				AND CustomerName = ? 
				AND ChargeDate = ?`, 1, "502_FSS", chargeDate)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		u := new(ChargeLine)

		err := rows.Scan(&u.ChargeName, &u.ChargeDescription, &u.ChargeDate, &u.ChargeAmount)
		if err != nil {
			log.Fatal(err)
		}
		returnCharges = append(returnCharges, *u)
		// fmt.Println(u)

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return returnCharges
}
func RunCharges() {
	chargeDate := time.Date(2017, time.July, 28, 0, 0, 0, 0, time.UTC)
	db, err := sql.Open("mysql",
		"root:www.mike.com@tcp(127.0.0.1:3306)/billing")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	xmlFile, err := os.Open("data/idrive/502_FSS/usage.xml")
	if err != nil {
		fmt.Println("Error opening up file:", err)
		return
	}
	defer xmlFile.Close()
	b, _ := ioutil.ReadAll(xmlFile)
	var q Query

	xml.Unmarshal(b, &q)
	stmt, err := db.Prepare("INSERT INTO charges(CostCategoryID, ChargeName, ChargeDescription, ChargeDate, ChargeAmount, CustomerName) VALUES(?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range q.UserList {
		_, err := stmt.Exec(costCategoryId, "Personal Storage", user.ChargeLine(), chargeDate, user.Cost(), customerName)
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Printf("\t%s\n", user)
	}
}
