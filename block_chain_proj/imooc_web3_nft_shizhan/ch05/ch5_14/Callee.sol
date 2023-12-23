// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

contract Callee {
    uint public x;
    function setX(uint _x) public returns (uint256) {
        x = _x;
        return x;
    }
    fallback() external  {
        x = 1234567;
    }
}