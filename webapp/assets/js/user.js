$("#unfollow").on("click", unfollow);
$("#follow").on("click", follow);
$("#edit-user").on("submit", edit);
$("#update-password").on("submit", updatePassword);

function unfollow() {
    const UserID = $(this).data("user-id");
    $(this).prop("disabled", true);

    $.ajax({
        url: `/users/${UserID}/unfollow`,
        method: "POST"
    }).done(function () {
        window.location = `/users/${UserID}`;
    }).fail(function () {
        swal.fire(
            "Ops...",
            "Failed to unfollow the user!",
            "error"
        );
        $(this).prop("disabled", false);
    });
}

function follow() {
    const UserID = $(this).data("user-id");
    $(this).prop("disabled", true);

    $.ajax({
        url: `/users/${UserID}/follow`,
        method: "POST"
    }).done(function () {
        window.location = `/users/${UserID}`;
    }).fail(function () {
        swal.fire(
            "Ops...",
            "Failed to follow the user!",
            "error"
        );
        $(this).prop("disabled", false);
    });
}

function edit(event) {
    event.preventDefault();
    $(this).prop("disabled", true);

    $.ajax({
        url: `/edit-user`,
        method: "PUT",
        data: {
            name: $("#name").val(),
            email: $("#email").val(),
            nick: $("#nick").val()
        }
    }).done(function () {
        swal.fire(
            "Success!",
            "User updated successfully!",
            "success"
        ).then(function () {
            window.location = `/profile`;
        });
    }).fail(function () {
        swal.fire(
            "Ops...",
            "Failed to update the user!",
            "error"
        );
        $(this).prop("disabled", false);
    });
}

function updatePassword(event) {
    event.preventDefault();
    $(this).prop("disabled", true);

    if($("#new").val() != $("#confirm").val()) {
        swal.fire(
            "Ops...",
            "Password not match!",
            "warning"
        );
        return;
    }

    $.ajax({
        url: `/update-password`,
        method: "POST",
        data: {
            current: $("#current").val(),
            new: $("#new").val()
        }
    }).done(function () {
        swal.fire(
            "Success!",
            "Password updated successfully!",
            "success"
        ).then(function () {
            window.location = `/profile`;
        });
    }).fail(function () {
        swal.fire(
            "Ops...",
            "Failed to update the password!",
            "error"
        );
        $(this).prop("disabled", false);
    });
}
