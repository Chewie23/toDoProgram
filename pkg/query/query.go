package query

import (
  "fmt"
  "time"
  "log"
  "toDoProgram/internal/db"
  "database/sql"
)
// TODO:
//Figure out WHERE to defer db.Close() or stmt.Close(), since currently, this closes
//the connection everytime you call the method "PrintRow()"
type query struct {
  //TODO
  //I need to rethink how to structure this. Since the db connection needs to be
  //longed lived, and thus either needs to be a global OR this is the highest level
  //and everything lives under this. Basically keep connections open, and pass it around
  //But for this case? I think having everything live under this struct will work
  //Since a ToDo app is literally get and set stuff into the DB
  conn *sql.DB
  stmt *sql.Stmt
  err error
}

func NewQuery() *query {
  q := new(query)
  q.conn = db.GetLogin()
  return q
}

func (q query) CloseDB() error {
    //Will be able to call this in the main, as
    //defer q.CloseDB()
    return q.conn.Close()
}

func (q *query) SetStatement(db_str string) {
  //Fun fact: In Golang, there are pointer "receivers" (the "q *query" here).
  //This makes the method be able to read AND write to a struct, instead of just read
  //If you want read/write access without it, then you need a higher level function
  //to call that function (like with the case of "NewQuery", but even then, that
  //has a pointer return)

  //See:
  //https://nathanleclaire.com/blog/2014/08/09/dont-get-bitten-by-pointer-vs-non-pointer-method-receivers-in-golang/
  //https://stackoverflow.com/questions/27775376/value-receiver-vs-pointer-receiver

  //In general, if you aren't using concurrency, then if you have ONE pointer receiver then
  //for consistency, just have it all pointer receivers. This is just for sake of ease and readability
  //BUT, in terms of this case? I think it'll be fine to leave as is. Since we only need a pointer
  //receiver to affect the struct data, the rest will be alright to work with a value receiver
  //(this is because the struct data itself is unaffected, but the SQL pointer struct is affected, but we
  //don't care at this scope)

  q.stmt, q.err = q.conn.Prepare(db_str)
  if q.err != nil {
    log.Fatal(q.err)
  }
}

func (q query) PrintRows() {

  var (
    date time.Time
    entry string
  )

  var db_str = "SELECT date, entry FROM to_do"

  q.SetStatement(db_str)

  //This is super weird. It is assuming that "q.conn" will be there
  //when it is called. I'm more used to the Python aspect of you can define
  //the instance variable (or whatever "q.conn" is) as soon as you declare it

  rows, err := q.stmt.Query()

  if err != nil {
    log.Fatal(err)
  }
  //Remember to always close your connections, as a defer, OUTSIDE a loop/func. They are long lived
  //Don't open/close connections over and over again. Keep it open until finished
  //i.e. pass connection into functions
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

func (q query) AddEntry(entry_str string) {
  //TODO
  //Figure out how to insert a Golang time stamp INTO MySQL
  //Maybe helpful:
  //http://rafalgolarz.com/blog/2017/08/27/mysql_timestamps_in_go/

  //This is a placeholder and will not work as is
  var db_str = fmt.Sprintf("INSERT INTO to_do(date, entry) VALUES(<DATE>, %s)", entry_str)
  fmt.Println(db_str)
}
