Steps For installation development environment and packaging application in arch linux
1. Install stringer tool ```go install golang.org/x/tools/cmd/stringer@latest``` 
2. Install Qt5.13.0 https://download.qt.io/archive/qt/5.13/5.13.0/qt-opensource-linux-x64-5.13.0.run
3. Install qtsetup ```export GO111MODULE=off; go get -v github.com/therecipe/qt/cmd/... && $(go env GOPATH)/bin/qtsetup test && $(go env GOPATH)/bin/qtsetup -test=false && go mod vendor``` if not working try with root privileges
4. go.mod vendor syncs only go files. Remove qt from vendor/github.com/therecipe and clone it manually from git ```git clone https://github.com/therecipe/qt.git```
5. Run ```sudo QT_DIR=$QT5.13.0_INSTALL_PATH qtsetup``` for installing dependencies in vendor
6. For linux packaging run ```cd ./cmd/unicode-converter-gui``` and execute ```sudo QT_DIR=$QT5.13.0_INSTALL_PATH qtdeploy build linux```
7. Docker pull therecipe/qt:windows_64_static
8. rsrc install for application icon ```go install github.com/akavel/rsrc@latest```
9. rsrc -ico icon.ico -o icon.syso -arch=amd64 for setting the application icon replace the icon.ico file with your icon
10. For windows packaging run ```cd ./cmd/unicode-converter-gui``` and execute ```sudo qtdeploy -docker build windows_64_static```


Steps for installation cli application
1. make tidy
2. make build