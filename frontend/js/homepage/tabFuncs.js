export function NewBook(user, e) {
  e.preventDefault();
  var title = document.getElementById("bookName").value;
  var author = document.getElementById("bookAuthor").value;
  var genre = document.getElementById("bookGenre").value;
  var about = document.getElementById("bookAbout").value;
  var image = document.getElementById("bookImage").value;

  var myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/json");
  var raw = JSON.stringify({
    name: title,
    author: author,
    owner: user,
    genre: genre,
    about: about,
  });
  var requestOptions = {
    method: "POST",
    headers: myHeaders,
    body: raw,
    redirect: "follow",
  };

  fetch("https://libraryz.herokuapp.com/addbook", requestOptions)
    .then((response) => response.json())
    .then((result) => {
      console.log(result);
      if (result.success == true) {
        swal.fire("Ta Da!", "New book added!", "success");
      } else {
        swal.fire("Oh oh!", result.error, "error");
        button.disabled = false;
      }
    });
}
