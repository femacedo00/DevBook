$("#unfollow").on("click", unfollow);
$("#follow").on("click", follow);

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
