package usedataset

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type man_data struct {
	EmpID    int
	Email    string
	Password string
}

type pos_data struct {
	Name          string
	SalaryPerHour float64
}
type shift_data struct {
	EmpID       int
	HoursWorked int
	ShiftDate   string
}

type emp_data struct {
	Firstname  string
	Lastname   string
	ManagerID  int
	PositionID int
}

type DBstruct struct {
	Session *sql.DB
}



func initDB() DBstruct {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	session, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening DB: %s", err.Error())
	}
	if err = session.Ping(); err != nil {
		log.Fatalf("Could not ping: %s", err.Error())
	}
	log.Println("Connected to DB succesfully")
	return DBstruct{session}
}

func PopulateDB() {
	db := initDB()
	db.createDB()
	db.posJSON()
	db.manJSON()
	db.empJSON()
	db.shiJSON()
	db.constraintAndCleanup()
	defer db.Session.Close()
}

func (db DBstruct) createDB() {
	if _, err := db.Session.Exec(`
	
	CREATE SCHEMA "emt";
	
	CREATE TABLE "emt"."manager" (
		"id" int unique primary key GENERATED BY DEFAULT AS IDENTITY (START WITH 888000 INCREMENT BY 1),
		"emp_id" INT UNIQUE NOT NULL,
		"email" VARCHAR(127) UNIQUE NOT NULL,
		"password" VARCHAR NOT NULL,
		"created_at" timestamptz DEFAULT (now()),
		"updated_at" timestamptz DEFAULT (now())
	);
	
	--CREATE SEQUENCE IF NOT EXISTS "emt"."unique_manager_id" START 888000 INCREMENT 5 MINVALUE 888000 OWNED BY "emt"."manager"."id";
	
	CREATE TABLE "emt"."position" (
		"id" int unique primary key GENERATED BY DEFAULT AS IDENTITY (START WITH 60 INCREMENT BY 1),
		"name" VARCHAR(255) UNIQUE NOT NULL,
		"salary_per_hr" NUMERIC NOT NULL,
		"created_at" timestamptz DEFAULT (now()),
		"updated_at" timestamptz DEFAULT (now())
	);
	
	--CREATE SEQUENCE IF NOT EXISTS "emt"."unique_position_id" START 60 INCREMENT 3 MINVALUE 60 OWNED BY "emt"."position"."id";
	
	CREATE TABLE "emt"."employee" (
		"id" int unique primary key GENERATED BY DEFAULT AS IDENTITY (START WITH 222000 INCREMENT BY 1),
		"firstname" VARCHAR(127) NOT NULL,
		"lastname" VARCHAR(127) NOT NULL,
		"position_id" INT DEFAULT (-60),
		"manager_id" INT DEFAULT (-888000),
		"created_at" timestamptz DEFAULT (now()),
		"updated_at" timestamptz DEFAULT (now())
	);
	
	-- CREATE SEQUENCE IF NOT EXISTS "emt"."unique_emp_id" START 222000 INCREMENT 1 MINVALUE 222000 OWNED BY "emt"."employee"."id";
	
	CREATE TABLE "emt"."shift" (
		"id" uuid PRIMARY KEY DEFAULT (gen_random_uuid ()),
		"emp_id" INT NOT NULL PRIMARY KEY,
		"hours_worked" INT NOT NULL,
		"shift_date" DATE NOT NULL,
		"created_at" timestamptz DEFAULT (now()),
		"updated_at" timestamptz DEFAULT (now())
	);
	
	
	
	ALTER TABLE
		"emt"."employee"
	ADD
		FOREIGN KEY ("manager_id") REFERENCES "emt"."manager" ("id") ON DELETE
	SET
		DEFAULT;
	
	ALTER TABLE
		"emt"."employee"
	ADD
		FOREIGN KEY ("position_id") REFERENCES "emt"."position" ("id") ON DELETE
	SET
		DEFAULT;
	
	ALTER TABLE
		"emt"."shift"
	ADD
		FOREIGN KEY ("emp_id") REFERENCES "emt"."employee" ("id") ON DELETE RESTRICT;
	`); err != nil {
		log.Fatalf("Could not execute table creation: %s", err.Error())
	} else {
		log.Println("successful creation of DB")
	}
}
func (db DBstruct) manJSON() {
	managerFile, err := os.Open("../dataset/manager.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened manager.json")

	defer managerFile.Close()
	v, _ := ioutil.ReadAll(managerFile)
	var man []man_data
	if err := json.Unmarshal(v, &man); err != nil {
		log.Fatal(err.Error())
	}
	for _, v := range man {

		if _, err := db.Session.Exec(`insert into emt.manager(emp_id, email, password) values($1, $2, $3)`, v.EmpID, v.Email, v.Password); err != nil {
			log.Fatalf("DB code could not run with success: %v", err.Error())
		} else {
			log.Println("Successful creation:", v)
		}
	}
}

func (db DBstruct) posJSON() {
	positionFile, err := os.Open("../dataset/position.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened position.json")

	defer positionFile.Close()
	v, _ := ioutil.ReadAll(positionFile)
	var pos []pos_data
	if err := json.Unmarshal(v, &pos); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%+v", pos)
	for _, v := range pos {
		if _, err := db.Session.Exec(`insert into emt.position(name, salary_per_hour) values($1, $2)`, v.Name, v.SalaryPerHour); err != nil {
			log.Fatalf("DB code could not run with success: %v", err.Error())
		} else {
			log.Println("Successful creation:", v)
		}
	}
}

func (db DBstruct) empJSON() {
	employeeFile, err := os.Open("../dataset/employee.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened employee.json")

	defer employeeFile.Close()
	v, _ := ioutil.ReadAll(employeeFile)
	var emp []emp_data
	if err := json.Unmarshal(v, &emp); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%+v", emp)
	for _, v := range emp {
		if v.ManagerID > 888006 {
			v.ManagerID = 888006 - v.ManagerID%2
		}
		if _, err := db.Session.Exec(`insert into emt.employee(firstname, lastname, position_id, manager_id) values($1, $2, $3, $4)`, v.Firstname, v.Lastname, v.PositionID, v.ManagerID); err != nil {
			log.Fatalf("DB code could not run with success: %v", err.Error())
		} else {
			log.Println("Successful creation:", v)
		}
	}
}

func (db DBstruct) shiJSON() {
	shiftFile, err := os.Open("../dataset/shift.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened shift.json")

	defer shiftFile.Close()
	v, _ := ioutil.ReadAll(shiftFile)
	var shi []shift_data
	if err := json.Unmarshal(v, &shi); err != nil {
		log.Fatal(err.Error())
	}

	for _, v := range shi {
		layout := "2006-01-02"
		str := v.ShiftDate
		split := strings.Split(str, "T")
		workedAt, _ := time.Parse(layout, split[0])
		if _, err := db.Session.Exec(`insert into emt.shift(emp_id, hours_worked, shift_date) values($1, $2, $3)`, v.EmpID, v.HoursWorked, workedAt); err != nil {
			log.Fatalf("DB code could not run with success: %v", err.Error())
		} else {
			log.Println("Successful creation:", v)
		}
	}
}

func (db DBstruct) constraintAndCleanup() {
	if _, err := db.Session.Exec(`ALTER TABLE "emt"."manager" ADD FOREIGN KEY ("emp_id") REFERENCES "emt"."employee" ("id") ON DELETE CASCADE`); err != nil {
		log.Fatalf("DB code could not run with success: %v", err.Error())
	} else {
		log.Println("Successful addition of constraint")
	}
	if _, err := db.Session.Exec(`DELETE FROM emt.shift WHERE created_at IN (SELECT created_at FROM emt.shift EXCEPT SELECT MIN(created_at) FROM emt.shift GROUP BY emp_id, shift_date )`); err != nil {
		log.Fatalf("DB code could not run with success: %v", err.Error())
	} else {
		log.Println("Successful database cleanup")
	}
}
