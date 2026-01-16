$("#new-publication").on("submit", createPublication);
$(".like-publication").on("click", likePublication);

function createPublication(event) {
    event.preventDefault();

    $.ajax({
        url: "/publications",
        method: "POST",
        data: {
            title: $('#title').val(),
            content: $('#content').val(),
        }
    }).done(function () {
        alert("Post saved successfully!");
        window.location = "/home";
    }).fail(function () {
        alert("Error saving post!");
    });
}

function likePublication(event) {
    event.preventDefault();

    const likeElement = $(event.target);
    const publicationID = likeElement.closest("div.bg-body-tertiary").data("publication-id");

    likeElement.prop("disabled", true)
    $.ajax({
        url: `/publications/${publicationID}/like`,
        method: "POST"
    }).done(function () {
        const countLikes = likeElement.next("span");
        const amountLikes = parseInt(countLikes.text());

        countLikes.text(amountLikes + 1);
    }).fail(function () {
        alert("Error liking the post!")
    }).always(function () {
        likeElement.prop("disabled", false);
    });
}