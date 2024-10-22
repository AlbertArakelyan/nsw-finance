BINARY_NAME=NSWFinance.app
BINARY_NAME_WINDOWS=NSWFinance.exe
APP_ID_WINDOWS=com.nswfinance.aa
APP_NAME=NSWFinance
VERSION=0.1.7
BUILD_NO=1

## build: build binary and package app
build:
	rm -rf ${BINARY_NAME}
	fyne package -appVersion ${VERSION} -appBuild ${BUILD_NO} -name ${APP_NAME} -release

build-windows:
	rm ${BINARY_NAME_WINDOWS}
	fyne package -os windows -name ${BINARY_NAME_WINDOWS} -appID ${APP_ID_WINDOWS} -release

## run: builds and runs the application
run:
	go run .

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf ${BINARY_NAME}
	@echo "Cleaned!"