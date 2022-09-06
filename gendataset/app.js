const fs = require("fs");
const { pos, emp, man, shift } = require("./dataset.model");
fs.mkdir("../dataset", { recursive: true }, (err) => {
  if (err) throw err;
  console.log("dataset directory created");
});
function writefiles(filename, dataArr) {
  fs.writeFile(
    `../dataset/${filename}.json`,
    JSON.stringify(dataArr),
    (err, _) => {
      if (err) throw err;
      console.log(`${filename} JSON is saved.`);
    }
  );
}
//pos, man, emp, shift
const fileObject = new Map([
  ["position", pos],
  ["manager", man],
  ["employee", emp],
  ["shift", shift],
]);
fileObject.forEach(function (data_Arr, file_name) {
  writefiles(file_name, data_Arr);
});
