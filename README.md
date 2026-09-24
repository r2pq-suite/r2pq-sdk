> **Historical 2025 SDK prototype.** This repository belongs to an earlier
> R2PQ development lineage exploring transaction-oriented SDK interfaces
> and integration hooks. It is not the SDK or implementation of the current
> R2PQ cryptographic-assessment research system.
>
> This prototype contains placeholder cryptographic and Merkle components.
> The signing and KEM stubs contain unimplemented operations, and Merkle
> construction is unimplemented for nonempty input. Interface definitions
> and feature checkmarks do not establish complete cryptographic
> implementations.
>
> These disclosures describe this historical SDK prototype. They must not
> be generalized into claims about the implementation or capabilities of
> current R2PQ assessment tooling.
>
> Current public research scope and evidence boundaries:
> [r2pq.dev](https://r2pq.dev).

## Historical prototype documentation — 2025

# r2pq-sdk

Core SDK for the R2PQ suite: keys, PQ crypto stubs, codec registry, protocol messages, and verification hooks.

APIs compile today with placeholder implementations and provide a stable interface for all other R2PQ repositories to depend on.

---

## Purpose

- Developer-friendly high-level API  
- Unified PQ key + signature handling  
- Helpers for PQ transaction formatting  
- Integration with `r2pq-cli` and `r2pq-devnet`  
- Stable interfaces for long-term tooling & upgrades  

---

## Features (Current & Planned)

✅ Core type definitions  
✅ Placeholder PQ crypto interfaces  
✅ Serialization helpers  
✅ Codec registry  
✅ Signature + verification API  
✅ Error and logging utilities  
⬜ Integration tests  
⬜ Full PQ algorithm backends (future drop-in upgrade)  
⬜ FFI bindings (Rust/Go/Python/Node)  

---

## Folder Structure
(ci test trigger)
 
