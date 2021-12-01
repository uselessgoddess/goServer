const { exec } = require("child_process");
const path = require("path");

console.log(__dirname)
const forked = exec("go run server.go", { cwd: __dirname }, (error, stdout, stderr) => {
    if (error) {
        console.log(`error: ${error.message}`);
        return;
    }
    if (stderr) {
        console.log(`stderr: ${stderr}`);
        return;
    }
    console.log(`stdout: ${stdout}`);
});
console.log(path.join(__dirname, "..", "/result/server.go"))
forked.on("exit", () => console.log("exited"))

