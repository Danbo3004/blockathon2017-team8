generate:
	curl -i -H "Accept: application/json" \
	-X POST -d '{"name":"Runi", "token_name":"RUNI", "owner":"1231", "number_of_token":100, "issuer_address":"0x42aa941997a340aa34c69f1be4f8772c5b33816f"}' \
	localhost:8080/v1/companies/register
	# http:/35.186.153.113:8080/v1/companies/register

list:
	curl localhost:8080/v1/contracts
