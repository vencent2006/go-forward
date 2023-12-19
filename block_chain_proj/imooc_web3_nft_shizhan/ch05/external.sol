// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

contract ExternalDemo {
    address public caller;
    function first() public {
        this.second(); // 这个的调用方就是 ExternalDemo 自己，只要是this，就是合约自己，public也可以用this，但是是多余的
    }

    function second() external {
        caller = msg.sender;
    }
}

// eoa  0x5B38Da6a701c568545dCfcB03FcB875f56beddC4
// ExternalDemo 合约自己 0xd2a5bC10698FD955D1Fe6cb468a17809A08fd005
