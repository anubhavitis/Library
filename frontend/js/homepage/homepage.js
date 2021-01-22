import { CheckWelcome } from "./dashboadAuth.js";
import { loadCard } from "./store.js";
var tokenStr;

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
  if (token == null) document.getElementById("logout").click();

  tokenStr = JSON.parse(token);
  CheckWelcome(tokenStr);

  setTimeout(function () {
    document.getElementById("loading").classList.add("hidden");
    document.getElementById("tabpanel").classList.replace("hidden", "block");
    document
      .getElementById("bookContainer")
      .classList.replace("hidden", "block");
  }, 2000);
};
