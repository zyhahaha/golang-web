package test

// gorm   github.com/jinzhu/gorm
import (
	"fmt"
	// "log"
	// "net/http"
	// "github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func main() {
    var err error
		Db, err = sql.Open("mysql", "root:password@tcp(47.96.234.53:3306)/test")
    if err != nil {
			fmt.Print(err.Error())
		}
		err = Db.Ping() // 校验连接
		if err != nil {
			fmt.Print(err.Error())
		}
		defer Db.Close()
		
		stmt, err := Db.Prepare("CREATE TABLE person (id int NOT NULL AUTO_INCREMENT, first_name varchar(40), last_name varchar(40), PRIMARY KEY (id));")
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = stmt.Exec()
		if err != nil {
			fmt.Print(err.Error())
		} else {
			fmt.Printf("Person Table successfully migrated....")
		}
}
