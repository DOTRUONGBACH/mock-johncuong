// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"mock-project/grpc/customer-grpc/ent/customer"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Customer is the model entity for the Customer schema.
type Customer struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// PhoneNumber holds the value of the "phone_number" field.
	PhoneNumber string `json:"phone_number,omitempty"`
	// IdentifyNumber holds the value of the "identify_number" field.
	IdentifyNumber string `json:"identify_number,omitempty"`
	// DateOfBirth holds the value of the "date_of_birth" field.
	DateOfBirth time.Time `json:"date_of_birth,omitempty"`
	// MemberCode holds the value of the "member_code" field.
	MemberCode string `json:"member_code,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Customer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case customer.FieldID:
			values[i] = new(sql.NullInt64)
		case customer.FieldName, customer.FieldEmail, customer.FieldAddress, customer.FieldPhoneNumber, customer.FieldIdentifyNumber, customer.FieldMemberCode:
			values[i] = new(sql.NullString)
		case customer.FieldDateOfBirth, customer.FieldCreatedAt, customer.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Customer fields.
func (c *Customer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case customer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int64(value.Int64)
		case customer.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case customer.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				c.Email = value.String
			}
		case customer.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				c.Address = value.String
			}
		case customer.FieldPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number", values[i])
			} else if value.Valid {
				c.PhoneNumber = value.String
			}
		case customer.FieldIdentifyNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field identify_number", values[i])
			} else if value.Valid {
				c.IdentifyNumber = value.String
			}
		case customer.FieldDateOfBirth:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date_of_birth", values[i])
			} else if value.Valid {
				c.DateOfBirth = value.Time
			}
		case customer.FieldMemberCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field member_code", values[i])
			} else if value.Valid {
				c.MemberCode = value.String
			}
		case customer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case customer.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Customer.
// This includes values selected through modifiers, order, etc.
func (c *Customer) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// Update returns a builder for updating this Customer.
// Note that you need to call Customer.Unwrap() before calling this method if this Customer
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Customer) Update() *CustomerUpdateOne {
	return NewCustomerClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Customer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Customer) Unwrap() *Customer {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Customer is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Customer) String() string {
	var builder strings.Builder
	builder.WriteString("Customer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(c.Email)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(c.Address)
	builder.WriteString(", ")
	builder.WriteString("phone_number=")
	builder.WriteString(c.PhoneNumber)
	builder.WriteString(", ")
	builder.WriteString("identify_number=")
	builder.WriteString(c.IdentifyNumber)
	builder.WriteString(", ")
	builder.WriteString("date_of_birth=")
	builder.WriteString(c.DateOfBirth.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("member_code=")
	builder.WriteString(c.MemberCode)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Customers is a parsable slice of Customer.
type Customers []*Customer
