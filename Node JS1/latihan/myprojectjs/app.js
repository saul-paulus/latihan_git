//use path module
const path = require("path"); //Choose your project path
//use express module
const express = require('express'); //npm install express
//use hbs view engine
const hbs = require("hbs"); //npm install hbs
const app = express();
//set views file
app.set("views", path.join(__dirname, "views"));
//set view engine
app.set("view engine", "hbs");
//set public folder as static folder for static file
app.use(express.static("public"));
//route untuk halaman home
app.get("/", (req, res) => {
    //render file index.hbs
    res.render("index", { name: "Komang" });
});
//route untuk halaman home dengan parameter name
app.get("/:name", (req, res) => {
    //render file index.hbs
    res.render("index", { name: req.params.name });
});

app.listen(8000, () => {
    console.log("Server is running at port 8000");
});