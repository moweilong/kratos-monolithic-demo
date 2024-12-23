// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// MenuColumns holds the columns for the "menu" table.
	MenuColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true, Comment: "id"},
		{Name: "status", Type: field.TypeEnum, Nullable: true, Comment: "状态", Enums: []string{"OFF", "ON"}, Default: "ON"},
		{Name: "create_time", Type: field.TypeTime, Nullable: true, Comment: "创建时间"},
		{Name: "update_time", Type: field.TypeTime, Nullable: true, Comment: "更新时间"},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true, Comment: "删除时间"},
		{Name: "create_by", Type: field.TypeUint32, Nullable: true, Comment: "创建者ID"},
		{Name: "order_no", Type: field.TypeInt32, Nullable: true, Comment: "排序ID", Default: 0},
		{Name: "name", Type: field.TypeString, Nullable: true, Size: 32, Comment: "菜单名称", Default: ""},
		{Name: "title", Type: field.TypeString, Nullable: true, Comment: "菜单标题", Default: ""},
		{Name: "type", Type: field.TypeEnum, Nullable: true, Comment: "菜单类型 FOLDER: 目录 MENU: 菜单 BUTTON: 按钮", Enums: []string{"FOLDER", "MENU", "BUTTON"}},
		{Name: "path", Type: field.TypeString, Nullable: true, Comment: "路径,当其类型为'按钮'的时候对应的数据操作名,例如:/user.service.v1.UserService/Login", Default: ""},
		{Name: "component", Type: field.TypeString, Nullable: true, Comment: "前端页面组件", Default: ""},
		{Name: "icon", Type: field.TypeString, Nullable: true, Size: 128, Comment: "图标", Default: ""},
		{Name: "is_ext", Type: field.TypeBool, Nullable: true, Comment: "是否外链", Default: false},
		{Name: "ext_url", Type: field.TypeString, Nullable: true, Size: 255, Comment: "外链地址"},
		{Name: "permissions", Type: field.TypeJSON, Nullable: true, Comment: "权限代码 例如:sys:menu", SchemaType: map[string]string{"mysql": "json", "postgres": "jsonb"}},
		{Name: "redirect", Type: field.TypeString, Nullable: true, Comment: "跳转路径"},
		{Name: "current_active_menu", Type: field.TypeString, Nullable: true, Comment: "当前激活路径"},
		{Name: "keep_alive", Type: field.TypeBool, Nullable: true, Comment: "是否缓存", Default: false},
		{Name: "show", Type: field.TypeBool, Nullable: true, Comment: "是否显示", Default: true},
		{Name: "hide_tab", Type: field.TypeBool, Nullable: true, Comment: "是否显示在标签页导航", Default: true},
		{Name: "hide_menu", Type: field.TypeBool, Nullable: true, Comment: "是否显示在菜单导航", Default: true},
		{Name: "hide_breadcrumb", Type: field.TypeBool, Nullable: true, Comment: "是否显示在面包屑导航", Default: true},
		{Name: "parent_id", Type: field.TypeInt32, Nullable: true, Comment: "上一层菜单ID"},
	}
	// MenuTable holds the schema information for the "menu" table.
	MenuTable = &schema.Table{
		Name:       "menu",
		Columns:    MenuColumns,
		PrimaryKey: []*schema.Column{MenuColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "menu_menu_children",
				Columns:    []*schema.Column{MenuColumns[23]},
				RefColumns: []*schema.Column{MenuColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OrganizationColumns holds the columns for the "organization" table.
	OrganizationColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "create_time", Type: field.TypeTime, Nullable: true, Comment: "创建时间"},
		{Name: "update_time", Type: field.TypeTime, Nullable: true, Comment: "更新时间"},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true, Comment: "删除时间"},
		{Name: "status", Type: field.TypeEnum, Nullable: true, Comment: "状态", Enums: []string{"OFF", "ON"}, Default: "ON"},
		{Name: "create_by", Type: field.TypeUint32, Nullable: true, Comment: "创建者ID"},
		{Name: "remark", Type: field.TypeString, Nullable: true, Comment: "备注", Default: ""},
		{Name: "name", Type: field.TypeString, Nullable: true, Size: 128, Comment: "名字", Default: ""},
		{Name: "order_no", Type: field.TypeInt32, Nullable: true, Comment: "排序ID", Default: 0},
		{Name: "parent_id", Type: field.TypeUint32, Nullable: true, Comment: "上一层部门ID", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
	}
	// OrganizationTable holds the schema information for the "organization" table.
	OrganizationTable = &schema.Table{
		Name:       "organization",
		Columns:    OrganizationColumns,
		PrimaryKey: []*schema.Column{OrganizationColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "organization_organization_children",
				Columns:    []*schema.Column{OrganizationColumns[9]},
				RefColumns: []*schema.Column{OrganizationColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "organization_id",
				Unique:  false,
				Columns: []*schema.Column{OrganizationColumns[0]},
			},
		},
	}
	// PositionColumns holds the columns for the "position" table.
	PositionColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "create_time", Type: field.TypeTime, Nullable: true, Comment: "创建时间"},
		{Name: "update_time", Type: field.TypeTime, Nullable: true, Comment: "更新时间"},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true, Comment: "删除时间"},
		{Name: "status", Type: field.TypeEnum, Nullable: true, Comment: "状态", Enums: []string{"OFF", "ON"}, Default: "ON"},
		{Name: "create_by", Type: field.TypeUint32, Nullable: true, Comment: "创建者ID"},
		{Name: "remark", Type: field.TypeString, Nullable: true, Comment: "备注", Default: ""},
		{Name: "name", Type: field.TypeString, Size: 128, Comment: "角色名称", Default: ""},
		{Name: "code", Type: field.TypeString, Size: 128, Comment: "角色标识", Default: ""},
		{Name: "order_no", Type: field.TypeInt32, Comment: "排序ID", Default: 0},
		{Name: "parent_id", Type: field.TypeUint32, Nullable: true, Comment: "上一层角色ID", Default: 0, SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
	}
	// PositionTable holds the schema information for the "position" table.
	PositionTable = &schema.Table{
		Name:       "position",
		Columns:    PositionColumns,
		PrimaryKey: []*schema.Column{PositionColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "position_position_children",
				Columns:    []*schema.Column{PositionColumns[10]},
				RefColumns: []*schema.Column{PositionColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "position_id",
				Unique:  false,
				Columns: []*schema.Column{PositionColumns[0]},
			},
		},
	}
	// RoleColumns holds the columns for the "role" table.
	RoleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "create_time", Type: field.TypeTime, Nullable: true, Comment: "创建时间"},
		{Name: "update_time", Type: field.TypeTime, Nullable: true, Comment: "更新时间"},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true, Comment: "删除时间"},
		{Name: "status", Type: field.TypeEnum, Nullable: true, Comment: "状态", Enums: []string{"OFF", "ON"}, Default: "ON"},
		{Name: "create_by", Type: field.TypeUint32, Nullable: true, Comment: "创建者ID"},
		{Name: "remark", Type: field.TypeString, Nullable: true, Comment: "备注", Default: ""},
		{Name: "name", Type: field.TypeString, Unique: true, Nullable: true, Size: 128, Comment: "角色名称"},
		{Name: "code", Type: field.TypeString, Nullable: true, Size: 128, Comment: "角色标识", Default: ""},
		{Name: "order_no", Type: field.TypeInt32, Nullable: true, Comment: "排序ID", Default: 0},
		{Name: "parent_id", Type: field.TypeUint32, Nullable: true, Comment: "上一层角色ID", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
	}
	// RoleTable holds the schema information for the "role" table.
	RoleTable = &schema.Table{
		Name:       "role",
		Columns:    RoleColumns,
		PrimaryKey: []*schema.Column{RoleColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_role_children",
				Columns:    []*schema.Column{RoleColumns[10]},
				RefColumns: []*schema.Column{RoleColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "role_id",
				Unique:  false,
				Columns: []*schema.Column{RoleColumns[0]},
			},
		},
	}
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, Comment: "id", SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "create_by", Type: field.TypeUint32, Nullable: true, Comment: "创建者ID"},
		{Name: "create_time", Type: field.TypeTime, Nullable: true, Comment: "创建时间"},
		{Name: "update_time", Type: field.TypeTime, Nullable: true, Comment: "更新时间"},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true, Comment: "删除时间"},
		{Name: "status", Type: field.TypeEnum, Nullable: true, Comment: "状态", Enums: []string{"OFF", "ON"}, Default: "ON"},
		{Name: "username", Type: field.TypeString, Unique: true, Nullable: true, Size: 50, Comment: "用户名"},
		{Name: "password", Type: field.TypeString, Nullable: true, Size: 255, Comment: "登录密码"},
		{Name: "role_id", Type: field.TypeUint32, Nullable: true, Comment: "角色ID"},
		{Name: "org_id", Type: field.TypeUint32, Nullable: true, Comment: "部门ID"},
		{Name: "position_id", Type: field.TypeUint32, Nullable: true, Comment: "职位ID"},
		{Name: "work_id", Type: field.TypeUint32, Nullable: true, Comment: "员工工号"},
		{Name: "nick_name", Type: field.TypeString, Nullable: true, Size: 128, Comment: "昵称"},
		{Name: "real_name", Type: field.TypeString, Nullable: true, Size: 128, Comment: "真实名字"},
		{Name: "email", Type: field.TypeString, Nullable: true, Size: 127, Comment: "电子邮箱"},
		{Name: "phone", Type: field.TypeString, Nullable: true, Size: 11, Comment: "手机号码", Default: ""},
		{Name: "avatar", Type: field.TypeString, Nullable: true, Size: 1023, Comment: "头像"},
		{Name: "gender", Type: field.TypeEnum, Nullable: true, Comment: "性别", Enums: []string{"UNKNOWN", "MALE", "FEMALE"}},
		{Name: "address", Type: field.TypeString, Nullable: true, Size: 2048, Comment: "地址", Default: ""},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 1023, Comment: "个人说明"},
		{Name: "authority", Type: field.TypeEnum, Nullable: true, Comment: "授权", Enums: []string{"SYS_ADMIN", "CUSTOMER_USER", "GUEST_USER", "REFRESH_TOKEN"}, Default: "CUSTOMER_USER"},
		{Name: "last_login_time", Type: field.TypeInt64, Nullable: true, Comment: "最后一次登录的时间"},
		{Name: "last_login_ip", Type: field.TypeString, Nullable: true, Size: 64, Comment: "最后一次登录的IP", Default: ""},
	}
	// UserTable holds the schema information for the "user" table.
	UserTable = &schema.Table{
		Name:       "user",
		Columns:    UserColumns,
		PrimaryKey: []*schema.Column{UserColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_id",
				Unique:  false,
				Columns: []*schema.Column{UserColumns[0]},
			},
			{
				Name:    "user_id_username",
				Unique:  true,
				Columns: []*schema.Column{UserColumns[0], UserColumns[6]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MenuTable,
		OrganizationTable,
		PositionTable,
		RoleTable,
		UserTable,
	}
)

func init() {
	MenuTable.ForeignKeys[0].RefTable = MenuTable
	MenuTable.Annotation = &entsql.Annotation{
		Table:     "menu",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	OrganizationTable.ForeignKeys[0].RefTable = OrganizationTable
	OrganizationTable.Annotation = &entsql.Annotation{
		Table:     "organization",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	PositionTable.ForeignKeys[0].RefTable = PositionTable
	PositionTable.Annotation = &entsql.Annotation{
		Table:     "position",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	RoleTable.ForeignKeys[0].RefTable = RoleTable
	RoleTable.Annotation = &entsql.Annotation{
		Table:     "role",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	UserTable.Annotation = &entsql.Annotation{
		Table:     "user",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
}
