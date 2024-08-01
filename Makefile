-include .env
export

# 设置工作目录
WORKDIR := w3contract/contract

.PHONY: all test clean deploy fund help install snapshot format anvil 

PRIVATE_KEY :=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
DEFAULT_ANVIL_KEY := 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

help:
	@echo "Usage:"
	@echo "  make deploy [ARGS=...]\n    example: make deploy ARGS=\"--network sepolia\""
	@echo ""
	@echo "  make fund [ARGS=...]\n    example: make fund ARGS=\"--network sepolia\""

all: clean remove install update build

# Clean the repo
clean:
	@cd $(WORKDIR) && forge clean

# Remove modules
remove:
	@cd $(WORKDIR) && rm -rf .gitmodules && rm -rf .git/modules/* && rm -rf lib && touch .gitmodules && git add . && git commit -m "modules"

install:
	@cd $(WORKDIR) && forge install Openzeppelin/openzeppelin-contracts --no-commit && forge install foundry-rs/forge-std --no-commit

# Update Dependencies
update:
	@cd $(WORKDIR) && forge update

build:
	@cd $(WORKDIR) && forge build

test:
	@cd $(WORKDIR) && forge test 

snapshot:
	@cd $(WORKDIR) && forge snapshot

format:
	@cd $(WORKDIR) && forge fmt

anvil:
	@anvil -m 'test test test test test test test test test test test junk' --steps-tracing --block-time 1

NETWORK_ARGS := --rpc-url http://localhost:8545 --private-key $(DEFAULT_ANVIL_KEY) --broadcast

ifeq ($(findstring --network sepolia,$(ARGS)),--network sepolia)
	NETWORK_ARGS := --rpc-url $(SEPOLIA_RPC_URL) --private-key $(PRIVATE_KEY) --broadcast --verify --etherscan-api-key $(ETHERSCAN_API_KEY) -vvvv
endif

# deploy:
# 	@echo "Current directory: $$(pwd)"
# 	@echo "WORKDIR: $(WORKDIR)"
# 	@cd $(WORKDIR) && echo "Changed to: $$(pwd)" && forge script script/DeployFileStorage.s.sol:DeployFileStorage $(NETWORK_ARGS)

deploy:
	@echo "WORKDIR: $(WORKDIR)"
	@cd $(WORKDIR) && forge script script/DeployFileStorage.s.sol:DeployFileStorage $(NETWORK_ARGS)

# deploy:
#     @echo "WORKDIR: $(WORKDIR)"
#     @cd $(WORKDIR) && forge script script/DeployFileStorage.s.sol:DeployFileStorage --fork-url $(RPC_URL) --private-key $(PRIVATE_KEY) -vvvv

# 更新为你的合约地址
verify:
	@cd $(WORKDIR) && forge verify-contract --chain-id 11155111 --num-of-optimizations 200 --watch --constructor-args 0x00000000000000000000000000000000000000000000d3c21bcecceda1000000 --etherscan-api-key $(ETHERSCAN_API_KEY) --compiler-version v0.8.19+commit.7dd6d404 0x089dc24123e0a27d44282a1ccc2fd815989e3300 src/FileStorage.sol:FileStorage