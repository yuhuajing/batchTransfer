// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.20;

contract Distribute {
    address public owner;

    constructor() payable {
        owner = msg.sender;
    }

    receive() external payable {}

    modifier onlyOwner() {
        require(msg.sender == owner, "Only operator can mint new souls");
        _;
    }

    function updateOwner(address newowner) external onlyOwner {
        require(newowner != address(0), "Invalid Owner");
        owner = newowner;
    }

    function handOut(address[] calldata addresses, uint256[] calldata amounts)
        public
        payable
        onlyOwner
    {
        require(addresses.length != 0, "please enter the acceptance address");
        require(amounts.length == addresses.length, "Unmatched length");
        for (uint256 i = 0; i < addresses.length; i++) {
            payable(addresses[i]).transfer(amounts[i] * (10**14));
        }
    }

    function withdraw() public onlyOwner {
        payable(msg.sender).transfer(address(this).balance);
    }

    function getBalance() external view returns (uint256) {
        return address(this).balance;
    }
}
