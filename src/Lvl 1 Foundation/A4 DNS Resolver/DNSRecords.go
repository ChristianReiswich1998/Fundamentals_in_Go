package main

/*
import "fmt"

// TODO: Define the Record struct with Name (string), Type (string), Value (string), TTL (int)
type Record struct {
	Name  string
	Type  string
	Value string
	TTL   int
}


// TODO: Define the Zone struct with records (map[string]*Record)
type Zone struct {
	records map[string]*Record
}


// TODO: NewZone() *Zone - create and return a new zone with initialized map
func NewZone() *Zone {
	return &Zone{
		records: make(map[string]*Record),
	}
}


// TODO: (z *Zone) AddRecord(r *Record) - store record pointer using r.Name as key
func (z *Zone) AddRecord(r *Record) {
	z.records[r.Name] = r
}


// TODO: (z *Zone) GetRecord(name string) (*Record, bool) - look up record by name
func (z *Zone) GetRecord(name string) (*Record, bool) {
	for i := 0; i < len(z.records); i++ {
	}

	return &Record{
		Name:  "",
		Type:  "",
		Value: "",
		TTL:   0,
	}, true
}


// TODO: (z *Zone) ListRecords() []*Record - return all records
func (z *Zone) ListRecords() []*Record {

}


// TODO: (z *Zone) RemoveRecord(name string) bool - remove record, return true if found



func main() {
	zone := NewZone()
	zone.AddRecord(&Record{Name: "api.example.com", Type: "A", Value: "93.184.216.34", TTL: 300})
	zone.AddRecord(&Record{Name: "cdn.example.com", Type: "CNAME", Value: "d1234.cloudfront.net", TTL: 3600})
	fmt.Printf("Zone has %d records\n", len(zone.records))

	zone2 := NewZone()
	zone2.AddRecord(&Record{Name: "api.example.com", Type: "A", Value: "93.184.216.34", TTL: 300})
	zone2.AddRecord(&Record{Name: "cdn.example.com", Type: "CNAME", Value: "d1234.cloudfront.net", TTL: 3600})

	if r, ok := zone2.GetRecord("api.example.com"); ok {
		fmt.Printf("%s -> %s\n", r.Name, r.Value)
	}

	fmt.Printf("Records: %d\n", len(zone2.ListRecords()))
	zone2.RemoveRecord("cdn.example.com")
	fmt.Printf("After remove: %d\n", len(zone2.ListRecords()))

}
*/
