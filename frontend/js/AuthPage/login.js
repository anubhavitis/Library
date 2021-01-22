document.getElementById("signin").onsubmit = function (e) {
  Login(e);
};

function Login(e) {
  e.preventDefault();
  var username = document.getElementById("login_uname").value;
  var password = document.getElementById("login_password").value;

  var myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/json");
  var raw = JSON.stringify({ username: username, password: password });
  var requestOptions = {
    method: "POST",
    headers: myHeaders,
    body: raw,
    redirect: "follow",
  };

  fetch("https://libraryz.herokuapp.com/signin", requestOptions)
    .then((response) => response.json())
    .then((result) => {
      if (result.success == true) {
        localStorage.setItem("token", JSON.stringify(result.body.token));
        swal.fire("Welcome back!", "Your books waited for you!", "success");
        window.location.href = "https://anubhavitis.github.io/Library";
        // window.location.href = "http://127.0.0.1:5500/index.html";
      } else {
        swal.fire("Oh oh!", result.error, "error");
      }
    });
}
