# Install Go

Define your Go version, OS and architecture
```
export VERSION="1.14.7"
export OS="linux"
export ARCH="amd64"
```

Download the binary release
```
wget https://golang.org/dl/go$VERSION.$OS-$ARCH.tar.gz
```

Remove previous installation
```
sudo rm -Rf /usr/local/go/
```

Extract the release into /usr/local
```
sudo tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
```

Add the Go folder to the path in the profile file
```
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
```

Test the Go version
```
go version
> go version go1.14.7 linux/amd64
```