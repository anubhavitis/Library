document.getElementById("signup").onsubmit = function (e) {
  Register(e);
};


function Register(e) {
  e.preventDefault();
  var email = document.getElementById("new_email").value;
  var uname = document.getElementById("new_uname").value;
  var fname = document.getElementById("new_fname").value;
  var sname = document.getElementById("new_sname").value;
  var college = document.getElementById("new_college").value;
  var password = document.getElementById("new_password").value;
  var cpassword = document.getElementById("new_cpassword").value;

  if (password != cpassword) {
    swal.fire("Passwords do not match", "Please try again.", "error");
    return;
  }

  var myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/json");

  var raw = JSON.stringify({
    username: uname,
    fname: fname,
    lname: sname,
    email: email,
    college: college,
    password: password,
  });

  var requestOptions = {
    method: "POST",
    headers: myHeaders,
    body: raw,
    redirect: "follow",
  };

  fetch("https://libraryz.herokuapp.com/signup", requestOptions)
    .then((response) => response.json())
    .then((result) => {
      console.log(result);
      if (result.success == true) {
        swal.fire(
          "Hello, " + uname,
          "Check your email to complete your registration.",
          "success"
        );
        localStorage.setItem("token", result.body.token);
      } else swal.fire("Uh Oh!", result.error, "error");
    });
}

function GoogleSignUp() {
  helper();
}

const helper = async () => {
  const { value: url } = await Swal.fire({
    title: "Please enter your college",
    input: "text",
    inputLabel: "Your college name",
    inputPlaceholder: "E.g. JSS Noida",
  });

  console.log(url);
};
