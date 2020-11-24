cd services/models

protoc --micro_out=../ --go_out=../ *.proto

cd .. && cd ..