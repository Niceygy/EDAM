function toggleLoadingBtn() {
  var x = document.getElementById("loadingIcon");
  if (x.style.display === "none") {
    x.style.display = "block";
  } else {
    x.style.display = "none";
  }
}

window.onload = function () {
  document.getElementById("loadingIcon").style.display = "none";
};
