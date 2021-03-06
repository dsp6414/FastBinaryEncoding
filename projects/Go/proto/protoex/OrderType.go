// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// Version: 1.1.0.0

package protoex

import "encoding/json"
import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = fbe.Version
var _ = proto.Version

// OrderType enum
type OrderType byte

// OrderType enum values
//noinspection GoSnakeCaseUsage
const (
    OrderType_market = OrderType(0 + 0)
    OrderType_limit = OrderType(0 + 1)
    OrderType_stop = OrderType(0 + 2)
    OrderType_stoplimit = OrderType(0 + 3)
)

// Create a new OrderType enum
func NewOrderType() *OrderType {
    return new(OrderType)
}

// Convert enum to string
func (e OrderType) String() string {
    if e == OrderType_market {
        return "market"
    }
    if e == OrderType_limit {
        return "limit"
    }
    if e == OrderType_stop {
        return "stop"
    }
    if e == OrderType_stoplimit {
        return "stoplimit"
    }
    return "<unknown>"
}

// Convert enum to JSON
func (e OrderType) MarshalJSON() ([]byte, error) {
    return json.Marshal(byte(e))
}

// Convert JSON to enum
func (e *OrderType) UnmarshalJSON(b []byte) error {
    var result byte
    err := json.Unmarshal(b, &result)
    if err != nil {
        return err
    }
    *e = OrderType(result)
    return nil
}
