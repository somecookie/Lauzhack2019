pragma solidity ^0.5.13;

contract Doc{
    mapping(string => mapping(bytes32 => bool)) public validated;
    mapping(string => uint) public docMedOps;
    mapping(string => uint) public docSuccesses;

    function appendMedOp(string doc){
        docMedOps[doc] += 1;
    }

    function appendDocSuccess(string doc, bool success){
        if(success){
            docSuccesses[doc] += 1;
        }
    }

    function appendValidated(string doc, bytes32 hash) public {
        validated[doc][hash] = true;
        appendMedOp(doc);
        appendDocSuccess(doc, success);
    }
}