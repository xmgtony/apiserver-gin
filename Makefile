go_version = $(shell go version)
commit_id = $(shell git rev-parse HEAD)
branch_name = $(shell git name-rev --name-only HEAD)
build_time = $(shell date -u '+%Y-%m-%d_%H:%M:%S')
app_version = 1.0.0
version_package = apiserver-gin/pkg/version
app_name = apiserver-gin
work_dir = target
all: package

build: target
	@go build -ldflags \
	"-X ${version_package}.CommitId=${commit_id} \
	-X ${version_package}.BranchName=${branch_name} \
	-X ${version_package}.BuildTime=${build_time} \
	-X ${version_package}.AppVersion=${app_version}" -v \
	-o ${work_dir}/${app_name} ./cmd/.
# show go version
version:
	@$(go_version)
	@echo APP_VERSION $(app_version)
clean:
	@rm -rf target
target:
	@mkdir ${work_dir}

.ONESHELL:
package: build
	@# 使用tar命令对${word_dir下面的文件打包}
	cp -r conf  ${work_dir}/
	cp ./scripts/startup.sh ${work_dir}/
	cd ${work_dir}/ && tar -zcvf ${app_name}.tar.gz *

.PHONY: version clean build package all