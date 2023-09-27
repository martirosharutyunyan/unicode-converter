Steps For installation development environment and packaging application in arch linux

1. Install Qt5.13.0 https://download.qt.io/archive/qt/5.13/5.13.0/qt-opensource-linux-x64-5.13.0.run
2. Install qtsetup ```export GO111MODULE=on; go get -v github.com/therecipe/qt && go install -v -tags=no_env github.com/therecipe/qt/cmd/... && go mod vendor && rm -rf vendor/github.com/therecipe/env_linux_amd64_513 && git clone https://github.com/therecipe/env_linux_amd64_513.git vendor/github.com/therecipe/env_linux_amd64_513 && $(go env GOPATH)/bin/qtsetup```
3. ```go mod vendor``` if not working try with root privileges
4. go.mod vendor syncs only go files. Remove qt from vendor/github.com/therecipe and clone it manually from git ```git clone https://github.com/therecipe/qt.git```
5. Run ```sudo QT_DIR=$QT5.13.0_INSTALL_PATH qtsetup``` for installing dependencies in vendor
6. For linux packaging run ```cd ./cmd/unicode-converter-gui``` and execute ```sudo QT_DIR=$QT5.13.0_INSTALL_PATH qtdeploy build linux```
7. For windows packaging run ```cd ./cmd/unicode-converter-gui``` and execute ```sudo qtdeploy -docker build windows```

Steps for installation cli application
1. make tidy
2. make build