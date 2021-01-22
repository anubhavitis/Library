document.getElementById("forgot").onsubmit = function (e) {
  forgot(e);
};

function forgot(e){
    e.preventDefault();
    swal.fire(
        "Feature Unavailabe!",
        "We are working on this.",
        "error"
    );
}