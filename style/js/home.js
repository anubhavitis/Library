import {
    ToggleToForget, ToggleToSignIn, ToggleToSignUp
} from "./changeCard.js"

import {
    Login, Register, GoogleSignUp
} from "./auth.js"

document.getElementById("goToSignInf").onclick= function(){ ToggleToSignIn(); };
document.getElementById("goToSignUpf").onclick = function(){ ToggleToSignUp(); }
document.getElementById("goToSignUp").onclick = function(){ ToggleToSignUp(); }
document.getElementById("goToSignIn").onclick = function(){ ToggleToSignIn(); }
document.getElementById("goToForget").onclick = function(){ ToggleToForget(); }

document.getElementById("login").onclick= function() { Login(); }
document.getElementById("register").onclick= function() { Register(); }

document.getElementById("google_signup").onclick= function() { GoogleSignUp(); }

window.onload = function () {
    // ToggleToSignIn()
    ToggleToSignIn()
}

