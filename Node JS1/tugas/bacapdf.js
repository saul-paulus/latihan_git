const fs = require('fs') //load file system framework
const pdfparse = require('pdf-parse') //load pdf parser framework. kita harus menginstall terlebih dahulu ke alamat direktorinya dengan mengetik npm i pdf-parse
const pdffile = fs.readFileSync('isi.pdf') // membaca file isi.pdf yang ada pada direktori
    // get the information
pdfparse(pdffile).then(function(data) { // load seluruh properti data pdf isi
    console.log(data.numpages) //jumlah halaman
    console.log(data.info) // informasi seputar file pdf
    console.log(data.text) //teks yang ada di dalam konten pdf
    console.log(data.version) //versi pdf
    console.log(data.metadata)
    console.log(data.numrender)
})

//https://www.youtube.com/watch?reload=9&v=THauKA1p7po