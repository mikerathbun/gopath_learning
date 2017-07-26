package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

const costCategoryId = 1
const cost = 2.50
const totalSize = 1000
const tierBreakdown = 1000

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
func main() {
	xmlFile, err := os.Open("data/idrive/nec/usage.xml")
	if err != nil {
		fmt.Println("Error opening up file:", err)
		return
	}
	defer xmlFile.Close()
	b, _ := ioutil.ReadAll(xmlFile)
	var q Query

	xml.Unmarshal(b, &q)
	for _, user := range q.UserList {
		fmt.Printf("\t%s\n", user)
	}
}
