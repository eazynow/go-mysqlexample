Go Mysql Sample app

Small app to test Go mysql connectivity

Make sure you install mysql drivers first:

    go get github.com/go-sql-driver/mysql

Be sure to set up your gopath first. See go help gopath for details

then just run with

    go run mysqlexample.go

It expects a mysql server at 176.16.200.100 (vagrant db0) with voice database