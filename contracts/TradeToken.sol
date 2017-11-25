pragma solidity ^0.4.17;
import "./Token.sol";

contract TradeToken {

  Token public Tok;

  address TokenAType;
  address TokenBType;
  uint AmountA;
  uint AmountB;
  uint ConvertionRate;

  address A_Actor;
  address B_Actor;

  event exChangeToken(address _A_Actor, bytes32 _tokenNameA, uint _amountA, address _B_Actor, bytes32 _tokenNameB, uint _amountB);

  function TradeToken(address _tokenAType, uint _amountA, address _tokenBType, uint _amountB) public {
      require(getTotalAmountToken(_tokenAType, msg.sender) > _amountA);
      A_Actor = msg.sender;
      TokenAType = _tokenAType;
      TokenBType = _tokenBType;
      AmountA = _amountA;
      AmountB = _amountB;
  }

  function someOneAccept() {
    require(AmountA > 0);
    require(AmountB > 0);
    B_Actor = msg.sender;

    uint totalTokenB_ofA = getTotalAmountToken(TokenBType, A_Actor);
    uint totalTokenA_ofB = getTotalAmountToken(TokenAType, B_Actor);


    uint totalTokenA_ofA = getTotalAmountToken(TokenAType, A_Actor);
    uint totalTokenB_ofB = getTotalAmountToken(TokenBType, B_Actor);

    require(totalTokenA_ofA >= AmountA);
    require(totalTokenB_ofB >= AmountB);

    totalTokenA_ofA -= AmountA;
    totalTokenB_ofB -= AmountB;

    totalTokenB_ofA += AmountB;
    totalTokenA_ofB += AmountA;

    AmountA = 0;
    AmountB = 0;

    setTotalAmountToken(TokenAType, B_Actor, totalTokenA_ofB);
    setTotalAmountToken(TokenBType, B_Actor, totalTokenB_ofB);

    setTotalAmountToken(TokenAType, A_Actor, totalTokenA_ofA);
    setTotalAmountToken(TokenBType, A_Actor, totalTokenB_ofA);

    exChangeToken(A_Actor, getThisTokenName(TokenAType), AmountA, B_Actor, getThisTokenName(TokenBType), AmountB);
  }

  function getThisTokenName(address _tokenAddress) private returns (bytes32) {
    Tok = Token(_tokenAddress);
    return  Tok.getNameThisToken();
  }

  function getTotalAmountToken(address _tokenAddress, address _someone) private returns (uint) {
    Tok = Token(_tokenAddress);
    return Tok.getAmountThisToken(_someone);
  }

  function setTotalAmountToken(address _tokenAddress, address _somone, uint _totalAmount) private returns (bool) {
    Tok = Token(_tokenAddress);
    if(Tok.setAmoutThisToken(_somone, _totalAmount))
      return true;
    else
      return false;
  }

}
