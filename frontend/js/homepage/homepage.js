import {
  addtippy,
  AddBook,
  showAllBooks,
  addBooks,
  AllBooks,
  MyBooks,
  SavedBooks,
  showMyBooks,
  showSavedBooks,
} from "./buttons.js";

import { loadCard } from "./store.js";
import { NewBook } from "./tabFuncs.js";

var username, fname, sname, image;

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

        document.getElementById("title").innerHTML = fname + "'s Dashboard";
        document.getElementById("loading").classList.add("hidden");
        document.getElementById("tabpanel").classList.remove("hidden");
        document.getElementById("bookContainer").classList.remove("hidden");
        addtippy();
        AllBooks();
      }
    });
};

addBooks.onclick = function () {
  AddBook();
};
showAllBooks.onclick = function () {
  AllBooks();
};

showMyBooks.onclick= function(){
  MyBooks();
};

showSavedBooks.onclick= function(){
  SavedBooks();
}

document.getElementById("bookForm").onsubmit= function(e){
  NewBook(username, e);
}