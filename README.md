# xid
A command-line tool for generating and decoding XIDs (globally unique identifiers).

## Installation

```bash
go install github.com/s4s7/xid@latest
```

## Usage

### Generate a single XID
```bash
xid-cli
```

### Generate multiple XIDs
```bash
xid-cli -n 5
```

### Generate XID with specific timestamp
```bash
xid-cli -t "2024-01-22T15:04:05"
```

### Decode XID
```bash
xid-cli -d bvqdjn1eu6d7ms8rb6bg
```

### Show specific parts
```bash
xid-cli -d bvqdjn1eu6d7ms8rb6bg -time  # Show only timestamp
xid-cli -d bvqdjn1eu6d7ms8rb6bg -m     # Show only machine ID
xid-cli -d bvqdjn1eu6d7ms8rb6bg -p     # Show only process ID
xid-cli -d bvqdjn1eu6d7ms8rb6bg -c     # Show only counter
```