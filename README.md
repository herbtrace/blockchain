# Agricultural Supply Chain Blockchain System

![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)
![Go Version](https://img.shields.io/badge/go-v1.17-blue.svg)
![Hyperledger](https://img.shields.io/badge/hyperledger-2.4.x-green.svg)

A comprehensive **Hyperledger Fabric-based blockchain solution** for agricultural supply chain traceability, providing transparent and immutable tracking from farm to consumer.

## 📖 Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Architecture](#architecture)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [API Reference](#api-reference)
- [Testing](#testing)
- [Security](#security)
- [Contributing](#contributing)
- [License](#license)

## Overview

This enterprise-grade blockchain solution implements end-to-end supply chain traceability using Hyperledger Fabric. The system tracks agricultural products through their entire journey, from harvest to consumer, ensuring:

- **Transparency**: Complete visibility of product movement
- **Immutability**: Tamper-proof record of all transactions
- **Traceability**: Real-time tracking and history access
- **Compliance**: Automated regulatory compliance checking
- **Trust**: Multi-party verification at each step

### Use Case Flow

1. **Farmer** records harvest details
2. **Transporter** verifies and moves product
3. **Processor** validates and processes
4. **Distributor** manages packaging
5. **Retailer** receives and sells
6.**Consumer** verifies authenticity

## Features

### Core Functionality

- **Smart Contract Operations**
  - Asset creation and management
  - Ownership transfers
  - Multi-party approvals
  - Compliance verification
  - History tracking

- **QR Code System**
  - Dynamic QR generation
  - Secure verification process
  - History access through scanning
  - Mobile-optimized interface

- **User Management**
  - Role-based access control
  - Digital identity management
  - Multi-signature support
  - Organization management

### Technical Features

- **Blockchain Network**
  - Multi-organization support
  - Chaincode lifecycle management
  - Private data collections
  - State database (CouchDB)

- **Security**
  - TLS encryption
  - CA-based authentication
  - MSP-based authorization
  - Audit logging



### Directory Structure

```
blockchain/
├── api/                    # REST API implementation
├── chaincode-go/          # Smart contracts
├── config/                # Network configuration
├── test-network/         # Test network setup
└── builders/             # Custom chaincode builders
```

## Prerequisites

### System Requirements
- Linux-based OS (Ubuntu 20.04 LTS recommended)
- 4GB RAM minimum
- 20GB available storage
- Docker v20.10.x or higher
- Docker Compose v2.0.x or higher

### Required Software
```bash
# System packages
sudo apt-get update
sudo apt-get install -y git curl wget jq

# Docker installation
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# Go installation
wget https://go.dev/dl/go1.17.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.17.linux-amd64.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
source ~/.bashrc

# Node.js
curl -fsSL https://deb.nodesource.com/setup_14.x | sudo -E bash -
sudo apt-get install -y nodejs
```

## Installation

1. **Clone Repository**
```bash
git clone https://github.com/herbtrace/blockchain-supply-chain.git
cd blockchain-supply-chain
```

2. **Network Setup**
```bash
cd test-network
./network.sh up createChannel -c supplychannel -ca
```

3. **Deploy Chaincode**
```bash
./network.sh deployCC -ccn supplychain -ccp ../chaincode-go -ccl go
```

4. **Start API Server**
```bash
cd ../api
go mod tidy
go run main.go
```

## Usage

### Basic Operations

1. **Create Product Entry**
```bash
curl -X POST http://localhost:3000/api/product/create \
  -H "Content-Type: application/json" \
  -d '{
    "cropType": "Wheat",
    "location": {"lat": 51.507351, "long": -0.127758},
    "timestamp": "2023-09-18T10:00:00Z"
  }'
```

2. **Track Product**
```bash
curl -X GET http://localhost:3000/api/product/{productId}/history
```

### QR Code Operations

1. **Generate QR**
```bash
curl -X POST http://localhost:3000/api/qr/generate/{productId}
```

2. **Verify QR**
```bash
curl -X GET http://localhost:3000/api/qr/{qrId}/verify
```

## API Reference

### Product Management
| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/product/create` | POST | Create new product |
| `/api/product/{id}` | GET | Get product details |
| `/api/product/{id}/history` | GET | Get product history |

### Event Management
| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/events/create` | POST | Record new event |
| `/api/events/{id}` | GET | Get event details |

## Testing

```bash
# Run all tests
cd chaincode-go
go test ./... -v

# Run specific test suite
go test ./chaincode/smartcontract_test.go -v
```

## Security

### Network Security
- TLS 1.3 encryption
- Mutual TLS authentication
- Private data collections
- Channel-based isolation

### Access Control
- Role-based access (RBAC)
- Organization-level MSP
- Certificate-based identity
- Policy-based endorsement

## Contributing

1. Fork repository
2. Create feature branch
3. Commit changes
4. Submit pull request


## License

[Apache License 2.0](LICENSE)


