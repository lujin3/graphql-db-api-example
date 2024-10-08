package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.53

import (
	"context"
	"example/graph/model"
)

// Firefighting is the resolver for the firefighting field.
func (r *queryResolver) Firefighting(ctx context.Context, startTs *int, endTs *int, limit *int) ([]*model.FirefightingData, error) {
	return []*model.FirefightingData{
		{
			Ts:        1725850811000,
			Building:  "E栋",
			Floor:     "F1",
			Indicator: "灭火器",
			Value:     1,
		},
		{
			Ts:        1725957547000,
			Building:  "E栋",
			Floor:     "F1",
			Indicator: "灭火器",
			Value:     2,
		},
		{
			Ts:        1726655019000,
			Building:  "E栋",
			Floor:     "F2",
			Indicator: "灭火器",
			Value:     3,
		},
		{
			Ts:        1727962437000,
			Building:  "E栋",
			Floor:     "F2",
			Indicator: "灭火器",
			Value:     3,
		},
		{
			Ts:        1727962437000,
			Building:  "E栋",
			Floor:     "F2",
			Indicator: "水枪",
			Value:     5,
		},
		{
			Ts:        1727962437000,
			Building:  "F栋",
			Floor:     "F2",
			Indicator: "灭火器",
			Value:     6,
		},
	}, nil
}

// Factory is the resolver for the factory field.
func (r *queryResolver) Factory(ctx context.Context, startTs *int, endTs *int, limit *int) ([]*model.FactoryData, error) {
	return []*model.FactoryData{
		{
			Ts:       1725850811000,
			Building: "厂房一",
			Area:     "南区",
			Value:    1,
		},
		{
			Ts:       1725957547000,
			Building: "厂房二",
			Area:     "北区",
			Value:    2,
		},
	}, nil
}

// VendorOnmake is the resolver for the vendorOnmake field.
func (r *queryResolver) VendorOnmake(ctx context.Context, supplyCode *string, materialType *string) ([]*model.VendorOnmake, error) {
	val1 := 2
	val2 := 5

	return []*model.VendorOnmake{
		{
			SupplyCode:   "123456",
			SupplyName:   "双环",
			MaterialName: "55二级太阳轮",
			MaterialType: "齿轴",
			ProcessName:  "调质",
			Num:          10,
			OpID:         &val1,
		},
		{
			SupplyCode:   "123",
			SupplyName:   "亚太",
			MaterialName: "55一级太阳轮",
			MaterialType: "轴承",
			ProcessName:  "调质",
			Num:          9,
			OpID:         &val2,
		},
		{
			SupplyCode:   "123456",
			SupplyName:   "双环",
			MaterialName: "55二级太阳轮",
			MaterialType: "齿轴",
			ProcessName:  "库存",
			Num:          1,
			OpID:         &val1,
		},
		{
			SupplyCode:   "123",
			SupplyName:   "亚太",
			MaterialName: "55一级太阳轮",
			MaterialType: "轴承",
			ProcessName:  "库存",
			Num:          1,
			OpID:         &val2,
		},
		{
			SupplyCode:   "123",
			SupplyName:   "亚太",
			MaterialName: "55一级太阳轮",
			MaterialType: "铸件",
			ProcessName:  "库存",
			Num:          1,
			OpID:         &val2,
		},
		{
			SupplyCode:   "123",
			SupplyName:   "亚太",
			MaterialName: "55一级太阳轮",
			MaterialType: "轴承",
			ProcessName:  "在制加工",
			Num:          8,
			OpID:         &val2,
		},
		{
			SupplyCode:   "123",
			SupplyName:   "亚太",
			MaterialName: "55一级太阳轮",
			MaterialType: "轴承",
			ProcessName:  "调质",
			Num:          15,
			OpID:         &val2,
		},
		{
			SupplyCode:   "123",
			SupplyName:   "亚太",
			MaterialName: "55一级太阳轮",
			MaterialType: "轴承",
			ProcessName:  "铸后",
			Num:          5,
			OpID:         &val2,
		},
	}, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
