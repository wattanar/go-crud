package main

import (

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	_ "github.com/go-sql-driver/mysql"
    "database/sql"
    "net/http"

)

type Users struct{
	Id int
	Name string
	Age int
}

func main() {

	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
	  Layout: "layout",
	}))

  	m.Get("/" , func(r render.Render) {

  		db, _ := sql.Open("mysql", "root:@/test?charset=utf8")

  		rows, _ := db.Query("SELECT * FROM users")

	  	usr := []Users{}

	    for rows.Next() {
	        
	        b := Users{}

	        err := rows.Scan(&b.Id, &b.Name, &b.Age)

	        if err != nil {
		        panic(err)
		    }

	        usr = append(usr,b)
	    }


  		r.HTML(200,"home",usr)
  	})

  	m.Get("/about" , func(r render.Render) {

  		r.HTML(200,"about",nil);
  	})

  	m.Get("/edit/:id" , func(r render.Render , params martini.Params){
  		user_id := params["id"]

  		db, _ := sql.Open("mysql", "root:@/test?charset=utf8")

  		rows , err := db.Query("SELECT * FROM users WHERE id = ?" , user_id)

  		if err != nil {
	        panic(err)
	    }

  		usr := []Users{}

	    for rows.Next() {
	        
	        b := Users{}

	        err := rows.Scan(&b.Id, &b.Name, &b.Age)

	        if err != nil {
		        panic(err)
		    }

	        usr = append(usr,b)


	    }


  		r.HTML(200,"edit",usr)

  	})

  	m.Post("/update" , func(ren render.Render , r *http.Request) {

  		id := r.FormValue("UserID")
  		name  := r.FormValue("UserName")
  		age := r.FormValue("UserAge")

  		db, _ := sql.Open("mysql", "root:@/test?charset=utf8")

  		stmt, _ := db.Prepare("update users set name = ? , age = ? where id = ?")

  		stmt.Exec(name , age , id)
 

  		ren.Redirect("/")
  	})

  	m.Get("/add" , func(ren render.Render) {

  		ren.HTML(200,"add",nil)
  	})

  	m.Post("/create" , func(ren render.Render , r *http.Request) {

  		name := r.FormValue("UserName")
  		age := r.FormValue("UserAge")

  		db, _ := sql.Open("mysql", "root:@/test?charset=utf8")

  		stmt, _ := db.Prepare("INSERT INTO users(name,age) VALUES(?,?)")

  		stmt.Exec(name,age)

  		ren.Redirect("/")

  	})

  	m.Get("/delete/:id" , func(ren render.Render , params martini.Params) {

  		user_id := params["id"]

  		db, _ := sql.Open("mysql", "root:@/test?charset=utf8")

  		_, err := db.Exec("DELETE FROM users WHERE id = ? " , user_id)  // OK
  		
  		if err != nil {
  			panic(err)
  		}

  		ren.Redirect("/")
  	})

	m.Run()
}