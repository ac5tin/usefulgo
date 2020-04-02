# UsefulGo

## Get Started
```Go
import (
    uf "github.com/ac5tin/usefulgo"
)
```


# Array
### ArrContains
```Go
// ArrContains check is array contains a string
containstring := uf.ArrContains(arrstring,str)
```
### ArrRMS
```Go
// ArrRMS removes string element from array (ordered)
uf.ArrRMS(&arr,str)
```

# Binary
### GetBytes
```Go
// GetBytes returns binary byte slices
byteslices,err := uf.GetBytes(x)
```

# CSV
### CSVParser
```Go
// CSVParser parses csv and runs callback function
reader := bytes.NewReader(byteslice)
r := csv.NewReader(reader)
err := uf.CSVParser(r,func(res []string,chunk int){

})
```

### CSVParserH
```Go
// CSVParserH parses csv and runs callback function
reader := bytes.NewReader(byteslice)
r := csv.NewReader(reader)
err := uf.CSVParserH(r,func(res []map[string]string,chunk int){

})
```

# Hash
### HashPassword
```Go
// HashPassword returns hashed version of password
hashed,err := uf.HashPassword(pw,cost)
```
### CheckPasswordHash
```Go
// CheckPasswordHash compares hashes between password string
cmp := uf.CheckPasswordHash(pw,hashed)
```


# Random
### RandString
```Go
// RandString generates random string
rnd := uf.RandString(5)
```



# String
### RmDash
```Go
// RmDash removes dashes from strings
newstr := uf.RmDash(oldstr)
```


# UUID
### GenUUIDV4
```Go
uuid := uf.GenUUIDV4()
```