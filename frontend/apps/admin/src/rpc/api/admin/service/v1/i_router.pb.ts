// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.0
//   protoc               unknown
// source: admin/service/v1/i_router.proto

/* eslint-disable */
import { type Empty } from "../../../google/protobuf/empty.pb";

/** 路由元数据 */
export interface RouteMeta {
  orderNo?: number | null | undefined;
  title?: string | null | undefined;
  dynamicLevel?: number | null | undefined;
  realPath?: string | null | undefined;
  icon?: string | null | undefined;
  frameSrc?: string | null | undefined;
  transitionName?: string | null | undefined;
  affix?: boolean | null | undefined;
  carryParam?: boolean | null | undefined;
  single?: boolean | null | undefined;
  ignoreAuth?: boolean | null | undefined;
  ignoreKeepAlive?: boolean | null | undefined;
  ignoreRoute?: boolean | null | undefined;
  hideBreadcrumb?: boolean | null | undefined;
  hideChildrenInMenu?: boolean | null | undefined;
  hideTab?: boolean | null | undefined;
  hideMenu?: boolean | null | undefined;
  isLink?: boolean | null | undefined;
  hidePathForChildren?: boolean | null | undefined;
  currentActiveMenu?: string | null | undefined;
}

/** 路由项 */
export interface RouteItem {
  name?: string | null | undefined;
  alias?: string | null | undefined;
  path?: string | null | undefined;
  component?: string | null | undefined;
  redirect?: string | null | undefined;
  caseSensitive?: boolean | null | undefined;
  meta?: RouteMeta | null | undefined;
  children: RouteItem[];
}

/** 查询路由列表 - 回应 */
export interface ListRouteResponse {
  items: RouteItem[];
}

/** 查询权限码列表 - 回应 */
export interface ListPermissionCodeResponse {
  codes: string[];
}

/** 网站后台动态路由服务 */
export interface RouterService {
  /** 查询路由列表 */
  ListRoute(request: Empty): Promise<ListRouteResponse>;
  /** 查询权限码列表 */
  ListPermissionCode(request: Empty): Promise<ListPermissionCodeResponse>;
}
