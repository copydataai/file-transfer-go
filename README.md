# A simple file transfer (One To Many)

## Install 
Before to make something, you need clone the repo

``` sh
git clone https://github.com/copydataai/file-transfer-go
cd file-transfer-go
```
if do you'll install a specific CLI
*General*
| Server:               | Client:               | All-one            |
|-----------------------|-----------------------|--------------------|
| `make install-server` | `make install-client` | `make install-all` |


## Graphs 
clientSend -> Server 
    
    Server -> ClientReceive
    Server -> ClientReceive
    Server to N Devices
    

*NOTE:*It's similar to FTP, but this use for learn and understand network with Go💙.
## Custom protocol
I use a similar format from HTTP, but just for send filename, fileSize and file 

* For send a file 
``` rich-text-format
Channel: 1 PUT
Filename: hello.txt
FileSize: 138
[]bytes(content) <- file
```
* Send to Receive Devices
``` rich-text-format
Filename: hello.txt
FileSize: 138
[]bytes(content) <- file
```

## TODO
* Create TCP Clients to receive using actual path
* Create a Makefile to install using ENVs 
* Create a dir to binarys
| Example                                         |
| `export FPATH=$PATH:/home/$USER/filetransfer-go` |
