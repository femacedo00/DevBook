$("#form-register").on('submit', registerUser);

function registerUser(event) {
    event.preventDefault();

    if ($("#password").val() != $("#confirm-password").val()) {
        swal.fire(
            "Ops...",
            "Password not match!",
            "error"
        );
        return;
    }

    $.ajax({
        url: "/users",
        method: "POST",
        data: {
            name: $("#name").val(),
            email: $("#email").val(),
            nick: $("#nick").val(),
            password: $("#password").val()
        }
    }).done(function () {
        swal.fire(
            "Success!",
            "User Successufully Registered!",
            "success"
        ).then(function () {
            $.ajax({
                url: "/login",
                method: "POST",
                data: {
                    email: $("#email").val(),
                    password: $("#password").val()
                }
            }).done(function () {
                window.location = "/home";
            }).fail(function () {
                swal.fire(
                    "Ops...",
                    "User Authentication Failed!",
                    "error"
                );
            });
        });
    }).fail(function () {
        swal.fire(
            "Ops...",
            "User Registration Failed!",
            "error"
        );
    });
}