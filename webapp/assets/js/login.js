$("#form-login").on('submit', loginUser);

function loginUser(event) {
    event.preventDefault();
    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $("#email").val(),
            password: $("#password").val()
        }
    }).done(function () {
        window.location.href = "/home";
    }).fail(function (error) {
        swal.fire(
            "Ops...",
            "User Authentication Failed!",
            "error"
        );
    });
}