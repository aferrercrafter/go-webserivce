check_install:
    witch swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
    swagger generate spec -o ./swagger.yml --scan-models