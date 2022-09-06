const { faker } = require("@faker-js/faker");
const pos = new Array();
for (let i = 0; i < 10; i++) {
  pos.push(
    ({
      Name: faker.name.jobTitle(),
      SalaryPerHour: +faker.finance.amount(2000, 4000, 2),
    })
  );
}

const emp = new Array();
for (let i = 0; i < 50; i++) {
  emp.push(
    ({
      Firstname: faker.name.firstName(),
      Lastname: faker.name.lastName(),
      PositionID: faker.mersenne.rand(70, 60),
      ManagerID: faker.mersenne.rand(888010, 888000),
    })
  );
}

var a = 222000;
const man = new Array();
for (let i = 0; i < 7; i++) {
  man.push(
    ({
      EmpID: a + i,
      Email: faker.internet.email(
        emp[i].Firstname,
        emp[i].Lastname,
        "emtcorp.com",
        { allowSpecialCharacters: false }
      ),
      Password: faker.internet.password(
        faker.datatype.number({ min: 9, max: 15 }),
        false
      ),
    })
  );
}

const shift = new Array();
for (let i = 0; i < 1000; i++) {
  shift.push(
    ({
      EmpID: faker.mersenne.rand(222050, 222000),
      HoursWorked: faker.datatype.number({ min: 9, max: 15 }),
      ShiftDate: faker.date.between("2022-05-15", "2022-07-12"),
    })
  );
}

module.exports = {
  pos,
  emp,
  man,
  shift,
};
