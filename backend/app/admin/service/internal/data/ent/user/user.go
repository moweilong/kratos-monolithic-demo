// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateBy holds the string denoting the create_by field in the database.
	FieldCreateBy = "create_by"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldDeleteTime holds the string denoting the delete_time field in the database.
	FieldDeleteTime = "delete_time"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldRoleID holds the string denoting the role_id field in the database.
	FieldRoleID = "role_id"
	// FieldOrgID holds the string denoting the org_id field in the database.
	FieldOrgID = "org_id"
	// FieldPositionID holds the string denoting the position_id field in the database.
	FieldPositionID = "position_id"
	// FieldWorkID holds the string denoting the work_id field in the database.
	FieldWorkID = "work_id"
	// FieldNickName holds the string denoting the nick_name field in the database.
	FieldNickName = "nick_name"
	// FieldRealName holds the string denoting the real_name field in the database.
	FieldRealName = "real_name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldAuthority holds the string denoting the authority field in the database.
	FieldAuthority = "authority"
	// FieldLastLoginTime holds the string denoting the last_login_time field in the database.
	FieldLastLoginTime = "last_login_time"
	// FieldLastLoginIP holds the string denoting the last_login_ip field in the database.
	FieldLastLoginIP = "last_login_ip"
	// Table holds the table name of the user in the database.
	Table = "user"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreateBy,
	FieldCreateTime,
	FieldUpdateTime,
	FieldDeleteTime,
	FieldStatus,
	FieldUsername,
	FieldPassword,
	FieldRoleID,
	FieldOrgID,
	FieldPositionID,
	FieldWorkID,
	FieldNickName,
	FieldRealName,
	FieldEmail,
	FieldPhone,
	FieldAvatar,
	FieldGender,
	FieldAddress,
	FieldDescription,
	FieldAuthority,
	FieldLastLoginTime,
	FieldLastLoginIP,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// NickNameValidator is a validator for the "nick_name" field. It is called by the builders before save.
	NickNameValidator func(string) error
	// RealNameValidator is a validator for the "real_name" field. It is called by the builders before save.
	RealNameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultPhone holds the default value on creation for the "phone" field.
	DefaultPhone string
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// AvatarValidator is a validator for the "avatar" field. It is called by the builders before save.
	AvatarValidator func(string) error
	// DefaultAddress holds the default value on creation for the "address" field.
	DefaultAddress string
	// AddressValidator is a validator for the "address" field. It is called by the builders before save.
	AddressValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultLastLoginIP holds the default value on creation for the "last_login_ip" field.
	DefaultLastLoginIP string
	// LastLoginIPValidator is a validator for the "last_login_ip" field. It is called by the builders before save.
	LastLoginIPValidator func(string) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(uint32) error
)

// Status defines the type for the "status" enum field.
type Status string

// StatusON is the default value of the Status enum.
const DefaultStatus = StatusON

// Status values.
const (
	StatusOFF Status = "OFF"
	StatusON  Status = "ON"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusOFF, StatusON:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for status field: %q", s)
	}
}

// Gender defines the type for the "gender" enum field.
type Gender string

// Gender values.
const (
	GenderUNKNOWN Gender = "UNKNOWN"
	GenderMALE    Gender = "MALE"
	GenderFEMALE  Gender = "FEMALE"
)

func (ge Gender) String() string {
	return string(ge)
}

// GenderValidator is a validator for the "gender" field enum values. It is called by the builders before save.
func GenderValidator(ge Gender) error {
	switch ge {
	case GenderUNKNOWN, GenderMALE, GenderFEMALE:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for gender field: %q", ge)
	}
}

// Authority defines the type for the "authority" enum field.
type Authority string

// AuthorityCUSTOMER_USER is the default value of the Authority enum.
const DefaultAuthority = AuthorityCUSTOMER_USER

// Authority values.
const (
	AuthoritySYS_ADMIN     Authority = "SYS_ADMIN"
	AuthorityCUSTOMER_USER Authority = "CUSTOMER_USER"
	AuthorityGUEST_USER    Authority = "GUEST_USER"
	AuthorityREFRESH_TOKEN Authority = "REFRESH_TOKEN"
)

func (a Authority) String() string {
	return string(a)
}

// AuthorityValidator is a validator for the "authority" field enum values. It is called by the builders before save.
func AuthorityValidator(a Authority) error {
	switch a {
	case AuthoritySYS_ADMIN, AuthorityCUSTOMER_USER, AuthorityGUEST_USER, AuthorityREFRESH_TOKEN:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for authority field: %q", a)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateBy orders the results by the create_by field.
func ByCreateBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateBy, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByDeleteTime orders the results by the delete_time field.
func ByDeleteTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeleteTime, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByRoleID orders the results by the role_id field.
func ByRoleID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoleID, opts...).ToFunc()
}

// ByOrgID orders the results by the org_id field.
func ByOrgID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrgID, opts...).ToFunc()
}

// ByPositionID orders the results by the position_id field.
func ByPositionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPositionID, opts...).ToFunc()
}

// ByWorkID orders the results by the work_id field.
func ByWorkID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWorkID, opts...).ToFunc()
}

// ByNickName orders the results by the nick_name field.
func ByNickName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNickName, opts...).ToFunc()
}

// ByRealName orders the results by the real_name field.
func ByRealName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRealName, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByAvatar orders the results by the avatar field.
func ByAvatar(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatar, opts...).ToFunc()
}

// ByGender orders the results by the gender field.
func ByGender(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGender, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByAuthority orders the results by the authority field.
func ByAuthority(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthority, opts...).ToFunc()
}

// ByLastLoginTime orders the results by the last_login_time field.
func ByLastLoginTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastLoginTime, opts...).ToFunc()
}

// ByLastLoginIP orders the results by the last_login_ip field.
func ByLastLoginIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastLoginIP, opts...).ToFunc()
}