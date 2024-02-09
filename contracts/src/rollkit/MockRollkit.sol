// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import {EigenDARollupUtils} from "../libraries/EigenDARollupUtils.sol";
import {IEigenDAServiceManager} from "../interfaces/IEigenDAServiceManager.sol";

struct BlobData {
    uint32 blobIndex;
    bytes blobBatchHeaderHash;
}

/**
 * @title MockRollkit
 * @author Layr Labs, Inc.
 * @notice This contract is used to emulate a rollkit rollup contract to verify the blobs.
 */
contract MockRollkit {
    IEigenDAServiceManager public eigenDAServiceManager; // EigenDASM contract
    mapping(uint64 => BlobData) private _blobDatas; // blobId => blobData
    uint64 private _latestBlobId; // latest blob id

    constructor(IEigenDAServiceManager _eigenDAServiceManager) {
        eigenDAServiceManager = _eigenDAServiceManager;
    }

    /**
     * @notice a function for submitting blobs to the rollup contract, should be passed in the correct order
     * @param blobIds the blob ids
     * @param _blobHeaders the blob headers
     * @param _blobVerificationProofs the blob verification proofs
     */
    function submitBlobs(uint64[] calldata blobIds, IEigenDAServiceManager.BlobHeader[] calldata _blobHeaders, EigenDARollupUtils.BlobVerificationProof[] calldata _blobVerificationProofs, bytes[] memory _batchHeaderHashes) external {
        // ensure the length of the blobIds, _blobHeaders and _blobVerificationProofs are the same, correct order
        require(blobIds.length == _blobHeaders.length && _blobHeaders.length == _blobVerificationProofs.length, "Length of blobIds, blobHeaders, and blobVerificationProofs must be equal.");

        for (uint i = 0; i < _blobHeaders.length; i++) {
            // verify the blob, revert if verification fails
            EigenDARollupUtils.verifyBlob(_blobHeaders[i], eigenDAServiceManager, _blobVerificationProofs[i]);

            // store the blob header
            _blobDatas[blobIds[i]] = BlobData(_blobVerificationProofs[i].blobIndex, _batchHeaderHashes[i]);
            // update the latest blob id, we ensure the blobIds are passed in the correct order
            _latestBlobId = blobIds[i];
        }
    }

    /**
     * @notice a function for getting the blob identification data
     * @param blobId the blob id
     * @return the blob header
     */
    function getBlobIdentificationData(uint64 blobId) external view returns (BlobData memory) {
        return _blobDatas[blobId];
    }

    /**
     * @notice a function for getting the latest blob id
     * @return the latest blob id
     */
    function getLatestBlobId() external view returns (uint64) {
        return _latestBlobId;
    }
}