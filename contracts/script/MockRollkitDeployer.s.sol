// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "forge-std/Script.sol";
import "../src/rollkit/MockRollkit.sol";
import {IEigenDAServiceManager} from "../src/interfaces/IEigenDAServiceManager.sol";

contract MockRollkitDeployer is Script {
    MockRollkit public mockRollkit;

    // forge script script/MockRollupDeployer.s.sol:MockRollupDeployer --sig "run(address, bytes32, uint256)" <DASM address> <security hash> <stake> --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast -vvvv
    // <security hash> = keccak256(abi.encode(blobHeader.quorumBlobParams))
    function run(address _eigenDAServiceManager) external {
        vm.startBroadcast();

        mockRollkit = new MockRollkit(IEigenDAServiceManager(_eigenDAServiceManager));

        string memory output = "eigenDA mock rollkit deployment output";
        vm.serializeAddress(output, "mockRollkit", address(mockRollkit));

        string memory finalJson = vm.serializeString(output, "object", output);

        vm.createDir("./script/output", true);
        vm.writeJson(finalJson, "./script/output/mock_rollkit_deploy_output.json");
        vm.stopBroadcast();
    }
}
