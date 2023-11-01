// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kratos-monolithic-demo/app/admin/service/internal/data/ent/dictdetail"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// DictDetail is the model entity for the DictDetail schema.
type DictDetail struct {
	config `json:"-"`
	// ID of the ent.
	// id
	ID uint32 `json:"id,omitempty"`
	// 创建时间
	CreateTime *time.Time `json:"create_time,omitempty"`
	// 更新时间
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// 删除时间
	DeleteTime *time.Time `json:"delete_time,omitempty"`
	// 创建者ID
	CreateBy *uint32 `json:"create_by,omitempty"`
	// 字典ID
	DictID *uint32 `json:"dict_id,omitempty"`
	// 排序ID
	OrderNo *int32 `json:"order_no,omitempty"`
	// 字典标签
	Label *string `json:"label,omitempty"`
	// 字典值
	Value        *string `json:"value,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DictDetail) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case dictdetail.FieldID, dictdetail.FieldCreateBy, dictdetail.FieldDictID, dictdetail.FieldOrderNo:
			values[i] = new(sql.NullInt64)
		case dictdetail.FieldLabel, dictdetail.FieldValue:
			values[i] = new(sql.NullString)
		case dictdetail.FieldCreateTime, dictdetail.FieldUpdateTime, dictdetail.FieldDeleteTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DictDetail fields.
func (dd *DictDetail) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dictdetail.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dd.ID = uint32(value.Int64)
		case dictdetail.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				dd.CreateTime = new(time.Time)
				*dd.CreateTime = value.Time
			}
		case dictdetail.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				dd.UpdateTime = new(time.Time)
				*dd.UpdateTime = value.Time
			}
		case dictdetail.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				dd.DeleteTime = new(time.Time)
				*dd.DeleteTime = value.Time
			}
		case dictdetail.FieldCreateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_by", values[i])
			} else if value.Valid {
				dd.CreateBy = new(uint32)
				*dd.CreateBy = uint32(value.Int64)
			}
		case dictdetail.FieldDictID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field dict_id", values[i])
			} else if value.Valid {
				dd.DictID = new(uint32)
				*dd.DictID = uint32(value.Int64)
			}
		case dictdetail.FieldOrderNo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_no", values[i])
			} else if value.Valid {
				dd.OrderNo = new(int32)
				*dd.OrderNo = int32(value.Int64)
			}
		case dictdetail.FieldLabel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field label", values[i])
			} else if value.Valid {
				dd.Label = new(string)
				*dd.Label = value.String
			}
		case dictdetail.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				dd.Value = new(string)
				*dd.Value = value.String
			}
		default:
			dd.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the DictDetail.
// This includes values selected through modifiers, order, etc.
func (dd *DictDetail) GetValue(name string) (ent.Value, error) {
	return dd.selectValues.Get(name)
}

// Update returns a builder for updating this DictDetail.
// Note that you need to call DictDetail.Unwrap() before calling this method if this DictDetail
// was returned from a transaction, and the transaction was committed or rolled back.
func (dd *DictDetail) Update() *DictDetailUpdateOne {
	return NewDictDetailClient(dd.config).UpdateOne(dd)
}

// Unwrap unwraps the DictDetail entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dd *DictDetail) Unwrap() *DictDetail {
	_tx, ok := dd.config.driver.(*txDriver)
	if !ok {
		panic("ent: DictDetail is not a transactional entity")
	}
	dd.config.driver = _tx.drv
	return dd
}

// String implements the fmt.Stringer.
func (dd *DictDetail) String() string {
	var builder strings.Builder
	builder.WriteString("DictDetail(")
	builder.WriteString(fmt.Sprintf("id=%v, ", dd.ID))
	if v := dd.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := dd.UpdateTime; v != nil {
		builder.WriteString("update_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := dd.DeleteTime; v != nil {
		builder.WriteString("delete_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := dd.CreateBy; v != nil {
		builder.WriteString("create_by=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := dd.DictID; v != nil {
		builder.WriteString("dict_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := dd.OrderNo; v != nil {
		builder.WriteString("order_no=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := dd.Label; v != nil {
		builder.WriteString("label=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := dd.Value; v != nil {
		builder.WriteString("value=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// DictDetails is a parsable slice of DictDetail.
type DictDetails []*DictDetail
