var db = require("./db_config");

db.connect(function(err) {
    if (err) throw err;
    let sql = `CREATE TABLE customer_bank (
                cust_id int(11) NOT NULL AUTO_INCREMENT,
                nama varchar(50) NOT NULL,
                alamat varchar(200) NOT NULL,
                kode_pos char(5) DEFAULT NULL,
                no_hp varchar(15) DEFAULT NULL,
                email varchar(50) DEFAULT NULL,
                PRIMARY KEY (cust_id)
                )`;
    db.query(sql, function(err, result) {
        if (err) throw err;
        console.log("Table created");
    });
});