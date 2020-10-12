var db = require("./db_config");

db.connect(function(err) {
    if (err) throw err;
    let sql = `CREATE TABLE mahasiswa
            (
                npm CHAR(8) NOT NULL,
                nama VARCHAR(255) NOT NULL,
                kelas CHAR(5) NOT NULL,
                alamat VARCHAR(255),
                PRIMARY KEY(npm)
            )`;
    db.query(sql, function(err, result) {
        if (err) throw err;
        console.log("Table created");
    });
});