
export project_name='prj_1' 
sudo apt-get -y install openssh-server


sudo chmod 777 /etc/apt/sources.list

wget -q -O - https://pkg.jenkins.io/debian/jenkins-ci.org.key | sudo apt-key add -
wget -q -O - https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
sudo sh -c 'echo deb http://pkg.jenkins.io/debian-stable binary/ > /etc/apt/sources.list.d/jenkins.list'
sudo sh -c 'echo deb http://apt.kubernetes.io/ kubernetes-xenial main > /etc/apt/sources.list.d/kubernetes.list'

sudo chmod 644 /etc/apt/sources.list
sudo apt-get -y update
sudo apt-get install -y kubelet kubeadm kubectl kubernetes-cni
#sudo apt-get -y install jenkins
cat /etc/default/jenkins |  sed -e "s/HTTP_PORT=8080/HTTP_PORT=666/" > /tmp/jenkins
sudo cp -f /tmp/jenkins /etc/default/jenkins
sudo /etc/init.d/jenkins restart

sudo apt-get -y install git
mkdir -p $project_name/ping  
mkdir -p $project_name/revel  
cd $project_name
sudo apt-get -y install unzip

rm master.zip*

wget . https://github.com/joToui/pingponggo/archive/master.zip

echo `pwd`

unzip -o master.zip

cp pingponggo-master/requirements.txt ping/.
cp pingponggo-master/ping.py ping/.
cp pingponggo-master/restserver.go  revel/.
cd ping 
sudo apt install -y docker.io

echo " killing all running docker containers .... "
sudo  docker kill $(sudo docker ps -q)


echo "building Dockerfile .... "
cat >Dockerfile << DOCFL 

FROM ubuntu:14.04
# Update OS
RUN sed -i 's/# \(.*multiverse$\)/\1/g' /etc/apt/sources.list
RUN apt-get -y update
RUN apt-get -y upgrade

# Install Python
RUN apt-get -y install python-pip
RUN apt-get -y install python-dev 
RUN apt-get -y install python-pycurl
#COPY requirements.txt .

COPY requirements.txt /tmp/

# Install uwsgi Python web server
#RUN pip install uwsgi
# Install app requirements

# Set the default directory for our environment

RUN mkdir -p /jinny/$project_name/ping

ENV HOME /jinny/$project_name/ping

WORKDIR /jinny/$project_name/ping 


ENV FLASK_APP /jinny/$project_name/ping/ping.py

COPY ping.py /jinny/$project_name/ping/.

#COPY test1 $HOME/test1
RUN echo `pwd `

#RUN echo `ls /webapp/requirements.txt`

RUN pip install -r /tmp/requirements.txt

# Create app directory


#ENTRYPOINT ["python"]

#RUN ["young_30.py"]
RUN echo \`pwd \`
RUN echo \`ls \`
RUN echo \`which flask \`

ENTRYPOINT [ "python", "/jinny/$project_name/ping/ping.py", "-p", "1948", "-h", "0.0.0.0" ]
# Expose port 1948 for uwsgi
EXPOSE 1948


DOCFL

echo "ping out_of Dockerfile  "
echo "building image .... "
sudo docker build . -t flask_image
echo "building image .... "
echo "running  image .... "

echo " image is up .... "

echo "so new image is needed for the go restserver "

################################################################

cd ..
mkdir -p revel
cd revel
echo "building Dockerfile .... "
cat >Dockerfile << DOCFLG

FROM ubuntu:14.04
# Update OS
RUN sed -i 's/# \(.*multiverse$\)/\1/g' /etc/apt/sources.list
RUN apt-get -y update
RUN apt-get -y upgrade

# Install Python
RUN sudo apt-get -y install golang-go
#COPY requirements.txt .

# Set the default directory for our environment

RUN mkdir -p /jinny/$project_name/revel

ENV HOME /jinny/$project_name/revel

WORKDIR /jinny/$project_name/revel



COPY restserver.go /jinny/$project_name/revel/.

#COPY test1 $HOME/test1
RUN echo `pwd `

#RUN echo `ls /webapp/requirements.txt`

#RUN ["young_30.py"]

ENTRYPOINT [ "go", "run", "restserver.go" ]
# Expose port 1949 for uwsgi
EXPOSE 1949


DOCFLG

echo "ping out_of Dockerfile  "
echo "building image .... "
sudo docker build . -t go_revel
echo "building image .... "
echo "running  image .... "
revel_server_host=$(sudo docker run -itd -p 0.0.0.0:1949:1949 go_revel)

revel_server_host_ip=$(sudo docker inspect --format='{{.NetworkSettings.Networks.bridge.IPAddress}}' $revel_server_host)

sudo docker run -itd  --add-host revel_server_host:$revel_server_host_ip -p 0.0.0.0:1948:1948 flask_image &


echo " image is up .... "

echo "so new image is needed for the go restserver "




