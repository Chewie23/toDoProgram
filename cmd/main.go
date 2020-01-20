package main

import "fmt"

/*
So, Golang depends on drivers to access the SQL. Seems like I need the under-
lying MySQL to do things. Which makes sense.

Following this tutorial:
http://go-database-sql.org/index.html

And using best practices with a config file for SQL access
*/
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "github.com/spf13/viper"


func main(){

  //May be able to put in a function? But for now, just reads the config
  //file
  viper.SetConfigName("config")
  viper.AddConfigPath("$GOPATH/src/toDoProgram")
  viper.SetConfigType("yml")

  if err := viper.ReadInConfig(); err != nil {
    fmt.Printf("Error reading config file, %s", err)
  }

  mysqluser, mysqlpass, mysqlport, mysqldb := viper.GetString("database.dbuser"),
  viper.GetString("database.dbpassword"), viper.GetString("server.port"), viper.GetString("database.dbname")

  dbstr := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s", mysqluser, mysqlpass, mysqlport, mysqldb)

  db, err := sql.Open("mysql", dbstr)
	if err != nil {
		fmt.Println("Could not open")
	}
  err = db.Ping()
  if err != nil {
  	fmt.Println("Something went wrong pinging db")
    fmt.Println(err)
  }
  //Remember to always close your connections. They are long lived
	defer db.Close()
  fmt.Println("Hello World")
}
