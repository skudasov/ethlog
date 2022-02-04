#!/bin/bash
abi_dir="abi"
mkdir -p ${abi_dir}
contract1="events_test_contract"
contract1_camel="EventsTestContract"
erc20="erc20"
mkdir -p ${contract1}
contract2="another_events_test_contract"
contract2_camel="AnotherEventsTestContract"
mkdir -p ${erc20}
erc20_camel="ERC20"
mkdir -p ${contract2}
solc --abi ${contract1}.sol --bin ${contract1}.sol -o ${abi_dir} --overwrite
solc --abi ${contract2}.sol --bin ${contract2}.sol -o ${abi_dir} --overwrite
abigen --abi=abi/${contract1_camel}.abi --bin=abi/${contract1_camel}.bin --pkg=${contract1} --out=${contract1}/${contract1_camel}.go
abigen --abi=abi/${contract2_camel}.abi --bin=abi/${contract2_camel}.bin --pkg=${contract2} --out=${contract2}/${contract2_camel}.go
abigen --abi=abi/${erc20_camel}.abi --pkg=${erc20} --out=${erc20}/${erc20_camel}.go
