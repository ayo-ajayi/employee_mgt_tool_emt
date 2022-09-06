# Employee-Management-Tool (emt)


Managers can register employee as well as their daily activities like days worked and number of hours worked on those days.
Critical decisions can be made about an employee based on the data available about such employee.
Employee salary can also be calculated and paid using the salary scale for each position.
Fake data was used to populate the DB. The data was generated using fakerjs package in NodeJS runtime.


## we use:
1. gin-gonic HTTP web framework 
2. postgresql db
3. faker js


## code execution
Ensure that nodejs and npm are installed on your machine. 
Install  the required packages using:
```sh
npm install
```
Then proceed to generate the required JSON files by running:
```sh
node ./gendataset/app.js
```
Execute the main program using the main.go file:
```sh
go run main.go
```
Alternatively:
```sh
go build main.go
./main
```

Thank you ❤
 


