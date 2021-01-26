export var showAllBooks = document.getElementById("showAllBooks");
export var showMyBooks = document.getElementById("showMyBooks");
export var showSavedBooks = document.getElementById("saved");
export var addBooks = document.getElementById("addbook");
export var cur = showAllBooks;

export var username, fname, sname, image;

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

  fetch("http://localhost:8000/welcome", requestOptions)
    .then((response) => response.json())
    .then((result) => {
      if (!result.success) {
        document.getElementById("logout").click();
      } else {
        console.log(result);
        username = result.body.username;
        fname = result.body.fname;
        sname = result.body.sname;
        var arr = result.body.book;
        var bookCards = "";
        arr.forEach((ele) => {
          var Card = `
          <div class="m-4 h-120 rounded-lg shadow-xl flex ring-2 ring-gray-200 bg-gray-100 hover:bg-gray-200">
            <div class="w-1/2 h-100">
              <a href="#">
                <img class=" h-64 w-full rounded-l-lg" src=" ${ele.image}" />
              </a>
            </div>
            <div class="p-4 w-1/2 rounded-r-lg text-left flex flex-col justify-between">
              <div class="flex justify-between">
                <p class="text-gray-600 cursor-pointer"> ${ele.name} </p>
              </div>

              <p class="mt-4 text-gray-400 uppercase text-sm font-bold"> ${
                ele.author
              } </p>
              <p class="text-gray-500  text-sm"> ${ele.genre} </p>

              <div class="flex">
                      <div class="h-6 w-6 text-yellow-500"> 
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                          <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                        </svg>
                      </div>
                      <div class="h-6 w-6 text-yellow-500"> 
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                          <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                        </svg>
                      </div>
                      <div class="h-6 w-6 text-gray-400"> 
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                          <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                        </svg>
                      </div>
              </div>

              <div class="flex justify-between">
                <svg class="w-6 h-6 ${
                  ele.likes ? `text-red-500` : `text-gray-300`
                } mt-auto mb-auto cursor-pointer" xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd"
                    d="M3.172 5.172a4 4 0 015.656 0L10 6.343l1.172-1.171a4 4 0 115.656 5.656L10 17.657l-6.828-6.829a4 4 0 010-5.656z"
                    clip-rule="evenodd" />
                </svg>
                <a href="#">
                  <img class="rounded-full h-12 w-12 border-2 border-blue-200" src=" ${
                    ele.owner
                  } " />
                </a>
              </div>
            </div>
          </div>
          `;

          bookCards+=Card
        });
        document.getElementById("bookContainer").innerHTML = bookCards;
        document.getElementById("title").innerHTML = fname + "'s Dashboard";
        document.getElementById("loading").classList.add("hidden");
        document.getElementById("tabpanel").classList.remove("hidden");
        document.getElementById("bookContainer").classList.remove("hidden");
        addtippy();
      }
    });
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
