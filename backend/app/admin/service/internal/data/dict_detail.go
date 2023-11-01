package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/tx7do/go-utils/entgo/query"
	util "github.com/tx7do/go-utils/time"

	"kratos-monolithic-demo/app/admin/service/internal/data/ent"
	"kratos-monolithic-demo/app/admin/service/internal/data/ent/dictdetail"

	"kratos-monolithic-demo/gen/api/go/common/pagination"
	"kratos-monolithic-demo/gen/api/go/system/service/v1"
)

type DictDetailRepo struct {
	data *Data
	log  *log.Helper
}

func NewDictDetailRepo(data *Data, logger log.Logger) *DictDetailRepo {
	l := log.NewHelper(log.With(logger, "module", "dict-detail/repo/admin-service"))
	return &DictDetailRepo{
		data: data,
		log:  l,
	}
}

func (r *DictDetailRepo) convertEntToProto(in *ent.DictDetail) *v1.DictDetail {
	if in == nil {
		return nil
	}
	return &v1.DictDetail{
		Id:         in.ID,
		DictId:     in.DictID,
		OrderNo:    in.OrderNo,
		Label:      in.Label,
		Value:      in.Value,
		CreatorId:  in.CreateBy,
		CreateTime: util.TimeToTimeString(in.CreateTime),
		UpdateTime: util.TimeToTimeString(in.UpdateTime),
		DeleteTime: util.TimeToTimeString(in.DeleteTime),
	}
}

func (r *DictDetailRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().DictDetail.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *DictDetailRepo) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListDictDetailResponse, error) {
	builder := r.data.db.Client().DictDetail.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(r.data.db.Driver().Dialect(),
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), dictdetail.FieldCreateTime)
	if err != nil {
		r.log.Errorf("解析SELECT条件发生错误[%s]", err.Error())
		return nil, err
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, err
	}

	items := make([]*v1.DictDetail, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &v1.ListDictDetailResponse{
		Total: int32(count),
		Items: items,
	}, nil
}

func (r *DictDetailRepo) Get(ctx context.Context, req *v1.GetDictDetailRequest) (*v1.DictDetail, error) {
	ret, err := r.data.db.Client().DictDetail.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		r.log.Errorf("query one data failed: %s", err.Error())
		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *DictDetailRepo) Create(ctx context.Context, req *v1.CreateDictDetailRequest) (*v1.DictDetail, error) {
	ret, err := r.data.db.Client().DictDetail.Create().
		SetNillableDictID(req.Detail.DictId).
		SetNillableOrderNo(req.Detail.OrderNo).
		SetNillableLabel(req.Detail.Label).
		SetNillableValue(req.Detail.Value).
		SetCreateBy(req.GetOperatorId()).
		SetCreateTime(time.Now()).
		Save(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *DictDetailRepo) Update(ctx context.Context, req *v1.UpdateDictDetailRequest) (*v1.DictDetail, error) {
	builder := r.data.db.Client().DictDetail.UpdateOneID(req.Id).
		SetNillableDictID(req.Detail.DictId).
		SetNillableOrderNo(req.Detail.OrderNo).
		SetNillableLabel(req.Detail.Label).
		SetNillableValue(req.Detail.Value).
		SetUpdateTime(time.Now())

	ret, err := builder.Save(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *DictDetailRepo) Delete(ctx context.Context, req *v1.DeleteDictDetailRequest) (bool, error) {
	err := r.data.db.Client().DictDetail.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("delete one data failed: %s", err.Error())
	}

	return err == nil, err
}