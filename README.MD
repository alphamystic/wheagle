
# WHEAGLE
Wheagle is an open source SOC Toolkit that allows for both malicious/unethical and ethical features. It's a purple teaming toolkit
This allows for training of teams and simulation of RED and BLUE teaming exercises through purple teaming.
One can create their own payload generator and use wheagle as it,s c2 or have your own EDR endpoint and seive through results with your SOC team
UPX,Themida,PeCompact,AsPack,Kkrunchy WWPack32
### Wheagle Features
* Wheagle Net
    An agent allowing gui connections to the RatNet Network
* EDR
    Awesome EDR features like:
      1. OS Memory Dumping
      2. Realtime Detection and event management
      3. CVE Manager
      4. Asset Management
      5. Live Discovery
      6. Backup management
      7. Policy Management
      8. Custom Plugins for exploit testing, Detection techniques
* APT Manager
    1. Botnet Management
    2. RED Teaming and phishing campaigns
    3.  Tracking of APT's and organization's threats



## SERVERS
Wheagle servers:
1. HTTP Server
    Admin panel http/https server
2. FTP Server
    For uploading virus type/files (I have this somewhere and forgot where)
3. SSH Server
    TO be implemented soon
4. TCP Server
    Serves for both tls creating simple listeners
5. DNS Server
    Similar to tcp and for exfil



## Regards
All thanks to some great projects that are in play here removing most bottleneck for me
1. Screenshot at github.com/kbinani/screenshot
2. Sliver for basics for a C2 (Will try creating a dll/so for this) https://github.com/BishopFox/sliver
3. Merlin also helped in basic C2 Knowledge (Will also create adlll for this) https://github.com/Ne0nd0g/merlin



## INSTALLATION
Git clone
> git clone
> cd into directory
> go build wheagle.go
> ./wheagle

### Prerequisites
When running on kali or linux, install migw to create windows executables and dlls
> sudo apt install -y mingw-w64
> sudo apt -y install build-essential


### Running the HTTP Server
Cd into server directory
> cd server

Go run server.go or go build
Check if already installed
If not, prompt create user page
#### PyQt instalation for the GUI
(This only covers for debian instalation)
> sudo apt-get update && sudo apt-get upgrade

> sudo apt-get -y install build-essential openssl libssl-dev libssl1.0 libgl1-mesa-dev libqt5x11extras5

cd into that package directory and install
> chmod +x qt*.run


#### Generate keys for HTTPS server
Execute this while inside the server directory ~/wheagle/server

1. Create CA key
> openssl genrsa 2048 > certs/wheagle_ca_key.pem

2. Generate X-509 certificate for the CA
> openssl req -new -x509 -nodes -days 365000  -key certs/wheagle_ca_key.pem -out certs/wheagle_ca_cert.pem

3. Create Server Certificate Keys
> openssl req -newkey rsa:2048 -nodes -days 365000 -keyout certs/wheagle_server_key.pem -out certs/wheagle_server_req.pem

4. Create Server Certificate and Key
> openssl req -newkey rsa:2048 -nodes -days 365000 -keyout  certs/wheagle_server_key.pem -out certs/wheagle_server_req.pem

5. Generate Server x-509 certificate
> openssl x509 -req -days 365000 -set_serial 01 -in certs/wheagle_server_req.pem -out certs/wheagle_server_cert.pem -CA certs/wheagle_ca_cert.pem -CAkey certs/wheagle_ca_key.pem

6. Generate client keys and certificate
> openssl req -newkey rsa:2048 -nodes -days 365000 -keyout certs/wheagle_client_key.pem -out certs/wheagle_client_req.pem

7. Generate Client X509 certificate
> openssl x509 -req -days 365000 -set_serial 01 -in certs/wheagle_client_req.pem -out certs/wheagle_client_cert.pem -CA certs/wheagle_ca_cert.pem -CAkey certs/wheagle_ca_key.pem

8. Verify Server Certificate
> openssl verify -CAfile certs/wheagle_ca_cert.pem  certs/wheagle_server_cert.pem

9. Verify client certificate
> openssl verify -CAfile certs/wheagle_ca_cert.pem certs/wheagle_ca_cert.pem certs/wheagle_client_cert.pem


##COLLABORATION
Whatever you do with your git push please don't change our Awesome certs.(But you can and should change yours locally) DO NOT EVEN TOUCH IT UNLESS YOU ARE CHANGING ONLY YOURS.
Other than that I'm new at this so let's see what happens
