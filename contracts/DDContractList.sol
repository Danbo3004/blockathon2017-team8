pragma solidity ^0.4.17;

contract DDContractList {

  address[] ListOfContracts;

  function push(address _address) {
    ListOfContracts.push(_address);
  }

  function getLength() returns (uint) {
    return ListOfContracts.length;
  }

  function getContracList() returns (address[]) {
    return ListOfContracts;
  }

  function deleteByIndex(uint _deleteIndex) {
    require(_deleteIndex < ListOfContracts.length);
    ListOfContracts[_deleteIndex] = 0;
  }

}
