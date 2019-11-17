let blockClicked = null;
let storedBlocks = [];
let row = -1;

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
                str += "<tr onclick=\"changePage(this);\" style=\"cursor: pointer;\">";
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

let changePage = function(e) {
    index = e.rowIndex;
    localStorage.setItem('gas', storedBlocks[index - 1].gas);
    localStorage.setItem('gasPrice', storedBlocks[index - 1].gasPrice);
    localStorage.setItem('hash', storedBlocks[index - 1].hash);
    localStorage.setItem('input', storedBlocks[index - 1].input);
    localStorage.setItem('nonce', storedBlocks[index - 1].nonce);
    localStorage.setItem('to', storedBlocks[index - 1].to);
    location.href='block.html';
};



