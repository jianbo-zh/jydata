
VERSION:=v1.0.0

.PHONY: build
# build
build:
	mkdir -p .bin/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o ./.bin/jydata ./cmd/migrate

.PHONY: deployyk
deploy_root_yk := /data/wwwyokee/ykapi.grandrace.tech
deployyk:
	@ssh jiuyou "cd ${deploy_root_yk}/bin/ && { if [ -f jydata_yk.bak2 ]; then rm jydata_yk.bak2; fi } && { if [ -f jydata_yk.tmp ]; then rm jydata_yk.tmp; fi }"
	@scp .bin/jydata jiuyou:${deploy_root_yk}/bin/jydata_yk.tmp
	@ssh jiuyou "cd ${deploy_root_yk}/bin/ && { if [ -e jydata_yk.bak ]; then mv jydata_yk.bak jydata_yk.bak2; fi } && { if [ -e jydata_yk ]; then mv jydata_yk jydata_yk.bak; fi } && mv jydata_yk.tmp jydata_yk"
	@ssh jiuyou "${deploy_root_yk}/jydata_restart.sh"

.PHONY: deploydev
deploy_root_dev := /data/wwwdev/api.joyorover.com2
deploydev:
	@ssh jiuyou "cd ${deploy_root_dev}/bin/ && { if [ -f jydata_2dev.bak2 ]; then rm jydata_2dev.bak2; fi } && { if [ -f jydata_2dev.tmp ]; then rm jydata_2dev.tmp; fi }"
	@scp .bin/jydata jiuyou:${deploy_root_dev}/bin/jydata_2dev.tmp
	@ssh jiuyou "cd ${deploy_root_dev}/bin/ && { if [ -e jydata_2dev.bak ]; then mv jydata_2dev.bak jydata_2dev.bak2; fi } && { if [ -e jydata_2dev ]; then mv jydata_2dev jydata_2dev.bak; fi } && mv jydata_2dev.tmp jydata_2dev"
	@ssh jiuyou "${deploy_root_dev}/jydata_restart.sh"