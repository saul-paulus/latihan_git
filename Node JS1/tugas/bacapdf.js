const fs = require('fs') //load file system framework
const pdfparse = require('pdf-parse') //load pdf parser framework. kita harus menginstall terlebih dahulu ke alamat direktorinya dengan mengetik npm i pdf-parse
const pdffile = fs.readFileSync('isi.pdf') // membaca file isi.pdf yang ada pada direktori
    // get the information
pdfparse(pdffile).then(function(data) { // load seluruh properti data pdf isi    
    // number of pages
    console.log(data.numpages);
    // number of rendered pages
    console.log(data.numrender);
    // PDF info
    console.log(data.info);
    // PDF metadata
    console.log(data.metadata);
    // PDF.js version    
    console.log(data.version);
    // PDF text
    console.log(data.text);
})

//https://www.youtube.com/watch?reload=9&v=THauKA1p7po