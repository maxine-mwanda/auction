package models

import (
	"API Auction/entities"
	"API Auction/utils"
	"fmt"
)


func Create_user(user entities.NewUser) (id int64, err error) {
	query := "insert into user (user_name, email, password, user_type)" +
		"values (?,?,?,?)"
	db, err := utils.Connecttodb()
	if err != nil {
		fmt.Println("unable to connect todb")
		return
	}
	row, err := db.Exec(query, user.UserName, user.Email, user.Password, user.UserType)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, _ = row.LastInsertId()
	return

}


func FetchUser(user entities.NewUser) () {
	db, err := utils.Connecttodb()
	if err != nil {
		fmt.Println("unable to connect todb")
		return
	}
	query := "select * from user where user_name = ?;"
	row := db.QueryRow(query, user)
	if row == nil {
		fmt.Println("unable to fetch list of users")
		return
	}

	err = row.Scan(&user.UserName, &user.Email, &user.Password, &user.UserType, &user.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func Create_product(product entities.NewProduct) (id int64, err error) {
	query := "insert into product (type, availability)" +
		"values (?,?)"
	db, err := utils.Connecttodb()
	if err != nil {
		fmt.Println("unable to connect todb")
		return
	}
	row, err := db.Exec(query, product.Type,product.Availability)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, _ = row.LastInsertId()
	return

}



func FetchProduct(product entities.NewProduct) {
	db, err := utils.Connecttodb()
	if err != nil {
		fmt.Println("unable to connect todb")
		return
	}
	query := "select * from product where type = ?;"
	row := db.QueryRow(query, product)
	if row == nil {
		fmt.Println("unable to fetch list of available products")
		return
	}

	err = row.Scan(&product.type, &product.Availability)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}