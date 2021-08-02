// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Journal {
    string public entry;
    uint public timestamp;
    address public author;

    function newEntry(string memory _entry) public {
        entry = _entry;
        timestamp = block.timestamp;
        author = msg.sender;
    }
}
