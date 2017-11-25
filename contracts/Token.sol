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

  function setAmoutThisToken(address _someone, uint _newAmount) returns (bool) {
    // 1. Check iff address of msg.sender valid
    // 2. Set new amount Token
    totalTokens[_someone] = _newAmount;
    return true;
  }





}
