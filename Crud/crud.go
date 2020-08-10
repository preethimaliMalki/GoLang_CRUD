package Crud

import (
	"log"
	"net/http"
	"strconv"
	  
	"github.com/preethimaliMalki/crud/DBConnection"
)

type Customer struct {
	Id      int
	Name    string
	Address string
}

type CustomerData struct {
	Customers []Customer
}

var html = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM customer ")
	if err != nil {
		panic(err.Error())
	}

	cus := Customer{}

	res := []Customer{}

	for selDB.Next() {
		var id int
		var name, address string
		err = selDB.Scan(&id, &name, &address)
		if err != nil {
			panic(err.Error())
		}
		cus.Id = id
		cus.Name = name
		cus.Address = address
		res = append(res, cus)
	}
	dataCus := CustomerData{
		Customers: res,
	}

	html.ExecuteTemplate(w, "Index.html", dataCus)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM customer WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	cus := Customer{}

	for selDB.Next() {
		var id int
		var name, address string
		err = selDB.Scan(&id, &name, &address)
		if err != nil {
			panic(err.Error())
		}
		cus.Id = id
		cus.Name = name
		cus.Address = address
	}

	html.ExecuteTemplate(w, "Show.html", cus)
	defer db.Close()

}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM customer WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	cus := Customer{}
	for selDB.Next() {
		var id int
		var name, address string
		err = selDB.Scan(&id, &name, &address)
		if err != nil {
			panic(err.Error())
		}
		cus.Id = id
		cus.Name = name
		cus.Address = address
	}
	html.ExecuteTemplate(w, "Edit.html", cus)
	defer db.Close()
}

func Insert(ww http.ResponseWriter, rr *http.Request) {
	db := dbConn()
	if rr.Method == "POST" {

		id, err := strconv.Atoi(rr.FormValue("id"))
		if err != nil {
			panic(err.Error())
		}
		name := rr.FormValue("name")
		address := rr.FormValue("address")

		insert, error := db.Query("INSERT INTO customer VALUES(?,?,?)", id, name, address)
		if error != nil {
			panic(error.Error())
		}

		defer insert.Close()

	}
	defer db.Close()
	http.Redirect(ww, rr, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		address := r.FormValue("address")
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			panic(err.Error())
		}
		insForm, err := db.Prepare("UPDATE customer SET  Name=?, Address=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, address, id)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM customer WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(nId)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}
