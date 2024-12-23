// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.0
//   protoc               unknown
// source: user/service/v1/user.proto

/* eslint-disable */
import { type Empty } from "../../../google/protobuf/empty.pb";
import { type Timestamp } from "../../../google/protobuf/timestamp.pb";
import { type PagingRequest } from "../../../pagination/v1/pagination.pb";

/** 验证密码结果码 */
export enum VerifyPasswordResult {
  /** SUCCESS - 验证成功 */
  SUCCESS = "SUCCESS",
  /** ACCOUNT_NOT_EXISTS - 账号不存在 */
  ACCOUNT_NOT_EXISTS = "ACCOUNT_NOT_EXISTS",
  /** WRONG_PASSWORD - 密码错误 */
  WRONG_PASSWORD = "WRONG_PASSWORD",
  /** FREEZE - 已冻结 */
  FREEZE = "FREEZE",
  /** DELETED - 已删除 */
  DELETED = "DELETED",
}

/** 用户权限 */
export enum UserAuthority {
  /** SYS_ADMIN - 系统超级用户 */
  SYS_ADMIN = "SYS_ADMIN",
  /** SYS_MANAGER - 系统管理员 */
  SYS_MANAGER = "SYS_MANAGER",
  /** CUSTOMER_USER - 普通用户 */
  CUSTOMER_USER = "CUSTOMER_USER",
  /** GUEST_USER - 游客 */
  GUEST_USER = "GUEST_USER",
  /** REFRESH_TOKEN - 刷新令牌 */
  REFRESH_TOKEN = "REFRESH_TOKEN",
}

/** 用户 */
export interface User {
  /** 用户ID */
  id: number;
  /** 角色ID */
  roleId?:
    | number
    | null
    | undefined;
  /** 工号 */
  workId?:
    | number
    | null
    | undefined;
  /** 部门ID */
  orgId?:
    | number
    | null
    | undefined;
  /** 岗位ID */
  positionId?:
    | number
    | null
    | undefined;
  /** 创建者ID */
  creatorId?:
    | number
    | null
    | undefined;
  /** 登录名 */
  userName?:
    | string
    | null
    | undefined;
  /** 昵称 */
  nickName?:
    | string
    | null
    | undefined;
  /** 真实姓名 */
  realName?:
    | string
    | null
    | undefined;
  /** 头像 */
  avatar?:
    | string
    | null
    | undefined;
  /** 邮箱 */
  email?:
    | string
    | null
    | undefined;
  /** 手机号 */
  phone?:
    | string
    | null
    | undefined;
  /** 性别 */
  gender?:
    | string
    | null
    | undefined;
  /** 住址 */
  address?:
    | string
    | null
    | undefined;
  /** 个人描述 */
  description?:
    | string
    | null
    | undefined;
  /** 最后登录时间 */
  lastLoginTime?:
    | string
    | null
    | undefined;
  /** 最后登录IP */
  lastLoginIp?:
    | string
    | null
    | undefined;
  /** 用户状态 */
  status?:
    | string
    | null
    | undefined;
  /** 权限 */
  authority?: UserAuthority | null | undefined;
  createTime?: Timestamp | null | undefined;
  updateTime?: Timestamp | null | undefined;
  deleteTime?: Timestamp | null | undefined;
}

/** 获取用户列表 - 答复 */
export interface ListUserResponse {
  items: User[];
  total: number;
}

/** 获取用户数据 - 请求 */
export interface GetUserRequest {
  id: number;
}

export interface GetUserByUserNameRequest {
  userName: string;
}

/** 创建用户 - 请求 */
export interface CreateUserRequest {
  operatorId?: number | null | undefined;
  user: User | null;
  password?: string | null | undefined;
}

/** 更新用户 - 请求 */
export interface UpdateUserRequest {
  operatorId?: number | null | undefined;
  user: User | null;
  password?: string | null | undefined;
  updateMask: string[] | null;
  allowMissing?: boolean | null | undefined;
}

/** 删除用户 - 请求 */
export interface DeleteUserRequest {
  operatorId?: number | null | undefined;
  id: number;
}

/** 验证密码 - 请求 */
export interface VerifyPasswordRequest {
  /** 用户名 */
  userName: string;
  /** 密码 */
  password: string;
}

/** 验证密码 - 答复 */
export interface VerifyPasswordResponse {
  result: VerifyPasswordResult;
}

/** 用户是否存在 - 请求 */
export interface UserExistsRequest {
  userName: string;
}

/** 用户是否存在 - 答复 */
export interface UserExistsResponse {
  exist: boolean;
}

/** 用户服务 */
export interface UserService {
  /** 查询用户列表 */
  ListUser(request: PagingRequest): Promise<ListUserResponse>;
  /** 查询用户详情 */
  GetUser(request: GetUserRequest): Promise<User>;
  /** 创建用户 */
  CreateUser(request: CreateUserRequest): Promise<Empty>;
  /** 更新用户 */
  UpdateUser(request: UpdateUserRequest): Promise<Empty>;
  /** 删除用户 */
  DeleteUser(request: DeleteUserRequest): Promise<Empty>;
  /** 查询用户详情 */
  GetUserByUserName(request: GetUserByUserNameRequest): Promise<User>;
  /** 验证密码 */
  VerifyPassword(request: VerifyPasswordRequest): Promise<VerifyPasswordResponse>;
  /** 用户是否存在 */
  UserExists(request: UserExistsRequest): Promise<UserExistsResponse>;
}
