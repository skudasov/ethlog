pragma solidity 0.8.6;
pragma experimental ABIEncoderV2;


contract AnotherEventsTestContract {
    uint storedData;

    // event used with input/output data
    event ValueEvent(address indexed sender, uint value);


    struct Account {
        string name;
        uint64 balance;
        uint dailylimit;
    }
    // events
    // 1 index by default is selector signature

    // basic types events
    event NoIndexEvent(address sender);
    event OneIndexEvent(string indexed a);
    event TwoIndexEvent(uint256 indexed roundId, address indexed startedBy);
    event ThreeIndexEvent(uint256 indexed roundId, address indexed startedBy, uint256 startedAt);

    // structs
    event NoIndexStructEvent(Account a);

    constructor() {}

    function set(uint x) public returns (uint) {
        storedData = x;
        emit ValueEvent(msg.sender, x);
        x -= 5;
        return x;
    }

    function get() public view returns (uint) {
        return storedData - 10;
    }

    function fireNoIndexEvent() public {
        emit NoIndexEvent(msg.sender);
    }

    function fireOneIndexEvent() public {
        emit OneIndexEvent("event_data_string");
    }

    function fireTwoIndexEvent() public {
        emit TwoIndexEvent(1, address(msg.sender));
    }

    function fireThreeIndexEvent() public {
        emit ThreeIndexEvent(1, address(msg.sender), 3);
    }

    function fireNoIndexStructEvent() public {
        emit NoIndexStructEvent(Account("John", 5, 10));
    }
}