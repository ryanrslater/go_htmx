function pageLoad() {
  console.log("is loaded");

  document.getElementById("title").addEventListener("click", () => {
    console.log("clicked");
    document.getElementById("content").innerHTML = "You clicked";
  });
}

window.addEventListener("load", pageLoad);
