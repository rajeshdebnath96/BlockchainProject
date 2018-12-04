docker rm -f $(docker ps -aq)

docker-compose -f docker-compose-simple.yaml up

docker exec -it chaincode bash
cd mychaincode
go build
CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=mycc:0 ./mychaincode

docker exec -it cli bash
peer chaincode install -p chaincodedev/chaincode/mychaincode -n mycc -v 0
peer chaincode instantiate -n mycc -v 0 -c '{"Args":["a","10"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["submit", "a", "b","2000","bank","654354546"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["submit", "b", "c","400","bKash","7ttu546"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["submit", "c", "a","67000","Rocket","7u5utuybtg"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["view","900788600000"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["verify","900788600000","false"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["mytransactions","a"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["alltransactions"]}' -C myc

