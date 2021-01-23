document.getElementById("signup").onsubmit = function (e) {
  Register(e);
};

var loadingDiv = `<div class="flex justify-center"> 
<svg class="h-6 w-6 animate-pulse text-white" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
<path d="M11 3a1 1 0 10-2 0v1a1 1 0 102 0V3zM15.657 5.757a1 1 0 00-1.414-1.414l-.707.707a1 1 0 001.414 1.414l.707-.707zM18 10a1 1 0 01-1 1h-1a1 1 0 110-2h1a1 1 0 011 1zM5.05 6.464A1 1 0 106.464 5.05l-.707-.707a1 1 0 00-1.414 1.414l.707.707zM5 10a1 1 0 01-1 1H3a1 1 0 110-2h1a1 1 0 011 1zM8 16v-1h4v1a2 2 0 11-4 0zM12 14c.015-.34.208-.646.477-.859a4 4 0 10-4.954 0c.27.213.462.519.476.859h4.002z" />
</svg>
<svg class="h-6 w-6 animate-pulse text-white" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
<path d="M11 3a1 1 0 10-2 0v1a1 1 0 102 0V3zM15.657 5.757a1 1 0 00-1.414-1.414l-.707.707a1 1 0 001.414 1.414l.707-.707zM18 10a1 1 0 01-1 1h-1a1 1 0 110-2h1a1 1 0 011 1zM5.05 6.464A1 1 0 106.464 5.05l-.707-.707a1 1 0 00-1.414 1.414l.707.707zM5 10a1 1 0 01-1 1H3a1 1 0 110-2h1a1 1 0 011 1zM8 16v-1h4v1a2 2 0 11-4 0zM12 14c.015-.34.208-.646.477-.859a4 4 0 10-4.954 0c.27.213.462.519.476.859h4.002z" />
</svg> 
<p class="text-white font-bold animate-pulse"> LOADING </p>
<svg class="h-6 w-6 animate-pulse text-white" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
<path d="M11 3a1 1 0 10-2 0v1a1 1 0 102 0V3zM15.657 5.757a1 1 0 00-1.414-1.414l-.707.707a1 1 0 001.414 1.414l.707-.707zM18 10a1 1 0 01-1 1h-1a1 1 0 110-2h1a1 1 0 011 1zM5.05 6.464A1 1 0 106.464 5.05l-.707-.707a1 1 0 00-1.414 1.414l.707.707zM5 10a1 1 0 01-1 1H3a1 1 0 110-2h1a1 1 0 011 1zM8 16v-1h4v1a2 2 0 11-4 0zM12 14c.015-.34.208-.646.477-.859a4 4 0 10-4.954 0c.27.213.462.519.476.859h4.002z" />
</svg><svg class="h-6 w-6 animate-pulse text-white" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
<path d="M11 3a1 1 0 10-2 0v1a1 1 0 102 0V3zM15.657 5.757a1 1 0 00-1.414-1.414l-.707.707a1 1 0 001.414 1.414l.707-.707zM18 10a1 1 0 01-1 1h-1a1 1 0 110-2h1a1 1 0 011 1zM5.05 6.464A1 1 0 106.464 5.05l-.707-.707a1 1 0 00-1.414 1.414l.707.707zM5 10a1 1 0 01-1 1H3a1 1 0 110-2h1a1 1 0 011 1zM8 16v-1h4v1a2 2 0 11-4 0zM12 14c.015-.34.208-.646.477-.859a4 4 0 10-4.954 0c.27.213.462.519.476.859h4.002z" />
</svg></div>`;

function Register(e) {
  e.preventDefault();

  var butt = document.getElementById("register");
  butt.innerHTML = loadingDiv;

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
        localStorage.setItem("token", JSON.stringify(result.body.token));
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
