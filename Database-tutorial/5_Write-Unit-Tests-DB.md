# Write unit tests for CRUD

## Convention
1. Unit test files in Go put into the same folder as the function files

2. The test file should have the same name with the suffix test(main.go => main_test.go) and same package
3. 

## Use Go Testify package


## Use cases in this turtorial
1. The func CreateAccount needs to connect to the database FIRST before creating an account

2. We will need to establish the connection first => Where to write this code? => main_test in package "db"

3. All unit tests will need a object "Queries" => declare as a global variable
4. 