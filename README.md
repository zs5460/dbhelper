# dbhelper

dbhelper is a paging query tool written for the old version of MSSQL

## Install

```bash
go get github.com/zs5460/dbhelper
```

## Example

```go
package main

import (
    "fmt"

    _ "github.com/denisenkom/go-mssqldb"
    "github.com/jmoiron/sqlx"
    "github.com/zs5460/dbhelper"

)

var db *sqlx.DB

// Article ...
type Article struct {
    ID      int    `db:"id"`
    Title   string `db:"title"`
    Image   string `db:"image"`
    Content string `db:"content"`
}

func main() {
    db = sqlx.MustConnect("mssql", "server=x.x.x.x;port=1433;database=xxx;user id=xxx;password=xxx;encrypt=disable")

    articles := []Article{}

    err := dbhelper.GetPage(
        db,               //*sqlx.DB
        &articles,        //query result
        "id,title,image", //fields
        "article",        //table
        "siteid=2",       //where
        "id",             //orderby
        10,               //pagesize
        1)                //pageindex
    if err != nil {
        fmt.Println(err)
    }

    for _, a := range articles {
        fmt.Println(a)
    }

}


```

## License

Released under MIT license, see [LICENSE](LICENSE) for details.
