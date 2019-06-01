goagen -d github.com/IkuEisou/goa-adder/design app
goagen -d github.com/IkuEisou/goa-adder/design swagger
rm -f main.go
goagen -d github.com/IkuEisou/goa-adder/design main
go build
./goa-adder
