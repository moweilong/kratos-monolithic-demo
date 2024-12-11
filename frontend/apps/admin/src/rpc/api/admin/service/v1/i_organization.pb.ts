// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.0
//   protoc               unknown
// source: admin/service/v1/i_organization.proto

/* eslint-disable */
import { type Empty } from "../../../google/protobuf/empty.pb";
import { type PagingRequest } from "../../../pagination/v1/pagination.pb";
import {
  type CreateOrganizationRequest,
  type DeleteOrganizationRequest,
  type GetOrganizationRequest,
  type ListOrganizationResponse,
  type Organization,
  type UpdateOrganizationRequest,
} from "../../../user/service/v1/organization.pb";

/** 组织部门管理服务 */
export interface OrganizationService {
  /** 查询部门列表 */
  ListOrganization(request: PagingRequest): Promise<ListOrganizationResponse>;
  /** 查询部门详情 */
  GetOrganization(request: GetOrganizationRequest): Promise<Organization>;
  /** 创建部门 */
  CreateOrganization(request: CreateOrganizationRequest): Promise<Empty>;
  /** 更新部门 */
  UpdateOrganization(request: UpdateOrganizationRequest): Promise<Empty>;
  /** 删除部门 */
  DeleteOrganization(request: DeleteOrganizationRequest): Promise<Empty>;
}
