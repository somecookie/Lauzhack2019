file_content = null;

let sendFormToBackend = function() {

    let checked = document.querySelector('.custom-control-input').checked;

    console.log(file_content)

    let toSend = {
        namePatient : document.getElementById('patient').value,
        nameDoctor : document.getElementById('doctor').value,
        success : String(checked),
        toHash : file_content,
    };

    $.ajax({
        type: "POST",
        url: '/write',
        data: toSend,
        success: function(answer) {
            console.log(answer);
            alert('Done')
        }
    })
};


function fileSelectionHandler(e) {

    file = e.target.files[0]

    var reader = new FileReader();
     reader.readAsText(file,'UTF-8');

    reader.onload = readerEvent => {
        var content = readerEvent.target.result;
        file_content = content;
   }
}