# Todo List Application - Consensus Breaking Update

This version of the Todo List Application introduces a significant update to the task structure - the addition of a "completed" boolean field for each task. This feature enhances the application's functionality by allowing users to mark tasks as completed. However, it's important to understand the implications of this change on the blockchain's consensus.

## New Features

**Completed Field:**
  - Each task now includes a `completed` field, represented as a boolean.
  - Users can mark a task as completed, which updates the `completed` status of the task on the blockchain.

## Consensus Breaking Changes

Introducing the `completed` field to the task structure constitutes a consensus-breaking change. This means that the updated application is not backward compatible with earlier versions. Here's how it affects consensus:

1. **Transaction Structure Change**:
   - The addition of the `completed` field changes the structure of the task-related transactions. Nodes running the old software will not recognize or correctly validate transactions with this new field.

2. **Network Fork Risk**:
   - If some nodes in the network adopt this update while others do not, it could lead to a network fork. This results in two divergent chains, one following the old rules and one following the new rules.

3. **Mandatory Software Upgrade**:
   - To maintain a unified network, all nodes must upgrade to the new version of the software. Failure to do so will result in those nodes being unable to validate and agree on new blocks.

4. **Hard Fork Considerations**:
   - This change is equivalent to a hard fork in blockchain terminology. Coordinating the upgrade across all nodes is crucial to avoid disruptions in the network.

## Upgrade Instructions

To ensure seamless integration and maintain network consensus, upgrade all nodes to the new software version concurrently.
For detailed usage instructions with the new `completed` field feature, please refer to the main branch README.
