git add .
git commit -a -m 'Deploying now'
git push -u origin master
ssh -i ~/aws/sarawagi/sarawagi_ec2.pem ubuntu@52.3.212.51
cd ./the-exam
git pull
export GOPATH=/home/ubuntu/the-exam/server
export GOBIN=/home/ubuntu/the-exam/server/bin
gulp server:build
./server/bin/server
echo "Done"