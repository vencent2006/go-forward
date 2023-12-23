// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract Gas {
    uint public i = 0;
    uint public remained;
 
    // Using up all of the gas that you send causes your transaction to fail.
    // State changes are undone.
    // Gas spent are not refunded.
    function forever() public {
        // Here we run a loop until all of the gas are spent
        // and the transaction fails
        while (true) {
            if(i > 100)
                return;
            if(i == 10)
                remained = gasleft();
            i += 1;
        }
    }

}


// contract GasCaller{
//     Gas gas;
//     constructor(Gas _gas) {
//         gas = _gas;
//     }
//     function callForever() public {
//          gas.forever{gas: 100000}();
//     }
// }

contract GasCaller{
    address gas;
    constructor(address _gas) {
        gas = _gas;
    }
    function callForever() public {
        bytes memory cd = abi.encodeWithSignature("forever()");
        (bool suc, bytes memory data) = gas.call{gas: 100000}(cd);
        if(!suc){
            revert("gas not enough");
        }
    }
}