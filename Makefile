.DEFAULT_GOAL := build
bin=loxilb
dock?=loxilb

loxilbid=$(shell docker ps -f name=$(dock) | grep -w $(dock) | cut  -d " "  -f 1 | grep -iv  "CONTAINER")

subsys:
	cd loxilb-ebpf && $(MAKE) 

subsys-clean:
	cd loxilb-ebpf && $(MAKE) clean

build: subsys
	@go build -o ${bin} -ldflags="-X 'main.buildInfo=${shell date '+%Y_%m_%d'}-${shell git branch --show-current}'"

clean: subsys-clean
	go clean

test:
	go test .

check:
	go test .

run:
	./$(bin)

docker-cp: build
	docker cp loxilb $(loxilbid):/root/loxilb-io/loxilb/loxilb
	docker cp /opt/loxilb/llb_ebpf_main.o $(loxilbid):/opt/loxilb/llb_ebpf_main.o
	docker cp /opt/loxilb/llb_ebpf_emain.o $(loxilbid):/opt/loxilb/llb_ebpf_emain.o
	docker cp /opt/loxilb/llb_xdp_main.o $(loxilbid):/opt/loxilb/llb_xdp_main.o
	docker cp loxilb-ebpf/kernel/loxilb_dp_debug  $(loxilbid):/usr/local/sbin/

docker-cp-ebpf: build
	docker cp /opt/loxilb/llb_ebpf_main.o $(loxilbid):/opt/loxilb/llb_ebpf_main.o
	docker cp /opt/loxilb/llb_ebpf_emain.o $(loxilbid):/opt/loxilb/llb_ebpf_emain.o
	docker cp /opt/loxilb/llb_xdp_main.o $(loxilbid):/opt/loxilb/llb_xdp_main.o

docker-run:
	@docker stop $(dock) 2>&1 >> /dev/null || true
	@docker rm $(dock) 2>&1 >> /dev/null || true
	docker run -u root --cap-add SYS_ADMIN   --restart unless-stopped --privileged -dt --entrypoint /bin/bash  --name $(dock) ghcr.io/loxilb-io/loxilb:latest

docker-rp: docker-run docker-cp
	docker commit ${loxilbid} ghcr.io/loxilb-io/loxilb:latest
	@docker stop $(dock) 2>&1 >> /dev/null || true
	@docker rm $(dock) 2>&1 >> /dev/null || true

docker-rp-ebpf: docker-run docker-cp-ebpf
	docker commit ${loxilbid} ghcr.io/loxilb-io/loxilb:latest
	@docker stop $(dock) 2>&1 >> /dev/null || true
	@docker rm $(dock) 2>&1 >> /dev/null || true

docker:
	docker build -t ghcr.io/loxilb-io/loxilb:latest .

docker-arm64:
	docker  buildx build --platform linux/arm64 -t ghcr.io/loxilb-io/loxilb:latest-arm64 .

lint:
	golangci-lint run --enable-all

stop:
	@rm -rf /opt/loxilb/dp/bpf/bd_stats_map
	@rm -rf /opt/loxilb/dp/bpf/cpu_map
	@rm -rf /opt/loxilb/dp/bpf/crc32c_map
	@rm -rf /opt/loxilb/dp/bpf/ct_ctr
	@rm -rf /opt/loxilb/dp/bpf/ct_map
	@rm -rf /opt/loxilb/dp/bpf/ct_stats_map
	@rm -rf /opt/loxilb/dp/bpf/dmac_map
	@rm -rf /opt/loxilb/dp/bpf/fc_v4_map
	@rm -rf /opt/loxilb/dp/bpf/fc_v4_stats_map
	@rm -rf /opt/loxilb/dp/bpf/fcas
	@rm -rf /opt/loxilb/dp/bpf/fw_v4_map
	@rm -rf /opt/loxilb/dp/bpf/fw_v4_stats_map
	@rm -rf /opt/loxilb/dp/bpf/gparser
	@rm -rf /opt/loxilb/dp/bpf/intf_map
	@rm -rf /opt/loxilb/dp/bpf/intf_stats_map
	@rm -rf /opt/loxilb/dp/bpf/live_cpu_map
	@rm -rf /opt/loxilb/dp/bpf/mirr_map
	@rm -rf /opt/loxilb/dp/bpf/nat_map
	@rm -rf /opt/loxilb/dp/bpf/nat_stats_map
	@rm -rf /opt/loxilb/dp/bpf/nh_map
	@rm -rf /opt/loxilb/dp/bpf/pgm_tbl
	@rm -rf /opt/loxilb/dp/bpf/pkt_ring
	@rm -rf /opt/loxilb/dp/bpf/pkts
	@rm -rf /opt/loxilb/dp/bpf/polx_map
	@rm -rf /opt/loxilb/dp/bpf/rt_v4_map
	@rm -rf /opt/loxilb/dp/bpf/rt_v4_stats_map
	@rm -rf /opt/loxilb/dp/bpf/rt_v6_map
	@rm -rf /opt/loxilb/dp/bpf/rt_v6_stats_map
	@rm -rf /opt/loxilb/dp/bpf/sess_v4_map
	@rm -rf /opt/loxilb/dp/bpf/sess_v4_stats_map
	@rm -rf /opt/loxilb/dp/bpf/smac_map
	@rm -rf /opt/loxilb/dp/bpf/tmac_map
	@rm -rf /opt/loxilb/dp/bpf/tmac_stats_map
	@rm -rf /opt/loxilb/dp/bpf/tx_bd_stats_map
	@rm -rf /opt/loxilb/dp/bpf/tx_intf_map
	@rm -rf /opt/loxilb/dp/bpf/tx_intf_stats_map
	@rm -rf /opt/loxilb/dp/bpf/xctk
	@rm -rf /opt/loxilb/dp/bpf/xfck
	@rm -rf /opt/loxilb/dp/bpf/xfis
	@ip link delete llb0 || true