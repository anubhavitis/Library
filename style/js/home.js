import {
    ToggleToForget, ToggleToSignIn, ToggleToSignUp
} from "./authCard.js"

document.getElementById("goToSignInf").onclick= function(){ ToggleToSignIn(); };
document.getElementById("goToSignUpf").onclick = function(){ ToggleToSignUp(); }
document.getElementById("goToSignUp").onclick = function(){ ToggleToSignUp(); }
document.getElementById("goToSignIn").onclick = function(){ ToggleToSignIn(); }
document.getElementById("goToForget").onclick = function(){ ToggleToForget(); }

window.onload = function () {
    // ToggleToSignIn()
    ToggleToSignIn()
}

