MODULE ?= apple
.PHONY:clean
all: linux


linux:
	mkdir -p dist/linux
	CGO_ENABLED=1 go build --buildmode=plugin -o dist/linux/$(MODULE).so plugin/main.go
    ifdef BEANPATH
	    cp dist/linux/$(MODULE).so $(BEANPATH)
    endif
# arm:
# 	mkdir -p dist/arm
# 	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="-s -w" -mod=vendor -o dist/arm/$(MODULE) plugin/main.go

clean:
	rm -rf dist