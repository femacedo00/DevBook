$("#unfollow").on("click", unfollow);
$("#follow").on("click", follow);
$("#edit-user").on("submit", edit);

function unfollow() {
    const UserID = $(this).data("user-id");
    $(this).prop("disablrd", true);

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
        $(this).prop("disablrd", false);
    });
}

function follow() {
    const UserID = $(this).data("user-id");
    $(this).prop("disablrd", true);

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
        $(this).prop("disablrd", false);
    });
}

function edit(event) {
    event.preventDefault();

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
        $(this).prop("disablrd", false);
    });
}
