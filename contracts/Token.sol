pragma solidity ^0.4.17;

contract Token {

  address public TokenIssuerAddress;
  bytes32 public TokenName;


  struct TokenShop {
    bool isActived;
    uint amountTokens;
    uint tokenPrice;
  }

  mapping (address => TokenShop) public tokenShops;
  mapping (address => uint) totalTokens;

  event SellToken_Ether(address _from, uint _amount, uint _tokenPrice, address _tokenIssuerAddress, bytes32 _tokenName);
  event BuyToken_Ehter(address _from, address _to,uint _totalTokenBuy, address _tokenIssuerAddress, bytes32 _tokenName);

  function Token(address _tokenIssuer, bytes32 _tokenName, uint _amountTokens) public {
    TokenIssuerAddress = _tokenIssuer;
    totalTokens[TokenIssuerAddress] = _amountTokens;
    TokenName = _tokenName;
  }


  function sellTokens(uint _amountTokens, uint _tokenPrice) public {

    require(totalTokens[msg.sender] > 0);
    require(totalTokens[msg.sender] > _amountTokens);
    require(_amountTokens > 0);
    require(_tokenPrice > 0);

    tokenShops[msg.sender].amountTokens = _amountTokens;
    tokenShops[msg.sender].tokenPrice = _tokenPrice;
    tokenShops[msg.sender].isActived = true;
    SellToken_Ether(msg.sender, _amountTokens, _tokenPrice, TokenIssuerAddress, TokenName);

  }

  function buyTokens(address _from) payable public {

    require(tokenShops[_from].isActived == true);
    uint tokenPrice = tokenShops[_from].tokenPrice;
    uint totalTokenBuy = msg.value / tokenPrice;

    require(tokenShops[_from].amountTokens > totalTokenBuy);

    tokenShops[_from].amountTokens -= totalTokenBuy;
    totalTokens[_from] -= totalTokenBuy;
    totalTokens[msg.sender] += totalTokenBuy;
    _from.transfer(msg.value);
    BuyToken_Ehter(_from, msg.sender, totalTokenBuy, TokenIssuerAddress, TokenName);

  }

  function getNameThisToken() view public returns (bytes32){
    // 1. Check if address of msg.sender valid
    // 2. return name of this Token
    return TokenName;
  }

  function getAmountThisToken(address _someone) view public returns (uint){
    // 1. Check if address of msg.sender valid
    // 2. Reutrn sumary of this Token
    return totalTokens[_someone];
  }

  function getIssuerAddress() public returns (address){
    return TokenIssuerAddress;
  }

  function setAmoutThisToken(address _someone, uint _newAmount) returns (bool) {
    // 1. Check iff address of msg.sender valid
    // 2. Set new amount Token
    totalTokens[_someone] = _newAmount;
    return true;
  }
}

// -------------------------------------------------------------------------
contract TradeToken {

  Token public Tok;

  address TokenAType;
  address TokenBType;
  uint AmountA;
  uint AmountB;
  uint ConvertionRate;

  address A_Actor;
  address B_Actor;

  event exChangeToken(address _A_Actor, address _tokenAddress, uint _amountA, address _B_Actor, address _tokenBAddress, uint _amountB);

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

    exChangeToken(A_Actor, TokenAType, AmountA, B_Actor, TokenBType, AmountB);

    // TODO: Give trading fee to platform owner (blocker)
  }

  function getPairToken() returns (address[]) {
    address[] pairAddress;
    pairAddress.push(TokenAType);
    pairAddress.push(TokenBType);
    return pairAddress;
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

// -------------------------------------------------------------------------
contract DDContractList {

  address[] ListOfContracts;

  function push(address _address) {
    ListOfContracts.push(_address);
  }

  function getLength() constant returns (uint) {
    return ListOfContracts.length;
  }

  function getContracList() constant returns (address[]) {
    return ListOfContracts;
  }

  function deleteByIndex(uint _deleteIndex) {
    require(_deleteIndex < ListOfContracts.length);
    ListOfContracts[_deleteIndex] = 0;
  }

}
