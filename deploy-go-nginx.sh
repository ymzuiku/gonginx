url=root@120.79.150.132
sshkey=~/.ssh/id_rsa
client=temp
server=go-nginx
remotePath=/db/go-server
deployFile=./deploy-go-nginx.sh

mkdir -p build || echo "have build dir"


function deployServer(){
  GOOS=linux GOARCH=amd64 go build -o ./build/$server main.go
  copyFile $deployFile $remotePath
  copyFile ./build/$server $remotePath
  remoteStartServer
}

function remoteStartServer(){
ssh -i $sshkey $url 2>&1 << eeooff
cd $remotePath
bash ./deploy.sh stop
nohup ./$server &
eeooff
}

function copyFile(){
  rsync -av --progress --rsh="ssh -i $sshkey" $1 $url:$2
}


if [ "$1" == "server" ]; then
  deployServer
fi


if [ "$1" == "stop" ]; then
  PROCESS=`ps -ef|grep $server|grep -v grep|grep -v PPID|awk '{ print $2}'`
  for i in $PROCESS
  do
    echo "Kill the $server process [ $i ]"
    kill -9 $i
  done
fi