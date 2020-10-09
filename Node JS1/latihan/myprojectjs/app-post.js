//use path module
const path = require("path"); //Choose your project path
//use express module
const express = require('express'); //npm install express
//use hbs view engine
const hbs = require("hbs"); //npm install hbs
const app = express();

var bodyParser = require('body-parser'); //npm install body-parser
// Create application/x-www-form-urlencoded parser  
var urlencodedParser = bodyParser.urlencoded({ extended: false })
    //set views file
app.set("views", path.join(__dirname, "views"));
//set view engine
app.set("view engine", "hbs");
//set public folder as static folder for static file
app.use(express.static("public"));
//route untuk halaman home
app.get("/post", (req, res) => {
    //render file views/post.hbs
    res.render("post");
});
//route untuk halaman home dengan parameter name
app.post('/post', urlencodedParser, function(req, res) {
    // Prepare output in JSON format  
    response = {
        name: req.body.name
    };
    console.log(response); //print to console log
    res.render("index", { name: req.body.name }); //render index.hbs and replace name value with post name input user
})

app.listen(8000, () => {
    console.log("Server is running at port 8000"); //start server
});