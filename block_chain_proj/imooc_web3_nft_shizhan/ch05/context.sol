// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

contract Callee {
    uint public x;
    address public caller; // 调用 callee 的 调用者，其实就是 caller
    address public eoaaddress;
    function setX(uint _x) public {
        caller = msg.sender;
        eoaaddress = tx.origin;
        x = _x;
    }
}

contract Caller {
    address calleeAddress;
    address public caller; // 调用 callee 的 调用者，其实就是 EOA
     address public eoaaddress;
    constructor(address _calleeAddress) {
        calleeAddress = _calleeAddress;
    }
    function setCalleeX(uint _x) public  {
        caller = msg.sender;
        eoaaddress = tx.origin;
        Callee callee = Callee(calleeAddress);
        callee.setX(_x);
    }
}

// callee 0xD7ACd2a9FD159E69Bb102A1ca21C9a3e3A5F771B

// caller 0x7EF2e0048f5bAeDe046f6BF797943daF4ED8CB47

// remix eoa 0x5B38Da6a701c568545dCfcB03FcB875f56beddC4