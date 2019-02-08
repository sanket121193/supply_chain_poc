echo "POST invoke chaincode on peers of Org1, Org2, Org3 and Org4"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/mycc \
  -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDk1NTkxMzgsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Ik9yZzEiLCJpYXQiOjE1NDk1MjMxMzh9.D6HXwPZKPROxTn6tqBVLA-VC9BvlWJMujvFq2pb6VU8" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.org1.rolex.com","peer0.org2.dealer.com", "peer0.org3.service.com"],
	"fcn":	"addProduct",
	"args":	["N1122","B1115","M4434","Rolex Shine","06-09-2018","C87565", "$1580", "Shine Finish"]
}')
echo "Transaction ID is $TRX_ID"
echo
echo


