#====================
AUTHOR         ?= The sacloud/iam-api-go Authors
COPYRIGHT_YEAR ?= 2025-

BIN            ?= iam-api-go
GO_FILES       ?= $(shell find . -name '*.go')

include includes/go/common.mk
include includes/go/single.mk
#====================

default: $(DEFAULT_GOALS)
tools: dev-tools
