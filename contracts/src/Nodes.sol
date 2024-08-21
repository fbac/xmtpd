// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

/**
A NFT contract for XMTP Node Operators.

The deployer of this contract is responsible for minting NFTs and assinging them to node operators.

All nodes on the network periodically check this contract to determine which nodes they should connect to.
 */
contract Nodes is ERC721, Ownable {
    constructor() ERC721("XMTP Node Operator", "XMTP") Ownable(msg.sender) {}

    // uint16 counter so that we cannot create more than 65k IDs
    // The ERC721 standard expects the tokenID to be uint256 for standard methods unfortunately
    uint16 private _nodeIdCounter;

    // A node, as stored in the internal mapping
    struct Node {
        bytes signingKeyPub;
        string httpAddress;
        bool isHealthy;
    }

    struct NodeWithId {
        uint16 nodeId;
        Node node;
    }

    event NodeUpdated(uint256 nodeId, Node node);

    // Mapping of token ID to Node
    mapping(uint256 => Node) private _nodes;

    /**
    Mint a new node NFT and store the metadata in the smart contract
     */
    function addNode(
        address to,
        bytes calldata signingKeyPub,
        string calldata httpAddress
    ) public onlyOwner returns (uint16) {
        uint16 nodeId = _nodeIdCounter;
        _mint(to, nodeId);
        _nodes[nodeId] = Node(signingKeyPub, httpAddress, true);
        _emitNodeUpdate(nodeId);
        _nodeIdCounter++;

        return nodeId;
    }

    /**
    Override the built in transferFrom function to block NFT owners from transferring
    node ownership.

    NFT owners are only allowed to update their HTTP address and MTLS cert.
     */
    function transferFrom(
        address from,
        address to,
        uint256 tokenId
    ) public override {
        require(
            _msgSender() == owner(),
            "Only the contract owner can transfer Node ownership"
        );
        super.transferFrom(from, to, tokenId);
    }

    /**
    Allow a NFT holder to update the HTTP address of their node
     */
    function updateHttpAddress(
        uint256 tokenId,
        string calldata httpAddress
    ) public {
        require(
            _msgSender() == ownerOf(tokenId),
            "Only the owner of the Node NFT can update its http address"
        );
        _nodes[tokenId].httpAddress = httpAddress;
        _emitNodeUpdate(tokenId);
    }

    /**
    The contract owner may update the health status of the node.

    No one else is allowed to call this function.
     */
    function updateHealth(uint256 tokenId, bool isHealthy) public onlyOwner {
        // Make sure that the token exists
        _requireOwned(tokenId);
        _nodes[tokenId].isHealthy = isHealthy;
        _emitNodeUpdate(tokenId);
    }

    /**
    Get a list of healthy nodes with their ID and metadata
     */
    function healthyNodes() public view returns (NodeWithId[] memory) {
        uint16 totalNodeCount = _nodeIdCounter;
        uint256 healthyCount = 0;

        // First, count the number of healthy nodes
        for (uint256 i = 0; i < totalNodeCount; i++) {
            if (_nodeExists(i) && _nodes[i].isHealthy) {
                healthyCount++;
            }
        }

        // Create an array to store healthy nodes
        NodeWithId[] memory healthyNodesList = new NodeWithId[](healthyCount);
        uint256 currentIndex = 0;

        // Populate the array with healthy nodes
        for (uint16 i = 0; i < totalNodeCount; i++) {
            if (_nodeExists(i) && _nodes[i].isHealthy) {
                healthyNodesList[currentIndex] = NodeWithId({
                    nodeId: i,
                    node: _nodes[i]
                });
                currentIndex++;
            }
        }

        return healthyNodesList;
    }

    /**
    Get all nodes regardless of their health status
     */
    function allNodes() public view returns (NodeWithId[] memory) {
        uint16 totalNodeCount = _nodeIdCounter;
        NodeWithId[] memory allNodesList = new NodeWithId[](totalNodeCount);

        for (uint16 i = 0; i < totalNodeCount; i++) {
            allNodesList[i] = NodeWithId({nodeId: i, node: _nodes[i]});
        }

        return allNodesList;
    }

    /**
    Get a node's metadata by ID
     */
    function getNode(uint256 tokenId) public view returns (Node memory) {
        _requireOwned(tokenId);
        return _nodes[tokenId];
    }

    function _emitNodeUpdate(uint256 tokenId) private {
        emit NodeUpdated(tokenId, _nodes[tokenId]);
    }

    function _nodeExists(uint256 tokenId) private view returns (bool) {
        address owner = _ownerOf(tokenId);
        return owner != address(0);
    }
}