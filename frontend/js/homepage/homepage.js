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
  username,
} from "./buttons.js";

import { loadCard } from "./store.js";
import { NewBook } from "./tabFuncs.js";

window.onload = function () {
  for (var i = 0; i < 9; i += 1)
    document.getElementById("Lcontainer").innerHTML += loadCard;

  AllBooks();
  addtippy();
};

addBooks.onclick = function () {
  AddBook();
};
showAllBooks.onclick = function () {
  AllBooks();
};

showMyBooks.onclick = function () {
  MyBooks();
};

showSavedBooks.onclick = function () {
  SavedBooks();
};

document.getElementById("bookForm").onsubmit = function (e) {
  NewBook(username, e);
};
