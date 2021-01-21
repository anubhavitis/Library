export function Login() {
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
      if (result.success == true){
        localStorage.setItem("token", result.body.token )
        console.log(result.body.token)
        swal.fire("Welcome back!", "Your books waited for you!", "success");
        window.location.href="https://anubhavitis.github.io/Library"
      }
      else {
        swal.fire("Oh oh!", result.error, "error");
      }
    });
}

export function Register() {
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
      if (result.success == true)
        swal.fire(
          "Hello, " + uname,
          "Check your email to complete your registration.",
          "success"
        );
        /*
        TODO
        - Save JWT in local storage 
        */
      else swal.fire("Uh Oh!", result.error, "error");
    });
}

export function GoogleSignUp() {
  
  helper()
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
