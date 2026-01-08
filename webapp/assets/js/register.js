$("#form-register").on('submit', registerUser);

function registerUser(event) {
    event.preventDefault();

    if ($("#password").val() != $("#confirm-password").val()) {
        alert("Password not match!");
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
    }).done(function() {
        alert("User Successufully Registered!");
    }).fail(function(error) {
        console.log(error);
        alert("User Registration Failed!")
    });
}