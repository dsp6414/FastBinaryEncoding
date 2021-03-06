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

// OrderSide enum
type OrderSide byte

// OrderSide enum values
//noinspection GoSnakeCaseUsage
const (
    OrderSide_buy = OrderSide(0 + 0)
    OrderSide_sell = OrderSide(0 + 1)
    OrderSide_tell = OrderSide(0 + 2)
)

// Create a new OrderSide enum
func NewOrderSide() *OrderSide {
    return new(OrderSide)
}

// Convert enum to string
func (e OrderSide) String() string {
    if e == OrderSide_buy {
        return "buy"
    }
    if e == OrderSide_sell {
        return "sell"
    }
    if e == OrderSide_tell {
        return "tell"
    }
    return "<unknown>"
}

// Convert enum to JSON
func (e OrderSide) MarshalJSON() ([]byte, error) {
    return json.Marshal(byte(e))
}

// Convert JSON to enum
func (e *OrderSide) UnmarshalJSON(b []byte) error {
    var result byte
    err := json.Unmarshal(b, &result)
    if err != nil {
        return err
    }
    *e = OrderSide(result)
    return nil
}
