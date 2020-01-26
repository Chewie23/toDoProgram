package db

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "github.com/spf13/viper"

import "log"
import "fmt"

func GetLogin() *sql.DB{
  //May be able to put in a function? But for now, just reads the config
  //file
  viper.SetConfigName("config")
  viper.AddConfigPath("$GOPATH/src/toDoProgram")
  viper.SetConfigType("yml")

  if err := viper.ReadInConfig(); err != nil {
    log.Fatal(err)
  }

  mysqluser, mysqlpass, mysqlport, mysqldb := viper.GetString("database.dbuser"),
  viper.GetString("database.dbpassword"), viper.GetString("server.port"),
  viper.GetString("database.dbname")

  dbstr := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?parseTime=true", mysqluser,
  mysqlpass, mysqlport, mysqldb)

  db, err := sql.Open("mysql", dbstr)
  if err != nil {
    log.Println("Could not open")
    log.Fatal(err)
  }
  err = db.Ping()
  if err != nil {
    log.Println("Something went wrong pinging db")
    log.Fatal(err)
  }

  return db
}
