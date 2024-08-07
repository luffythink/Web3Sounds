-include .env
export

# 设置工作目录
WORKDIR := w3contract/contract

.PHONY: all test clean deploy fund help install snapshot format anvil 

DEFAULT_ANVIL_KEY := $(OWNER_PRIVATE_KEY)
help:
	@echo "Usage:"
	@echo "  make deploy [ARGS=...]\n    example: make deploy ARGS=\"--network sepolia\""
	@echo ""
	@echo "  make fund [ARGS=...]\n    example: make fund ARGS=\"--network sepolia\""

all: clean remove install update build

# Clean the repo
clean:
	@cd $(WORKDIR) && forge clean && rm -rf broadcast

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

testecho:
	@echo "ETHERSCAN_API_KEY: $(ETHERSCAN_API_KEY)"
	@echo "OWNER_PRIVATE_KEY: $(OWNER_PRIVATE_KEY)"
	@echo "DEFAULT_ANVIL_KEY: $(DEFAULT_ANVIL_KEY)"

anvil:
	@anvil -m 'test test test test test test test test test test test junk' --steps-tracing --block-time 1

NETWORK_ARGS := --rpc-url $(RPC_URL) --private-key $(DEFAULT_ANVIL_KEY) --broadcast --gas-price 15000000000 --gas-limit 25718 -vvvv

ifeq ($(findstring --network sepolia,$(ARGS)),--network sepolia)
	NETWORK_ARGS := --rpc-url $(SEPOLIA_RPC_URL) --private-key $(PRIVATE_KEY) --broadcast --verify --etherscan-api-key $(ETHERSCAN_API_KEY) -vvvv
endif

# deploy:
# 	@echo "Current directory: $$(pwd)"
# 	@echo "WORKDIR: $(WORKDIR)"
# 	@cd $(WORKDIR) && echo "Changed to: $$(pwd)" && forge script script/DeployFileStorage.s.sol:DeployFileStorage $(NETWORK_ARGS)
CHAIN_ID := $(shell cast chain-id --rpc-url $(RPC_URL))
BALANCE := $(shell cast balance $(OWNER_ADDRESS) --rpc-url $(RPC_URL))
NONCE := $(shell cast nonce $(OWNER_ADDRESS) --rpc-url $(RPC_URL))
CURRENT_GAS_PRICE := $(shell cast gas-price --rpc-url $(RPC_URL))
# 可以选择性地增加一点 gas 价格，以确保交易能够快速被确认
GAS_PRICE_BUFFER := 1000000000  # 1 gwei 的缓冲
GAS_PRICE := $(shell echo "$(CURRENT_GAS_PRICE) + $(GAS_PRICE_BUFFER)" | bc)
deploy:
	@echo "WORKDIR: $(WORKDIR)"
	@echo "RPC_URL: $(RPC_URL)"
	@echo "CHAIN_ID: $(CHAIN_ID)"
	@echo "OWNER_ADDRESS: $(OWNER_ADDRESS)"
	@echo "balance: $(BALANCE)"
	@echo "nonce: $(NONCE)"
	@echo "ONline_CURRENT_GAS_PRICE: $(CURRENT_GAS_PRICE) wei || GAS_PRICE: $(GAS_PRICE) wei"
	@cast wallet address $(DEFAULT_ANVIL_KEY)
	@echo "==========exec deploy============"
	@cd $(WORKDIR) && forge script script/DeployFileStorage.s.sol:DeployFileStorage \
    --rpc-url $(RPC_URL) \
    --chain-id $(CHAIN_ID) \
    --private-key $(DEFAULT_ANVIL_KEY) \
    --broadcast \
    --gas-price $(GAS_PRICE) \
    --legacy \
    --slow \
    -vvvv
	
	
# @cd $(WORKDIR) && forge script script/DeployFileStorage.s.sol:DeployFileStorage --rpc-url $(RPC_URL) --chain-id $(CHAIN_ID)  --private-key $(DEFAULT_ANVIL_KEY) --broadcast --auto  --legacy  --slow  -vvvv
# --gas-price 18gwei --gas-limit 2000000
# deploy: 
#     @echo "WORKDIR: $(WORKDIR)"
#     @cd $(WORKDIR) && forge script script/DeployFileStorage.s.sol:DeployFileStorage --fork-url $(RPC_URL) --private-key $(PRIVATE_KEY) -vvvv

# 更新为你的合约地址
verify:
	@cd $(WORKDIR) && forge verify-contract --chain-id 3334 --num-of-optimizations 200 --watch --constructor-args 0x00000000000000000000000000000000000000000000d3c21bcecceda1000000 --etherscan-api-key $(ETHERSCAN_API_KEY) --compiler-version v0.8.19+commit.7dd6d404 0x089dc24123e0a27d44282a1ccc2fd815989e3300 src/FileStorage.sol:FileStorage