// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

import { GroupMessages } from "../src/GroupMessages.sol";

import { Utils } from "./utils/Utils.sol";
import { Environment } from "./utils/Environment.sol";

contract DeployGroupMessages is Utils, Environment {
    GroupMessages groupMessagesImpl;
    ERC1967Proxy proxy;

    address admin;
    address deployer;

    function run() external {
        admin = vm.envAddress("XMTP_GROUP_MESSAGES_ADMIN_ADDRESS");
        require(admin != address(0), "XMTP_GROUP_MESSAGES_ADMIN_ADDRESS not set");

        uint256 privateKey = vm.envUint("PRIVATE_KEY");
        require(privateKey != 0, "PRIVATE_KEY not set");

        deployer = vm.addr(privateKey);
        vm.startBroadcast(privateKey);

        // Deploy the implementation contract.
        groupMessagesImpl = new GroupMessages();
        require(address(groupMessagesImpl) != address(0), "Implementation deployment failed");

        // Deploy the proxy contract.
        proxy = new ERC1967Proxy(
            address(groupMessagesImpl), abi.encodeWithSelector(GroupMessages.initialize.selector, admin)
        );

        vm.stopBroadcast();

        _serializeDeploymentData();
    }

    function _serializeDeploymentData() internal {
        string memory parent_object = "parent object";
        string memory addresses = "addresses";

        string memory addressesOutput;

        addressesOutput = vm.serializeAddress(addresses, "deployer", deployer);
        addressesOutput = vm.serializeAddress(addresses, "proxyAdmin", admin);
        addressesOutput = vm.serializeAddress(addresses, "proxy", address(proxy));
        addressesOutput = vm.serializeAddress(addresses, "implementation", address(groupMessagesImpl));

        string memory finalJson;
        finalJson = vm.serializeString(parent_object, addresses, addressesOutput);
        finalJson = vm.serializeUint(parent_object, "deploymentBlock", block.number);
        finalJson = vm.serializeUint(parent_object, "latestUpgradeBlock", block.number);

        writeOutput(finalJson, XMTP_GROUP_MESSAGES_OUTPUT_JSON);
    }
}
