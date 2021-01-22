import { Login, Register, GoogleSignUp } from "./authfuncs.js";

document.getElementById("signin").onsubmit = function (e) {
  Login(e);
};
document.getElementById("signup").onsubmit = function (e) {
  Register(e);
};

document.getElementById("forget").onsubmit = function (e) {
  Register(e);
};

document.getElementById("google_signup").onclick = function () {
  GoogleSignUp();
};

window.onload = function () {
  // ToggleToSignIn()
  ToggleToSignIn();
};
