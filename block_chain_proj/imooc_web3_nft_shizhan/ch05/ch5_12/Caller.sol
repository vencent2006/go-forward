// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;
import "Callee.sol";

contract Caller {
    address calleeAddress;
    uint public xx;

    constructor(address _calleeAddress) {
        calleeAddress = _calleeAddress;
    }
    function setCalleeX(uint _x) public  {
        bytes memory cd = abi.encodeWithSignature("setX(uint256)", _x);
        (bool suc, bytes memory rst) = calleeAddress.call(cd);
        if(!suc) {
            // 终止交易进行
            revert("call failed");
        }
        (uint256 x) = abi.decode(rst, (uint256));
        xx = x;
    }
}