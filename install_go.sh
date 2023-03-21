curl https://go.dev/dl/go1.20.1.linux-amd64.tar.gz

# if go was previously installed
# rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.1.linux-amd64.tar.gz

#unfurl
tar -C /usr/local -xzf go1.20.1.linux-amd64.tar.gz

#add path env var
export PATH=$PATH:/usr/local/go/bin


#check version
go version