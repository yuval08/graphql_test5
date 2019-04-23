package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"training/graphql_test5/models"
)

func (d *Db) GetUserById(id string) (*models.User, error) {
    user := models.User{ID: &id}

	sqlStr, err := d.Prepare("select email from users where id = $1")
	if err != nil {
		fmt.Println("GetUserById Preparation Error: ", err)
		return nil, err
	}

	if err := sqlStr.QueryRow(id).Scan(&user.Email); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("GetUserById No rows error: " , err)
			err = fmt.Errorf("no data found for id %s", id)
			return nil, err
		}
		fmt.Println("Error scanning rows: " , err)
		return nil, err
	}

	posts, err := d.GetPostsByUserId(id)
	if err != nil {
		fmt.Println("Error loading posts: " , err)
		return nil, err
	}
	user.Posts = posts

	return &user, err
}

func (d *Db) CreateUser(email string) (*models.User, error) {
	user := models.User{
		Email: email,
	}

	sqlStr, err := d.Prepare("insert into users(email) values($1) returning id")
	if err != nil {
		fmt.Println("CreateUser Preparation Error: ", err)
		return nil, err
	}

	rows, err := sqlStr.Query(user.Email)
	if err != nil {
		fmt.Println("CreateUser Query Error: ", err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&user.ID)
		if err != nil {
			fmt.Println("Error scanning rows: " , err)
			return nil, err
		}
	}
	return &user, err
}

func (d *Db) RemoveUser(id string) (bool, error) {
	sqlStr, err := d.Prepare("delete from users where id = $1")
	if err != nil {
		fmt.Println("RemoveUser Preparation Error: ", err)
		log.Fatal(err)
	}

	_, err = sqlStr.Exec(id)
	if err != nil {
		fmt.Println("RemoveUser Query Error: ", err)
        return false, err
	}

	return true, nil
}

