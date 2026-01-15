$("#new-publication").on("submit", createPublication);

function createPublication(event) {
    event.preventDefault();

    $.ajax({
        url: "/publications",
        method: "POST",
        data: {
            title: $('#title').val(),
            content: $('#content').val(),
        }
    }).done(function(){
        alert("Post saved successfully!");
        window.location = "/home";
    }).fail(function(){
        alert("Error saving post!");
    });
}