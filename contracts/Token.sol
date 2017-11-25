pragma solidity ^0.4.17;

contract Token {

  address public TokenIssuerAddress;
  bytes32 public TokenIssuerName;

  struct TokenShop {
    bool isActived;
    uint amountTokens;
    uint tokenPrice;
  }

  mapping (address => TokenShop) public tokenShops;
  mapping (address => uint) totalTokens;

  event TokenForSale(address _from, uint _amountTokens, uint _tokenPrices);
  event BuyToken(address _seller, address _buyer, uint _amountTokens);

  function Token(address _tokenIssuer, bytes32 _tokenIssuerName, uint _amountTokens) public {
    TokenIssuerAddress = _tokenIssuer;
    totalTokens[TokenIssuerAddress] = _amountTokens;
    TokenIssuerName = _tokenIssuerName;
  }


  function sellTokens(uint _amountTokens, uint _tokenPrice) public {

    require(totalTokens[msg.sender] > _amountTokens);
    require(_amountTokens > 0);
    require(_tokenPrice > 0);

    tokenShops[msg.sender].amountTokens = _amountTokens;
    tokenShops[msg.sender].tokenPrice = _tokenPrice;
    tokenShops[msg.sender].isActived = true;
    TokenForSale(msg.sender, _amountTokens, _tokenPrice);

  }

  function byTokens(address _from) payable public {

    require(tokenShops[_from].isActived == true);
    uint tokenPrice = tokenShops[_from].tokenPrice;
    uint totalTokenBuy = msg.value / tokenPrice;

    require(tokenShops[_from].amountTokens > totalTokenBuy);

    tokenShops[_from].amountTokens -= totalTokenBuy;
    totalTokens[_from] -= totalTokenBuy;
    totalTokens[msg.sender] += totalTokenBuy;
    _from.transfer(msg.value);
    BuyToken(_from, msg.sender, totalTokenBuy);

  }

}
