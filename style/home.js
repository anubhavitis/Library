var SignInTab = document.getElementById("signin");
var SignUpTab = document.getElementById("signup");
var ForgotTab = document.getElementById("forget");

var goToSignInF = document.getElementById("goToSignInf");
var goToSignUpF = document.getElementById("goToSignUpf");
var goToSignUp = document.getElementById("goToSignUp");
var goToSignIn = document.getElementById("goToSignIn");
var goToForget = document.getElementById("goToForget");

goToSignUpF.onclick = function(){ ToggleToSignUp(); }
goToSignUp.onclick = function(){ ToggleToSignUp(); }
goToSignInF.onclick= function(){ ToggleToSignIn(); }
goToSignIn.onclick = function(){ ToggleToSignIn(); }
goToForget.onclick = function(){ ToggleToForget(); }

window.onload = function () {
    // ToggleToSignIn()
    ToggleToSignIn()
}

function removeall() {
    SignInTab.classList.remove("block")
    SignUpTab.classList.remove("block")
    ForgotTab.classList.remove("block")

    SignInTab.classList.add("hidden")
    SignUpTab.classList.add("hidden")
    ForgotTab.classList.add("hidden")
}

function ToggleToSignIn() {
    console.log("Toogled to SignIn")
    removeall()
    SignInTab.classList.remove("hidden")
    SignInTab.classList.add("block")
};

function ToggleToSignUp() {
    console.log("Toggled to SignUp")
    removeall()
    SignUpTab.classList.remove("hidden")
    SignUpTab.classList.add("block")
};

function ToggleToForget() {
    console.log("Toggled to Forget")
    removeall()
    
    ForgotTab.classList.remove("hidden")
    ForgotTab.classList.add("block")
};