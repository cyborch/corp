# Configurable options
IMAGE_NAME = cyborch/corp

ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

include scripts/*.mk

all: corp

clean: corp-clean docker-clean
