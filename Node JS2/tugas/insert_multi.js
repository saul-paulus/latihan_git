var db = require("./db_config");
db.connect(function(err) {
    if (err) throw err;

    let sql = "INSERT INTO mahasiswa (npm, nama, kelas, alamat) VALUES ?";
    var values = [
        ['12345677', 'Asep', '4AA01', 'Bandung'],
        ['12345676', 'Febri', '4AA01', 'Bekasi'],
        ['12345675', 'Putri', '4AA01', 'Jakarta'],
        ['12345674', 'Amelia', '4AA01', 'Depok '],
        ['12345673', 'Gilang', '4AA01', 'Jakarta']
    ];
    db.query(sql, [values], function(err, result) {
        if (err) throw err;
        console.log("1 record inserted");
    });
});