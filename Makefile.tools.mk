#!make

nproc ?= 2

clang ?= $(shell lsb_release -r | cut -f2 | sed s/22.04/clang-13/ | sed s/20.04/clang-10/)

TARGETS := linux/amd64 linux/arm64
DIST_DIRS := find * -type d -exec
GOPATH = $(shell go env GOPATH)
GOBIN  = $(GOPATH)/bin
GOX    = go run github.com/mitchellh/gox
SHA256 = sha256sum
ifeq ($(shell uname),Darwin)
	SHA256 = shasum -a 256
endif

.PHONY: depends
depends:
	@apt -y update
	@arch=$(arch | sed s/aarch64/arm64/ | sed s/x86_64/amd64/) && echo $arch && if [ "$arch" = "arm64" ] ; then apt install -y gcc-multilib-arm-linux-gnueabihf; else apt update && apt install -y gcc-multilib;fi
	@apt install -y $(clang) llvm libelf-dev libpcap-dev net-tools elfutils dwarves git libbsd-dev bridge-utils wget arping unzip build-essential bison flex sudo iproute2 pkg-config tcpdump iputils-ping keepalived curl bash-completion
	@apt -y autoremove
	@cp loxilb-ebpf/utils/mkllb_bpffs.sh /usr/local/sbin/mkllb_bpffs
	@mkdir -p /opt/loxilb/cert
	@cp api/certification/* /opt/loxilb/cert/
	@if [ ! -f /usr/local/sbin/bpftool ]; then git clone --recurse-submodules https://github.com/libbpf/bpftool.git && cd bpftool/src/ && make clean && make -j $(nproc) && cp -f ./bpftool /usr/local/sbin/bpftool && cd - && rm -fr bpftool; fi
	@if [ ! -f /usr/local/sbin/ntc ]; then wget https://github.com/cybwan/iproute2/archive/refs/heads/main.zip && unzip main.zip && cd iproute2-main/libbpf/src/ && mkdir build && DESTDIR=build make install && cd - && cd iproute2-main/ && export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:`pwd`/libbpf/src/ && LIBBPF_FORCE=on LIBBPF_DIR=`pwd`/libbpf/src/build ./configure && make && cp -f tc/tc /usr/local/sbin/ntc && cd - && cd iproute2-main/libbpf/src/ && make install && cd - && rm -fr main.zip iproute2-main; fi
	@if [ ! -f /usr/lib64/libbpf.so.0.4.0 ]; then cd bpf && make && make install && cd -; fi
	@if [ ! -f /usr/sbin/gobgp ]; then arch=${shell arch | sed s/aarch64/arm64/ | sed s/x86_64/amd64/} && echo https://github.com/osrg/gobgp/releases/download/v3.5.0/gobgp_3.5.0_linux_$${arch}.tar.gz.tar.gz && wget https://github.com/osrg/gobgp/releases/download/v3.5.0/gobgp_3.5.0_linux_$${arch}.tar.gz && tar -xzf gobgp_3.5.0_linux_$${arch}.tar.gz && rm gobgp_3.5.0_linux_$${arch}.tar.gz LICENSE README.md && mv gobgp* /usr/sbin/; fi
	@if [ ! -f /usr/local/go/bin/go ]; then arch=${shell arch | sed s/aarch64/arm64/ | sed s/x86_64/amd64/} && echo https://go.dev/dl/go1.19.linux-$${arch}.tar.gz && wget https://go.dev/dl/go1.19.linux-$${arch}.tar.gz && tar -xzf go1.19.linux-$${arch}.tar.gz --directory /usr/local/ && rm go1.19.linux-$${arch}.tar.gz;echo please export PATH=\$${PATH}:/usr/local/go/bin; fi
	@ kver=${shell uname -r | cut -d"-" -f1} && echo $${kver} && if $$(dpkg --compare-versions $${kver}  "lt" "5.14"); then ukv=${shell sudo apt list linux-image-5.*-generic 2>&1 | grep ^linux | cut -d '-' -f 3,4 | sort -rV | head -n1} && sudo apt install -y linux-modules-$${ukv}-generic linux-headers-$${ukv}-generic linux-image-$${ukv}-generic; fi
	@echo please export PATH=\$${PATH}:/usr/local/go/bin

subsys:
	mkdir -p /opt/loxilb/cert
	cp loxilb-ebpf/utils/mkllb_bpffs.sh /usr/local/sbin/mkllb_bpffs
	cp api/certification/* /opt/loxilb/cert/
	cd loxilb-ebpf && make

subsys-clean:
	cd loxilb-ebpf && make clean
