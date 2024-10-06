BINARY_NAME=NSWFinance.app
APP_NAME=NSWFinance
VERSION=0.1.4
BUILD_NO=1

## build: build binary and package app
build:
	rm -rf ${BINARY_NAME}
	fyne package -appVersion ${VERSION} -appBuild ${BUILD_NO} -name ${APP_NAME} -release
	rm -f go-for-gold

## run: builds and runs the application
run:
	go run .

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf ${BINARY_NAME}
	@echo "Cleaned!"