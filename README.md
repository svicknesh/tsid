## Golang Time-Sorted Unique Identifiers (TSID) library

An implementation for generating TSID using bit manipulation. 


### Usage

#### Creating default TSID without node ID

```golang
ts := tsid.NewDefault() // creates new instance of TSID without a node identifier
tsidValue := ts.Generate() // generates the TSID and returns a `TSIDValue` struct
fmt.Println(tsidValue.Number) // `Number` is the generated TSID
fmt.Println(tsidValue.String()) // prints the number in Crockford Base32 format
```

#### Creating TSID with node ID

```golang
// create new instance of node, with 6 bits being used for the node identifier and 63 being the actual node identifier
node, err := tsid.NewNode(6, 63)
if nil != err {
    // an error will be returned if the node bitsize exceeds the maximum size and if the node identifier does not fit the bit size
    log.Println(err)
    os.Exit(1)
}

ts := tsid.New(node) // creates new instance of TSID with a node identifier
tsidValue := ts.Generate() // generates the TSID and returns a `TSIDValue` struct
fmt.Println(tsidValue.Number) // `Number` is the generated TSID
fmt.Println(tsidValue.String()) // prints the number in Crockford Base32 format
```

#### Parsing TSID from string

```golang
tsidValue, err := tsid.NewFromString("0GV12PBB3YHH2")
if nil != err {
    // returns error if there the string length doesn't match
    log.Println(err)
    os.Exit(1)
}

fmt.Println(tsidValue.Number) // prints the TSID number
```

#### Parsing TSID from number

```golang
tsidValue := tsid.NewFromNumber(606898201282168354)
fmt.Println(tsidValue.Number) // prints the TSID number
```

#### Extracting details from TSID number

```golang
tsidValue, err := tsid.NewFromString("0GV12PBB3YHH2")
if nil != err {
    // returns error if there the string length doesn't match
    log.Println(err)
    os.Exit(1)
}

tsidValue.Parse(6) // this parsess the TSID number into its various parts, the value 6 refers to the number of bits used to define the node id, this must match what was used to create the TSID in the first place
fmt.Println(tsidValue.Number, tsidValue.Time, tsidValue.NodeID, tsidValue.Counter) // prints the TSID number, it has the breakdown of the time, node id and counter
```



### Explanation

TSID is a means of creating an incremental number to be used for database primary keys or unique numbers in a manageable pattern without revealing details about the application. It is nearly impossible to predict how many TSIDs have been generated and what the consecutive TSID will be in a meaningful manner.

The structure of TSID is as follows

```
                                            adjustable
                                           <---------->
|------------------------------------------|----------|------------|
       time (msecs since 2020-01-01)           node       counter
                42 bits                       10 bits     12 bits

- time:    2^42 = ~69 years or ~139 years (with adjustable epoch)
- node:    2^10 = 1,024 (with adjustable bits)
- counter: 2^12 = 4,096 (initially random)

Notes:
The node is adjustable from 0 to 20 bits.
The node bits affect the counter bits.
The time component can be used for ~69 years if stored in a **SIGNED** 64 bits integer field.
The time component can be used for ~139 years if stored in a **UNSIGNED** 64 bits integer field.
```



### Reading materials

[Original TSID note](https://github.com/f4b6a3/tsid-creator/wiki)
[TSIDs strike the perfect balance between integers and UUIDs for most databases](https://www.foxhound.systems/blog/time-sorted-unique-identifiers/)
[The best UUID type for a database Primary Key](https://vladmihalcea.com/uuid-database-primary-key/)
[Crockford Base32](https://www.crockford.com/base32.html)
