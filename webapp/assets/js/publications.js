$("#new-publication").on("submit", createPublication);
$("#update-publication").on("click", updatePublication);
$(".delete-publication").on("click", deletePublication);

$(document).on("click", ".like-publication", likePublication);
$(document).on("click", ".dislike-publication", dislikePublication);

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
        alert("Post successfully saved!");
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

        likeElement.addClass("dislike-publication")
        likeElement.addClass("text-danger")
        likeElement.removeClass("like-publication")
    }).fail(function () {
        alert("Error liking the post!")
    }).always(function () {
        likeElement.prop("disabled", false);
    });
}

function dislikePublication(event) {
    event.preventDefault();

    const likeElement = $(event.target);
    const publicationID = likeElement.closest("div.bg-body-tertiary").data("publication-id");

    likeElement.prop("disabled", true)
    $.ajax({
        url: `/publications/${publicationID}/dislike`,
        method: "POST"
    }).done(function () {
        const countLikes = likeElement.next("span");
        const amountLikes = parseInt(countLikes.text());

        countLikes.text(amountLikes - 1);

        likeElement.removeClass("dislike-publication")
        likeElement.removeClass("text-danger")
        likeElement.addClass("like-publication")
    }).fail(function () {
        alert("Error disliking the post!")
    }).always(function () {
        likeElement.prop("disabled", false);
    });
}

function updatePublication() {
    const publication = this;
    const publicationID =  $(publication).data("publication-id");
    
    $(publication).prop("disabled", true);

    $.ajax({
        url: `/publications/${publicationID}`,
        method: "PUT",
        data: {
            title: $("#title").val(),
            content: $("#content").val(),
        }
    }).done(function(){
        alert("Post successfully updated")
    }).fail(function(){
        alert("Error updating the post!")
    }).always(function(){
        $(publication).prop("disabled", false);
    });
}

function deletePublication(event) {
    event.preventDefault();

    const deleteElement = $(event.target);
    const publication = deleteElement.closest("div.bg-body-tertiary");
    const publicationID = publication.data("publication-id");

    deleteElement.prop("disabled", true);

    $.ajax({
        url: `/publications/${publicationID}`,
        method: "DELETE"
    }).done(function() {
        publication.fadeOut("slow", function() {
            $(this).remove();
        })
    }).fail(function() {
        alert("Error deleting post");
    });
}
