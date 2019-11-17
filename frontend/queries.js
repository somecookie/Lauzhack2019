let blockClicked = null;
let storedBlocks = [];
let storedUsers = [];
let row = -1;

let getDoctors = function() {
    $.ajax({
        type: "GET",
        url: '/doctors',
        dataType:'json',
        success : function(answer) {
            console.log(answer);
            let str = "";
            let docs = document.getElementById('listDocs');

            if (answer != null) {
                for(let i = 0 ; i < answer.length ; i++) {
                    let block = answer[i];
                    storedUsers.push(block);
                    str += "<tr onclick=\"changePageDoc(this);\" style=\"cursor: pointer;\" >";
                    str += "<th scope=\"row\">" + (i + 1) + "</th>";
                    str += "<td>" + block.name+ "</td>";
                    str += "<td>" + block.OpNb + "</td>";
                    str += "<td>" + block.SuccessNb + "</td>";
                    str += "</tr>";
                }
            }
            docs.innerHTML = str;
        }
    })
};

let displayDoctors = function() {
    str = "";
    let toDiplay = document.getElementById('docs');
    //console.log(rowBlock)
    //if (rowBlock != -1) {
    //curBlock = storedBlocks[rowBlock - 1];
    str += "<tr><th scope=\"row\">Name</th><td><i class=\"fas mr-2 grey-text\" aria-hidden=\"true\"></i>" + localStorage.getItem('Name') + "</td></tr>";
    str += "<tr><th scope=\"row\">Hash of the report</th><td><i class=\"fas mr-2 grey-text\" aria-hidden=\"true\"></i>"+ localStorage.getItem('hash') +"</td></tr>";
    str += "<tr><th scope=\"row\">Successful</th><td><i class=\"fas mr-2 grey-text\" aria-hidden=\"true\"></i>"+ localStorage.getItem('success') +"</td></tr>";
    str += "<tr><th scope=\"row\">Valid</th><td><i class=\"fas mr-2 grey-text\" aria-hidden=\"true\"></i>"+ localStorage.getItem('valid') +"</td></tr>";

    //}
    toDiplay.innerHTML = str;
};

let getBlocks = function() {
    $.ajax({
        type: "GET",
        url: '/get',
        dataType:'json',
        success : function(answer) {
            console.log(answer)
        let str = "";
        let blocks = document.getElementById('listBlocks');

        if (answer != null) {
            for(let i = 0 ; i < answer.length ; i++) {
                let block = answer[i];
                storedBlocks.push(block);
                str += "<tr onclick=\"changePageBlock(this);\" style=\"cursor: pointer;\">";
                str += "<th scope=\"row\">" + (i + 1) + "</th>";
                str += "<td>" + block.hash + "</td>";
                str += "<td>" + block.gasPrice + "</td>";
                str += "</tr>";
            }
        }
        blocks.innerHTML = str;
        }
    })
};

let displayBlock = function() {
    str = "";
    let toDiplay = document.getElementById('block');
    var curBlock = localStorage.getItem('block');
    //console.log(rowBlock)
    //if (rowBlock != -1) {
        //curBlock = storedBlocks[rowBlock - 1];
        str += "<tr><th scope=\"row\">Gas</th><td><i class=\"fas mr-2 grey-text\" aria-hidden=\"true\"></i>" + localStorage.getItem('gas') + "</td></tr>";
        str += "<tr><th scope=\"row\">Gas price</th><td><i class=\"fas mr-2 grey-text\" aria-hidden=\"true\"></i>"+ localStorage.getItem('gasPrice') +"</td></tr>";
        str += "<tr><th scope=\"row\">Hash</th><td><i class=\"fas mr-2 grey-text\" aria-hidden=\"true\"></i>"+ localStorage.getItem('hash') +"</td></tr>";
        str += "<tr><th scope=\"row\">Input</th><td><i class=\"fas mr-2 grey-text\" aria-hidden=\"true\"></i>"+ localStorage.getItem('input') +"</td></tr>";
        str += "<tr><th scope=\"row\">Nonce</th><td><i class=\"fas mr-2 grey-text\" aria-hidden=\"true\"></i>"+ localStorage.getItem('nonce') +"</td></tr>";
        str += "<tr><th scope=\"row\">To</th><td><i class=\"fas mr-2 grey-text\" aria-hidden=\"true\"></i>"+ localStorage.getItem('to') +"</td></tr>";

    //}
    toDiplay.innerHTML = str;
};

let changePageBlock = function(e) {
    index = e.rowIndex;
    localStorage.setItem('gas', storedBlocks[index - 1].gas);
    localStorage.setItem('gasPrice', storedBlocks[index - 1].gasPrice);
    localStorage.setItem('hash', storedBlocks[index - 1].hash);
    localStorage.setItem('input', storedBlocks[index - 1].input);
    localStorage.setItem('nonce', storedBlocks[index - 1].nonce);
    localStorage.setItem('to', storedBlocks[index - 1].to);
    location.href='block.html';
};

let changePageDoc = function(e) {
    index = e.rowIndex;
    localStorage.setItem('Name', storedUsers[index - 1].name);
    localStorage.setItem('OpNb', storedUsers[index - 1].OpNb);
    localStorage.setItem('SuccessNb', storedUsers[index - 1].SuccessNb);
    localStorage.setItem('hash', storedUsers[index - 1].ops[0].hash);
    localStorage.setItem('success', storedUsers[index - 1].ops[0].success);
    localStorage.setItem('valid', storedUsers[index - 1].ops[0].isValid);
    location.href='host.html';
};


function fileSelectionHandler(e) {

    file = e.target.files[0]

    var reader = new FileReader();
    reader.readAsText(file,'UTF-8');

    reader.onload = readerEvent => {
        var content = readerEvent.target.result;
        console.log(content)
        
        let hashFile = sjcl.hash.sha256.hash(content);

        let toSend = {
            hash : sjcl.codec.hex.fromBits(hashFile)
        };

        console.log("hello",toSend)


        $.ajax({
            type: "POST",
            url: '/validate',
            data: toSend,
            success: function(answer) {
                console.log(answer);
            }
        })
   }
}