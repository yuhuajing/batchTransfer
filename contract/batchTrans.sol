// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.18;

import "@openzeppelin/contracts/access/Ownable.sol";

contract Distribute is Ownable {
    receive() external payable {}

    function handOut(
        address[] calldata addresses,
        uint256[] calldata amounts
    ) public payable onlyOwner() {
        require(addresses.length > 0, "please enter the acceptance address");
        for (uint i = 0; i < addresses.length; i++) {
            payable(addresses[i]).transfer(amounts[i]*(10**14));
        }
    }

    function withdraw() public onlyOwner() {
        payable(msg.sender).transfer(address(this).balance);
    }

     function getBalance() external view onlyOwner() returns(uint256) {
        return address(this).balance;
    }

    function renounceOwnership() public virtual override onlyOwner {
        revert("renouncing ownership is disabled");
    }
}