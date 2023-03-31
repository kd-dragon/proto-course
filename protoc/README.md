

<h2> How to Decode `simple.bin` </h2>

<h3> decode_raw </h3>

```bash
% cat simple.bin | protoc --decode=Simple simple.proto
```

<h3> decode (e.g. simple.proto)</h3>

```bash
% cat simple.bin | protoc --decode=simple.Simple simple.proto
```

<h3> decode (e.g. simple.proto) with writing file </h3>

```bash
% cat simple.bin | protoc --decode=simple.Simple simple.proto > simple.txt
```

<h2> How to Encode `simple.pb` </h2>

```bash
cat simple.txt | protoc --encode=simple.Simple simple.proto > simple.pb
```

<h3> Compare the difference between simple.bin and simple.pb </h3>

```bash
% diff simple.bin simple.pb
# not difference found between simple.bin and simple.pb
```