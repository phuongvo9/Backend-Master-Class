package db

// All unit tests will need a object "Queries" => declare as a global variable
// {Queries struct} contans DBTX which can be a database connection or a database transaction.
// In these test cases, we will use testQueries as a db connection and use it to create a Queries object
var testQueries *Queries
