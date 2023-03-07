# P2PRC CHERI Qemu SSH 
The following repository was designed to SSH into 
the CHERI Qemu emulator from anywhere using 
P2PRC TURN based implementation. 

## Build instructions
- [Build the CHERI Qemu image](https://github.com/CTSRD-CHERI/cheribuild)
- Build the go project for the TURN implementation  
```
go build .
```
- Run the project
```
./p2prc-cheri-qemu
```

## Sample server 
Feel free to try out the sample server I am currently hosting. 
```
ssh root@64.227.168.102 -p34587
```