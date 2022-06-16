##
# File Transfer GO
#
# @file
# @version 0.1



# end


begin:
	mkdir /home/$USER/file-transfer
	export FPATH=/home/$USER/file-transfer
	export PATH=$PATH:$FPATH

install-server:
	cd backend/cmd/
	go build -o srvfile .
	mv srvfile $FPATH/
	cd ..
	cd ftp/
	go build -o file-transfer-go .
	mv file-transfer-go $FPATH/

install-client:
	cd client-go/
	go build -o cltfile .
	mv cltfile $FPATH/
