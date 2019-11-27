package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// реализация интерфейсов

// Vehicle interface
type Vehicle interface {
	move()
}

// Car type
type Car struct {
	model string
}

// Aircraft type
type Aircraft struct {
	model string
}

func (c Car) move() {
	fmt.Println(c.model, "It is car")
}

func (a Aircraft) move() {
	fmt.Println(a.model, "It is aircraft")
}

type Usr struct {
	id       int
	username string
	box string
	coach bool
	dark_theme bool
	email string
	firstname string
	gender string
	last_password_reset_date int
	lastname string
	password string
	registered int
	thumb_nail string
	user_pic string
}

func main() {
	//vaz := Car{model: "2107"}
	//tu := Aircraft{model: "140"}
	//vehicles := [...]Vehicle{vaz, tu}
	//for _, vehicle := range vehicles {
	//	vehicle.move()
	//}

	//for i := 1; i < 1000000; i++ {
	//	go func(n int) {
	//		fmt.Println(i)
	//	}(i)
	//}

	connStr := "user=user password=root dbname=coachu sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from usr")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	users := []Usr{}

	for rows.Next() {
		p := Usr{}
		err := rows.Scan(&p.id, &p.username, &p.box, &p.coach, &p.dark_theme, &p.email, &p.firstname, &p.gender,
			&p.last_password_reset_date, &p.last_password_reset_date, &p.password, &p.registered, &p.thumb_nail, &p.user_pic)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}
	for _, p := range users {
		fmt.Println(p.id, p.username)
	}

}
