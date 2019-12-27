$(document).ready(function () {
    $("#signIn").hide();
    $("#goToSignInSection").click(function () {
        $("#signInSection").toggleClass("signin");
        $("#login, #signIn").toggle();
    });
});