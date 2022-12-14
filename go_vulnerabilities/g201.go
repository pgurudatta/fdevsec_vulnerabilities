//G202 - SQL query construction using string concatenation
//https://securego.io/docs/rules/g201-g202.html

package main
import (
    "database/sql"
    "fmt"
    "os"
)
func main(){
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        panic(err)
    }
    q := fmt.Sprintf("SELECT * FROM foo where name = '%s'", os.Args[1])
    rows, err := db.Query(q)
    if err != nil {
        panic(err)
    }
    defer rows.Close()
}
