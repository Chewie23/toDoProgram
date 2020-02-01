package main

import "toDoProgram/pkg/query"
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

  q := query.NewQuery()
  defer q.CloseDB()
  q.PrintRows()



  //A Prepared Statement. Good for security and reusability, but if the statement
  //isn't used again, maaaay not be worth the overhead
  //In this case, it's probably not worth it

}
