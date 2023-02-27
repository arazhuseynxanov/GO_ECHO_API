# GO_ECHO_API
//===============================================//
To install Echo Go v1.13 or higher is required. Go v1.12 has limited support and some middlewares will not be available. Make sure your project folder is outside your $GOPATH.

$ mkdir myapp && cd myapp
$ go mod init myapp
$ go get github.com/labstack/echo/v4

//===============================================//
Install GORM and SQLite
go get -u gorm.io/gorm 
go get -u gorm.io/driver/sqlite
