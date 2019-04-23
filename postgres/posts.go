package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"training/graphql_test5/models"
)

func (d *Db) CreatePost(user string, title string, body string) (*models.Post, error) {
	post := models.Post{
        Title: title,
        Body: body,
	}

	sqlStr, err := d.Prepare("insert into posts(user_id, title, body) values($1, $2, $3) returning id")
	if err != nil {
		fmt.Println("CreatePost Preparation Error: ", err)
		return nil, err
	}

	if err := sqlStr.QueryRow(user, title, body).Scan(&post.ID); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Create Post No rows error: " , err)
			return nil, err
		}
		fmt.Println("Error scanning rows: " , err)
		return nil, err
	}
	return &post, err
}

func (d *Db) RemovePost(id string) (bool, error) {
	sqlStr, err := d.Prepare("delete from posts where id = $1")
	if err != nil {
		fmt.Println("RemovePost Preparation Error: ", err)
		log.Fatal(err)
	}

	_, err = sqlStr.Exec(id)
	if err != nil {
		fmt.Println("RemovePost Query Error: ", err)
		return false, err
	}
	return true, nil
}

func (d *Db) GetPostsByUserId(userId string) ([]models.Post, error) {
	var posts []models.Post

	sqlStr, err := d.Prepare("select title, body from posts where user_id = $1")
	if err != nil {
		fmt.Println("GetPostsByUserId Preparation Error: ", err)
		return nil, err
	}

	rows, err := sqlStr.Query(userId)
	if err != nil {
		fmt.Println("GetPostsByUserId query execution error: " , err)
		return nil, err
	}
	for rows.Next() {
		post := models.Post{}
		err = rows.Scan(&post.Title, &post.Body)
		if err != nil {
			fmt.Println("Error scanning rows: " , err)
			return nil, err
		}
        posts = append(posts, post)
	}
	return posts, err
}
