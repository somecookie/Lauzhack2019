let loginFunc = function() {

    $.ajax({
        type: "GET",
        url: '/login',
        data: {
            "username" : document.getElementById('login').value,
            "password" : document.getElementById('password').value,
        },
        success: function(answer) {
            console.log(answer);
            if (answer == true) {
                location.href='search.html';
            }
        }
    })
};