export GOPROXY=https://goproxy.io
export DB_Password=8280508
export DB_User=ubuntu

go build main.go webassembly

./webassembly

# nohup ./run.sh > run.log 2>&1 &