// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract ReceiveEther {
    string public caller;
 
    receive() external payable {
    //    caller = "receive"; // 会超过 _to.send的gas 2300的限制，所以会失败
    }

    fallback() external  {
        caller = "fallback";
    }
    function deposit()public {

    }

    function getBalance() public view returns (uint) {
        return address(this).balance;
    }
}

contract SendEther {
    function sendViaTransfer(address payable _to) public payable {
        _to.transfer(msg.value);
    }

    function sendViaSend(address payable _to) public payable {

        bool sent = _to.send(msg.value);
        require(sent, "Failed to send Ether");
    }

    function sendViaCall(address  _to) public payable {
        
        (bool sent, bytes memory data) = _to.call{value: msg.value}("");
        require(sent, "Failed to send Ether");
    }
    uint public received;
    function deposit()public payable{
        received += msg.value;

    }
     function sendTo(address payable _to,uint amount) public  {
        
        (bool sent, bytes memory data) = _to.call{value: amount}("");
        require(sent, "Failed to send Ether");
    }

    function mysend(address to, uint amount)public returns(bool){
        (bool sent, bytes memory data) = to.call{gas: 2300, value: amount}("");
        return  sent;
    }
  function mytransfer(address to, uint amount)public returns(bool){
        (bool sent, bytes memory data) = to.call{gas: 2300, value: amount}("");
        require(sent, "Failed to send Ether");
    }

}