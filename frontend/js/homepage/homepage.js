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

  tokenStr = localStorage.getItem("token");
  console.log("token: " + tokenStr);
  if (tokenStr == null) document.getElementById("logout").click();

  var login = CheckWelcome(tokenStr == null ? "" : tokenStr);
  if (login == false) document.getElementById("logout").click();

  setTimeout(function () {
    document.getElementById("loading").classList.add("hidden");
    document.getElementById("tabpanel").classList.replace("hidden", "block");
    document
      .getElementById("bookContainer")
      .classList.replace("hidden", "block");
  }, 2000);
};
