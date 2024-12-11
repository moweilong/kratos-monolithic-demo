// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kratos-monolithic-demo/app/admin/service/internal/data/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	// id
	ID uint32 `json:"id,omitempty"`
	// 创建者ID
	CreateBy *uint32 `json:"create_by,omitempty"`
	// 创建时间
	CreateTime *time.Time `json:"create_time,omitempty"`
	// 更新时间
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// 删除时间
	DeleteTime *time.Time `json:"delete_time,omitempty"`
	// 状态
	Status *user.Status `json:"status,omitempty"`
	// 用户名
	Username *string `json:"username,omitempty"`
	// 登录密码
	Password *string `json:"password,omitempty"`
	// 角色ID
	RoleID *uint32 `json:"role_id,omitempty"`
	// 部门ID
	OrgID *uint32 `json:"org_id,omitempty"`
	// 职位ID
	PositionID *uint32 `json:"position_id,omitempty"`
	// 员工工号
	WorkID *uint32 `json:"work_id,omitempty"`
	// 昵称
	NickName *string `json:"nick_name,omitempty"`
	// 真实名字
	RealName *string `json:"real_name,omitempty"`
	// 电子邮箱
	Email *string `json:"email,omitempty"`
	// 手机号码
	Phone *string `json:"phone,omitempty"`
	// 头像
	Avatar *string `json:"avatar,omitempty"`
	// 性别
	Gender *user.Gender `json:"gender,omitempty"`
	// 地址
	Address *string `json:"address,omitempty"`
	// 个人说明
	Description *string `json:"description,omitempty"`
	// 授权
	Authority *user.Authority `json:"authority,omitempty"`
	// 最后一次登录的时间
	LastLoginTime *int64 `json:"last_login_time,omitempty"`
	// 最后一次登录的IP
	LastLoginIP  *string `json:"last_login_ip,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID, user.FieldCreateBy, user.FieldRoleID, user.FieldOrgID, user.FieldPositionID, user.FieldWorkID, user.FieldLastLoginTime:
			values[i] = new(sql.NullInt64)
		case user.FieldStatus, user.FieldUsername, user.FieldPassword, user.FieldNickName, user.FieldRealName, user.FieldEmail, user.FieldPhone, user.FieldAvatar, user.FieldGender, user.FieldAddress, user.FieldDescription, user.FieldAuthority, user.FieldLastLoginIP:
			values[i] = new(sql.NullString)
		case user.FieldCreateTime, user.FieldUpdateTime, user.FieldDeleteTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = uint32(value.Int64)
		case user.FieldCreateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_by", values[i])
			} else if value.Valid {
				u.CreateBy = new(uint32)
				*u.CreateBy = uint32(value.Int64)
			}
		case user.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				u.CreateTime = new(time.Time)
				*u.CreateTime = value.Time
			}
		case user.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				u.UpdateTime = new(time.Time)
				*u.UpdateTime = value.Time
			}
		case user.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				u.DeleteTime = new(time.Time)
				*u.DeleteTime = value.Time
			}
		case user.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				u.Status = new(user.Status)
				*u.Status = user.Status(value.String)
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = new(string)
				*u.Username = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = new(string)
				*u.Password = value.String
			}
		case user.FieldRoleID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field role_id", values[i])
			} else if value.Valid {
				u.RoleID = new(uint32)
				*u.RoleID = uint32(value.Int64)
			}
		case user.FieldOrgID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field org_id", values[i])
			} else if value.Valid {
				u.OrgID = new(uint32)
				*u.OrgID = uint32(value.Int64)
			}
		case user.FieldPositionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field position_id", values[i])
			} else if value.Valid {
				u.PositionID = new(uint32)
				*u.PositionID = uint32(value.Int64)
			}
		case user.FieldWorkID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field work_id", values[i])
			} else if value.Valid {
				u.WorkID = new(uint32)
				*u.WorkID = uint32(value.Int64)
			}
		case user.FieldNickName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nick_name", values[i])
			} else if value.Valid {
				u.NickName = new(string)
				*u.NickName = value.String
			}
		case user.FieldRealName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field real_name", values[i])
			} else if value.Valid {
				u.RealName = new(string)
				*u.RealName = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = new(string)
				*u.Email = value.String
			}
		case user.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				u.Phone = new(string)
				*u.Phone = value.String
			}
		case user.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				u.Avatar = new(string)
				*u.Avatar = value.String
			}
		case user.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				u.Gender = new(user.Gender)
				*u.Gender = user.Gender(value.String)
			}
		case user.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				u.Address = new(string)
				*u.Address = value.String
			}
		case user.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				u.Description = new(string)
				*u.Description = value.String
			}
		case user.FieldAuthority:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field authority", values[i])
			} else if value.Valid {
				u.Authority = new(user.Authority)
				*u.Authority = user.Authority(value.String)
			}
		case user.FieldLastLoginTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field last_login_time", values[i])
			} else if value.Valid {
				u.LastLoginTime = new(int64)
				*u.LastLoginTime = value.Int64
			}
		case user.FieldLastLoginIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_login_ip", values[i])
			} else if value.Valid {
				u.LastLoginIP = new(string)
				*u.LastLoginIP = value.String
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	if v := u.CreateBy; v != nil {
		builder.WriteString("create_by=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := u.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := u.UpdateTime; v != nil {
		builder.WriteString("update_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := u.DeleteTime; v != nil {
		builder.WriteString("delete_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := u.Status; v != nil {
		builder.WriteString("status=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := u.Username; v != nil {
		builder.WriteString("username=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Password; v != nil {
		builder.WriteString("password=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.RoleID; v != nil {
		builder.WriteString("role_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := u.OrgID; v != nil {
		builder.WriteString("org_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := u.PositionID; v != nil {
		builder.WriteString("position_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := u.WorkID; v != nil {
		builder.WriteString("work_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := u.NickName; v != nil {
		builder.WriteString("nick_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.RealName; v != nil {
		builder.WriteString("real_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Email; v != nil {
		builder.WriteString("email=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Phone; v != nil {
		builder.WriteString("phone=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Avatar; v != nil {
		builder.WriteString("avatar=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Gender; v != nil {
		builder.WriteString("gender=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := u.Address; v != nil {
		builder.WriteString("address=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Authority; v != nil {
		builder.WriteString("authority=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := u.LastLoginTime; v != nil {
		builder.WriteString("last_login_time=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := u.LastLoginIP; v != nil {
		builder.WriteString("last_login_ip=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
