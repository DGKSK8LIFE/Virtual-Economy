go mod tidy 
go mod download
mongod --port 27017
go run main.go