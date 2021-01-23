import { CheckWelcome } from "./dashboadAuth.js";
import { loadCard } from "./store.js";

var username, fname, sname, image;

tippy("#showAllBooks", {
  content: "View all books",
  followCursor: "horizontal",
  animation: "shift-away",
});

tippy("#showMyBooks", {
  content: "View my books",
  followCursor: "horizontal",
  animation: "shift-away",
});

window.onload = function () {
  for (var i = 0; i < 9; i += 1)
    document.getElementById("Lcontainer").innerHTML += loadCard;

  var token = localStorage.getItem("token");
  if (token == null) {
    document.getElementById("logout").click();
  }

  var tokenStr = JSON.parse(token);
  var myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/json");
  var raw = JSON.stringify({ token: tokenStr });

  var requestOptions = {
    method: "POST",
    headers: myHeaders,
    body: raw,
    redirect: "follow",
  };

  fetch("https://libraryz.herokuapp.com/welcome", requestOptions)
    .then((response) => response.json())
    .then((result) => {

      if (!result.success) {
        document.getElementById("logout").click();
      } else {
        username = result.body.username;
        fname = result.body.fname;
        sname = result.body.sname;

        var title = document.createElement("title");
        title.innerHTML = fname + "'s Dashboard";
        document.getElementById("loading").classList.add("hidden");
        document
          .getElementById("tabpanel")
          .classList.replace("hidden", "block");
        document
          .getElementById("bookContainer")
          .classList.replace("hidden", "block");
      }
    });
};
