let blockClicked = null;
let storedBlocks = [];
let row = -1;

let getBlocks = function() {
    let get = $.ajax({
        type: "GET",
        url: '/get',
        dataType:'json',
        success : function(answer) {
        let str = "";
        let blocks = document.getElementById('listBlocks');

        if (answer != null) {
            for(let i = 0 ; i < answer.length ; i++) {
                let block = answer[i];
                storedBlocks.push(block);
                str += "<tr onclick=\"changePage(this);\" style=\"cursor: pointer;\">";
                str += "<th scope=\"row\">" + (i + 1) + "</th>";
                str += "<td>" + block.FirstName + "</td>";
                str += "<td>" + block.Name + "</td>";
                str += "<td>" + block.Location + "</td>";
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
        str += "<tr><th scope=\"row\">Name</th><td><i class=\"fas mr-2 grey-text\" aria-hidden=\"true\"></i>" + localStorage.getItem('FirstName') + "</td></tr>";
        str += "<tr><th scope=\"row\">Surname</th><td><i class=\"fas mr-2 grey-text\" aria-hidden=\"true\"></i>"+ localStorage.getItem('Name') +"</td></tr>";
        str += "<tr><th scope=\"row\">Location</th><td><i class=\"fas mr-2 grey-text\" aria-hidden=\"true\"></i>"+ localStorage.getItem('Location') +"</td></tr>";
        str += "<tr><th scope=\"row\">Success</th><td><i class=\"fas mr-2 grey-text\" aria-hidden=\"true\"></i>"+ localStorage.getItem('Success') +"</td></tr>";
        str += "<tr><th scope=\"row\">ReportHash</th><td><i class=\"fas mr-2 grey-text\" aria-hidden=\"true\"></i>"+ localStorage.getItem('ReportHash') +"</td></tr>";

    //}
    toDiplay.innerHTML = str;
};

let changePage = function(e) {
    index = e.rowIndex;
    localStorage.setItem('FirstName', storedBlocks[index - 1].FirstName);
    localStorage.setItem('Name', storedBlocks[index - 1].Name);
    localStorage.setItem('Location', storedBlocks[index - 1].Location);
    localStorage.setItem('Success', storedBlocks[index - 1].Success);
    localStorage.setItem('ReportHash', storedBlocks[index - 1].ReportHash);
    location.href='block.html';
};



