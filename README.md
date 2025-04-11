# MCSO Secret Interface (MSI)

**MSI** defines a gRPC-based contract for integrating any secret store with the [MCSO controller](https://github.com/hasAnybodySeenHarry/secret-rotator). It allows external plugins to handle secret fetching, making MCSO extensible across cloud providers and secret management systems.

## 📦 Features

- 🔌 Pluggable gRPC interface
- 🌐 Cloud-agnostic integration with secret stores
- 📜 Simple proto definition for easy implementation
- 🧪 Compatible with any gRPC-compatible language

