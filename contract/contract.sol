pragma solidity ^0.5.13;

contract Doc{
    mapping(string => mapping(bytes32 => bool)) public validated;
    mapping(string => uint) public docMedOps;
    mapping(string => uint) public docSuccesses;

    function appendMedOp(string memory doc) internal{
        docMedOps[doc] += 1;
    }

    function appendDocSuccess(string memory doc, bool success) internal{
        if(success){
            docSuccesses[doc] += 1;
        }
    }

    function appendValidated(string memory doc, bytes32 hash, bool success) public {
        validated[doc][hash] = false;
        appendMedOp(doc);
        appendDocSuccess(doc, success);
    }

    function updateValidated(string memory doc, bytes32 hash) public{
        validated[doc][hash] = true;
    }

    function getDocMedOps(string memory doc) view public returns (uint){
        return docMedOps[doc];
    }

    function getDocSuccess(string memory doc) view public returns (uint){
        return docSuccesses[doc];
    }
}