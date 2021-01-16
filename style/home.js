SignInTab = document.getElementById("signin");
SignUpTab = document.getElementById("signup");
ForgotTab = document.getElementById("forget");


goToSignUp = document.getElementById("goToSignUp");
goToSignIn = document.getElementById("goToSignIn");
goToForget = document.getElementById("goToForget");

goToSignUp.onclick = function(){ ToggleToSignUp(); }
goToSignIn.onclick = function(){ ToggleToSignIn(); }
goToForget.onclick = function(){ ToggleToForget(); }

window.onload = function () {
    // ToggleToSignIn()
    ToggleToSignUp()
}

function removeall() {
    console.log("Welcome to removeall")
    SignInTab.classList.remove("block")
    SignUpTab.classList.remove("block")
    ForgotTab.classList.remove("block")

    SignInTab.classList.add("hidden")
    SignUpTab.classList.add("hidden")
    ForgotTab.classList.add("hidden")
}

function ToggleToSignIn() {
    removeall()

    SignInTab.classList.remove("hidden")
    SignInTab.classList.add("block")

    console.log(SignInTab.classList)
};

function ToggleToSignUp() {
    removeall()
    console.log(SignUpTab)
    SignUpTab.classList.remove("hidden")
    SignUpTab.classList.add("block")

    console.log(SignInTab.classList)
    console.log(SignUpTab.classList)
    console.log(ForgotTab.classList)
};

function ToggleToForget() {
    removeall()
    
    ForgotTab.classList.remove("hidden")
    ForgotTab.classList.add("block")

    console.log(SignInTab.classList)
    console.log(SignUpTab.classList)
    console.log(ForgotTab.classList)
};