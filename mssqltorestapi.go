/*
//
1.go get  github.com/denisenkom/go-mssqldb
2.go get  github.com/gorilla/mux

*/

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/julienschmidt/httprouter"
)

var (
	key         string
	customerid  string
	companyname string
)

type jsongroup struct {
	Jcustomerid  string
	Jcompanyname string
}

func main() {

	r := httprouter.New()
	r.GET("/:name", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		key := ps.ByName("name")
		condb, errdb := sql.Open("mssql", "server=.;database=northwind;User ID=****;Password=****;")
		if errdb != nil {
			fmt.Println("ERROR: ", errdb.Error())
		}
		rows, err := condb.Query("select CustomerID,CompanyName from customers where CustomerID=?", key)
		if err != nil {
			log.Fatal(err)
		}
		defer condb.Close()

		for rows.Next() {
			err := rows.Scan(&customerid, &companyname)
			if err != nil {
				log.Fatal(err)
			}

			group := jsongroup{
				Jcustomerid:  customerid,
				Jcompanyname: companyname,
			}

			b, err := json.Marshal(group)
			if err != nil {
				fmt.Println(err)
			}

			log.Println(string(b))
			//os.Stdout.Write(b)
			fmt.Fprintln(w, string(b))

		}
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}

/*
CustomerID	CompanyName
ALFKI	Alfreds Futterkiste
ANATR	Ana Trujillo Emparedados y helados
ANTON	Antonio Moreno Taquería
AROUT	Around the Horn
BERGS	Berglunds snabbköp
*/
