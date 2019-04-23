// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type Comment struct {
	ID    *string `json:"id"`
	User  User    `json:"user"`
	Post  Post    `json:"post"`
	Title *string `json:"title"`
	Body  string  `json:"body"`
}

type Post struct {
	ID       *string   `json:"id"`
	User     User      `json:"user"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	Comment  *Comment  `json:"comment"`
	Comments []Comment `json:"comments"`
}

type User struct {
	ID        *string `json:"id"`
	Email     string  `json:"email"`
	Post      *Post   `json:"post"`
	Posts     []Post  `json:"posts"`
	Follower  *User   `json:"follower"`
	Followers []User  `json:"followers"`
	Folowee   *User   `json:"folowee"`
	Followees []User  `json:"followees"`
}
