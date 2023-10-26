// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";

contract SBT is ERC1155 {
    address public owner;

    constructor() ERC1155("") payable {
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only operator can mint new souls");
        _;
    }

    struct TokenInfo {
        uint256 tokenid;
        uint256 minted;
        uint256 totalamount;
        string name;
        string symbol;
        string url;
    }

    mapping(uint256 => TokenInfo) tokenIDInfo; // tokenid

    function settokenIDInfo(
        uint256 tokenid,
        uint256 totalamount,
        string memory name,
        string memory symbol,
        string memory url //https://ipfs.io/ipfs/QmW948aN4Tjh4eLkAAo8os1AcM2FJjA46qtaEfFAnyNYzY
    ) external onlyOwner {
        require(
            tokenIDInfo[tokenid].totalamount == 0,
            "TokenID already Initialized"
        );
        TokenInfo memory tokeninfo = TokenInfo({
            tokenid: tokenid,
            minted: 0,
            totalamount: totalamount,
            name: name,
            symbol: symbol,
            url: url
        });
        tokenIDInfo[tokenid] = tokeninfo;
    }

    function mint(
        address account,
        uint256 id,
        uint256 amount
    ) public onlyOwner {
        require(tokenIDInfo[id].totalamount != 0, "TokenID not Initialized");
        require(
            tokenIDInfo[id].minted + amount <= tokenIDInfo[id].totalamount,
            "Not Enough TokenID left"
        );
        tokenIDInfo[id].minted += amount;
        _mint(account, id, amount, bytes(""));
    }

    function batchmint(
        address[] memory account,
        uint256[] memory id,
        uint256[] memory amount
    ) external onlyOwner {
        require(
            account.length != 0 && account.length == amount.length,
            "Need valid account/amount list and Equal length"
        );
        if (id.length == 1) {
            for (uint256 i = 0; i < account.length; i++) {
                mint(account[i], id[0], amount[i]);
            }
        } else {
            require(
                account.length == id.length,
                "Need valid account/amount/id list and Equal length"
            );
            for (uint256 i = 0; i < account.length; i++) {
                mint(account[i], id[i], amount[i]);
            }
        }
    }

    function burn(uint256 id, uint256 value) external virtual {
        tokenIDInfo[id].minted -= value;
        _burn(msg.sender, id, value);
    }

    function updateOwner(address newowner) external onlyOwner {
        require(newowner != address(0), "Invalid Owner");
        owner = newowner;
    }

    function safeTransferFrom(
        address,
        address,
        uint256,
        uint256,
        bytes memory
    ) public pure virtual override {
        revert("Transfer not supported for soul bound token.");
    }

    function safeBatchTransferFrom(
        address,
        address,
        uint256[] memory,
        uint256[] memory,
        bytes memory
    ) public virtual override {
        revert("Transfer not supported for soul bound token.");
    }

    function setApprovalForAll(address, bool) public virtual override {
        revert("Transfer not supported for soul bound token.");
    }

    function isApprovedForAll(address, address)
        public
        view
        virtual
        override
        returns (bool)
    {
        revert("Transfer not supported for soul bound token.");
    }

    function updateTotalAmount(uint256 tokenid, uint256 totalamount)
        external
        onlyOwner
    {
        require(
            tokenIDInfo[tokenid].totalamount != 0,
            "TokenID not Initialized"
        );
        tokenIDInfo[tokenid].totalamount = totalamount;
    }

    function updateURL(uint256 tokenid, string memory url) external onlyOwner {
        require(
            tokenIDInfo[tokenid].totalamount != 0,
            "TokenID not Initialized"
        );
        tokenIDInfo[tokenid].url = url;
    }

    function updateNameSymbol(
        uint256 tokenid,
        string memory name,
        string memory symbol
    ) external onlyOwner {
        require(
            tokenIDInfo[tokenid].totalamount != 0,
            "TokenID not Initialized"
        );
        tokenIDInfo[tokenid].name = name;
        tokenIDInfo[tokenid].symbol = symbol;
    }

    function uri(uint256 id)
        public
        view
        virtual
        override
        returns (string memory)
    {
        return tokenIDInfo[id].url;
    }

    function getname(uint256 tokenid) external view returns (string memory) {
        return tokenIDInfo[tokenid].name;
    }

    function getsymbol(uint256 tokenid) external view returns (string memory) {
        return tokenIDInfo[tokenid].symbol;
    }

    function gettotalamount(uint256 tokenid) external view returns (uint256) {
        return tokenIDInfo[tokenid].totalamount;
    }

    function getminted(uint256 tokenid) external view returns (uint256) {
        return tokenIDInfo[tokenid].minted;
    }
}
