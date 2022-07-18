# Raspio
>**(Work in Progress)** This project is still under development and is not complete, do not use for a production system.

Raspio is development tool that after installed on a device (most likely a raspberry pi), will allow any clients with the authentication key to upload and update payloads on the device for easy testing and development.

## Installation
There are two different devices this program needs to be installed under, both the client or your development device and the server, raspberry pi, or any other device or server you are using to run the code you are developing.
> NOTE: This program is an all in one, aka all of the server and client functionality are installed in the same executable meaning you only have to download one file for both the client and the server. If you execute the built program with no arguments it will by default run in client mode however for server setup you will be adding spesial arguments to tell it to run in server mode.

### Client / Development Device
1. Download the latest version of the software from GitHub
2. Run the executable to see how to connect to a hosting device
> Make sure you don't use any of the arguments listed in the Hosting Device section or your client will become a host

### Raspberry Pi / Hosting Device
1. Download the latest version of the software from GitHub
2. Run the executable with three arguments;
   - `-server` to tell it to run in server mode
   - `-bind` followed by this `IP:PORT`
   - `-key` and the key you wish to use. This key will have to be sent by any clients to upload or update projects.

**MAKE SURE THIS KEY IS LONG AND COMPLEX, ANYONE WITH THIS KEY WILL BE ABLE TO UPLOAD AND UPDATE PROJECTS ON YOUR HOST DEVICE**

Full command will look something like `raspio -server -bind 127.0.0.1:3030 -key myreallygoodkey`

4. *(Optional but recommended)* Add the complete command into a startup item so when the computer is powered the server will automatically start and get ready to recieve clients.