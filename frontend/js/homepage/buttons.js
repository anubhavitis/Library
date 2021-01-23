export var showAllBooks = document.getElementById("showAllBooks");
export var showMyBooks = document.getElementById("showMyBooks");
export var showSavedBooks = document.getElementById("saved");
export var addBooks = document.getElementById("addbook");
export var cur = showAllBooks;

export function addtippy() {
  tippy("#showAllBooks", {
    content: "All books in your library",
    followCursor: "horizontal",
    animation: "shift-away",
  });

  tippy("#showMyBooks", {
    content: "View my books",
    followCursor: "horizontal",
    animation: "shift-away",
  });

  tippy("#saved", {
    content: "Books that you love",
    followCursor: "horizontal",
    animation: "shift-away",
  });

  tippy("#addbook", {
    content: "Add new book to Library",
    followCursor: "horizontal",
    animation: "shift-away",
  });
}

export function AddBook() {
  document.getElementById("bookContainer").classList.add("hidden");
  document.getElementById("addBookDiv").classList.remove("hidden");
  removeActive(cur);
  cur = document.getElementById("addbook");
  Active(cur);
}

export function AllBooks() {
  removeActive(cur);
  cur = showAllBooks;
  Active(cur);
  document.getElementById("addBookDiv").classList.remove("hidden");
  document.getElementById("addBookDiv").classList.add("hidden");
  document.getElementById("bookContainer").classList.remove("hidden");
}

export function SavedBooks() {
  removeActive(cur);
  cur = showSavedBooks;
  Active(cur);
  document.getElementById("addBookDiv").classList.remove("hidden");
  document.getElementById("addBookDiv").classList.add("hidden");
  document.getElementById("bookContainer").classList.remove("hidden");
}

export function MyBooks() {
  removeActive(cur);
  cur = showMyBooks;
  Active(cur);
  document.getElementById("addBookDiv").classList.remove("hidden");
  document.getElementById("addBookDiv").classList.add("hidden");
  document.getElementById("bookContainer").classList.remove("hidden");
}

export function removeActive(div) {
  div.classList.remove("shadow-xl");
  div.classList.remove("ring-4");
  div.classList.remove("ring-gray-300");
}

export function Active(div) {
  div.classList.add("shadow-xl");
  div.classList.add("ring-4");
  div.classList.add("ring-gray-300");
}
