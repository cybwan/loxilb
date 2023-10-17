wget https://github.com/loxilb-io/loxilb/raw/main/cicd/common/sctp_client
wget https://github.com/loxilb-io/loxilb/raw/main/cicd/common/udp_client
chmod 777 sctp_client
chmod 777 udp_client
echo "123.123.123.1 k8s-svc" >> /etc/hosts

sudo apt install -y bird2 socat

sleep 5

sudo cp -f /vagrant/bird.conf /etc/bird/bird.conf
if [ ! -f  /var/log/bird.log ]; then
  sudo touch /var/log/bird.log
fi
sudo chown bird:bird /var/log/bird.log
sudo systemctl restart bird
echo "Host is up"