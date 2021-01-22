export function CheckWelcome(token) {
  var myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/json");
  var raw = JSON.stringify({ token: token });

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
        document.getElementById("logout").click()
      }
    });
}
