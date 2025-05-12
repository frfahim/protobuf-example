# Protobuf Example: Go & Python

## What is Protobuf?
Protocol Buffers (Protobuf) is a language-neutral, platform-neutral, extensible mechanism for serializing structured data, developed by Google. It allows you to define data structures (messages) and services, then generate source code in various languages to efficiently encode, transmit, and decode data.

### Benefits of Protobuf
- **Compact & Fast:** Protobuf encodes data in a binary format, making it much smaller and faster to serialize/deserialize than text-based formats like JSON or XML.
- **Strongly Typed:** The schema is defined in `.proto` files, ensuring data consistency and type safety.
- **Backward/Forward Compatibility:** Fields can be added or removed without breaking deployed programs.
- **Multi-language Support:** Code can be generated for many languages (Go, Python, Java, etc.).

## Generating Protobuf Files

### For Go
```sh
protoc --go_out=go_pb employee.proto
```
For details check official tutorial https://protobuf.dev/getting-started/gotutorial/

### For Python
```sh
protoc --python_out=py_pb employee.proto
```

For details tutorial: https://protobuf.dev/getting-started/pythontutorial/

## File Size Comparison: Protobuf vs JSON
After serializing the same data:
- **employee.bin** (Protobuf binary): _smaller file size_
- **employee.json** (JSON): _larger file size_

Example output:
```
Binary file size: 32 bytes
JSON file size: 85 bytes
```


## Project Structure
- `employee.proto`: Protobuf schema definition

Golang:
- `go_pb/employee.pb.go`: Generated Go code
- `main.go`: Go example for serialization/deserialization

Python:
- `py_pb/employee_pb2.py`: Generated Python code
- `py_pb/test.py`: Python example


### Why is Protobuf instead of widely used JSON?
- **Efficiency:** Protobuf's binary format is more compact and faster to process.
- **Schema Enforcement:** Protobuf enforces a schema, reducing errors and improving maintainability.
- **Versioning:** Protobuf supports backward and forward compatibility, making it easier to evolve your data structures.
- **Cross-language:** Protobuf code generation works across many languages, while JSON is just a data format.

## Real-life Example Usages

Protobuf is used for:

- **Microservices Communication:** Protobuf is the default serialization format for gRPC, enabling fast and efficient communication between microservices in distributed systems.
- **Data Storage:** Databases and data pipelines (e.g., Google Bigtable, Apache Kafka) use Protobuf for compact, schema-based data serialization.
- **Mobile and Web APIs:** Used to transmit data between clients and servers in mobile and web applications, especially where bandwidth and speed are critical.
- **Game Development:** Online games use Protobuf to send structured data (like player state and game events) between servers and clients with minimal latency.
- **IoT Devices:** Protobuf is used in IoT systems to serialize sensor data and commands due to its compact size and speed.

These use cases benefit from Protobuf’s efficiency, strong typing, and cross-language support.

---

For more, see the [Protocol Buffers documentation](https://protobuf.dev/).
