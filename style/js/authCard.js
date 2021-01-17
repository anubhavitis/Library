var SignInTab = document.getElementById("signin");
var SignUpTab = document.getElementById("signup");
var ForgotTab = document.getElementById("forget");

function removeall() {
    SignInTab.classList.remove("block")
    SignUpTab.classList.remove("block")
    ForgotTab.classList.remove("block")

    SignInTab.classList.add("hidden")
    SignUpTab.classList.add("hidden")
    ForgotTab.classList.add("hidden")
}

export function ToggleToSignIn() {
    console.log("Toogled to SignIn")
    removeall()
    SignInTab.classList.remove("hidden")
    SignInTab.classList.add("block")
};

export function ToggleToSignUp() {
    console.log("Toggled to SignUp")
    removeall()
    SignUpTab.classList.remove("hidden")
    SignUpTab.classList.add("block")
};

export function ToggleToForget() {
    console.log("Toggled to Forget")
    removeall()
    
    ForgotTab.classList.remove("hidden")
    ForgotTab.classList.add("block")
};