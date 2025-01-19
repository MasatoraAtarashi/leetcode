export GOBIN:=$(CURDIR)/bin

$(GOBIN)/leetgode:
	go install github.com/budougumi0617/leetgode/cmd/leetgode

$(GOBIN)/gosimports:
	go install github.com/rinchsan/gosimports/cmd/gosimports@latest

init: $(GOBIN)/leetgode
	@echo "Please set up your LeetCode credentials:"
	@echo "1. Open chrome://settings/cookies/detail?site=leetcode.com"
	@echo "2. Copy LEETCODE_SESSION and csrftoken values"
	@echo "3. Export them as environment variables:"
	@echo "   export LEETCODE_SESSION=your_session_value"
	@echo "   export LEETCODE_TOKEN=your_csrf_token"
	@echo "\nAvailable commands:"
	@$(GOBIN)/leetgode help

check-credentials:
	@if [ -z "$$LEETCODE_SESSION" ]; then \
		read -p "Enter your LEETCODE_SESSION value: " session; \
		echo "export LEETCODE_SESSION=$$session" >> ~/.bashrc; \
		export LEETCODE_SESSION=$$session; \
	fi
	@if [ -z "$$LEETCODE_TOKEN" ]; then \
		read -p "Enter your LEETCODE_TOKEN (csrftoken) value: " token; \
		echo "export LEETCODE_TOKEN=$$token" >> ~/.bashrc; \
		export LEETCODE_TOKEN=$$token; \
	fi

# List all problems
list:
	$(GOBIN)/leetgode list

# Generate solution template for a problem
# Usage: make generate id=1234
generate:
	$(GOBIN)/leetgode generate $(id)

# Test your solution
# Usage: make test id=1234
test: check-credentials
	$(GOBIN)/leetgode test $(id)

# Submit your solution
# Usage: make exec id=1234
exec: check-credentials
	$(GOBIN)/leetgode exec $(id)

.PHONY: fmt
fmt: $(GOBIN)/gosimports ## Run goimports
	go fmt ./...
