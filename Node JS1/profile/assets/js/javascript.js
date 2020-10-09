function ColNavBar() {
  var x = document.getElementById("NavAtas");
  if (x.className === "navatas") {
    x.className += " responsive";
  } else {
    x.className = "navatas";
  }
}
