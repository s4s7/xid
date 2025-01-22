# xid
A command-line tool for generating and decoding XIDs (globally unique identifiers).

## Installation

```bash
go install github.com/s4s7/xid@latest
```

## Usage

### Generate a single XID
```bash
xid
```

### Generate multiple XIDs
```bash
xid -n 5
```

### Generate XID with specific timestamp
```bash
xid -t "2024-01-22T15:04:05"
```

### Decode XID
```bash
xid -d bvqdjn1eu6d7ms8rb6bg
```

### Show specific parts
```bash
xid -d bvqdjn1eu6d7ms8rb6bg -time  # Show only timestamp
xid -d bvqdjn1eu6d7ms8rb6bg -m     # Show only machine ID
xid -d bvqdjn1eu6d7ms8rb6bg -p     # Show only process ID
xid -d bvqdjn1eu6d7ms8rb6bg -c     # Show only counter
```