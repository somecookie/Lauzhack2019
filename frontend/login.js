let loginFunc = function() {

    /*let toSend = {
        username : document.getElementById('login').value,
        password : document.getElementById('password').value,
    };*/

    //console.log(toSend);

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