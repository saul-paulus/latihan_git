var fs = require('fs');
var PDFParser = require('pdf2json');
var path = osHomedir();
var homepath = path.replace(new RegExp('\\' + __dirname, ''), '/');
var pdfFilePath = homepath + 'isi.pdf';


if (fs.existsSync(pdfFilePath)) {
    //Read the content of the pdf from the downloaded path
    var pdfParser = new PDFParser(browser, 1);
    pdfParser.on("pdfParser_dataError", function(errData) {
        console.error(errData.parserError)
    });
    pdfParser.on("pdfParser_dataReady", function(pdfData) {
        //console.log('here is the content: '+pdfParser.getRawTextContent());
        browser.assert.ok(pdfParser.getRawTextContent().indexOf(textToVerify) > -1);
    });

    pdfParser.loadPDF(pdfFilePath);
} else {
    console.log('OOPs file not present in the downloaded folder');
    //Throw an error if the file is not found in the path mentioned
    browser.assert.ok(fs.existsSync(pdfFilePath));
}