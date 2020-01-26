package main

import "time"
import "log"
import "toDoProgram/internal/db"
//import cmd "toDoProgram/pkg/query"

/*
So, Golang depends on drivers to access the SQL. Seems like I need the under-
lying MySQL to do things. Which makes sense.

Following this tutorial:
http://go-database-sql.org/index.html

And using best practices with a config file for SQL access.

This is definitely a prototype and will definitely need to be refactored into
something more robust in terms of code organization. Having a massive main file
is not very useful
*/

func main(){

  conn := db.GetLogin()

  var (
    date time.Time
    entry string
  )

  //A Prepared Statement. Good for security and reusability, but if the statement
  //isn't used again, maaaay not be worth the overhead
  //In this case, it's probably not worth it
  stmt, err := conn.Prepare("select date, entry from to_do")
  if err != nil {
	  log.Fatal(err)
  }
  defer stmt.Close()

  rows, err := stmt.Query()

  if err != nil {
	  log.Fatal(err)
  }
  //Remember to always close your connections, as a defer, OUTSIDE a loop/func. They are long lived
  //Don't open/close connections over and over again. Keep it open until finished
  //i.e. pass connection into functions
	defer conn.Close()
  defer rows.Close()

  for rows.Next() {
    //Interesting. I wonder why we need to have the variable reference?
  	err := rows.Scan(&date, &entry)
  	if err != nil {
  		log.Fatal(err)
    }
  	log.Println(date, entry)
  }
  err = rows.Err()
  if err != nil {
	   log.Fatal(err)
   }

}
