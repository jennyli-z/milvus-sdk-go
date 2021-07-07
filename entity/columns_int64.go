// Code generated by go generate; DO NOT EDIT
// This file is generated by go genrated at 2021-07-05 22:52:01.455014639 &#43;0800 CST m=&#43;0.001674288

//package entity defines entities used in sdk
package entity 

import "github.com/milvus-io/milvus-sdk-go/internal/proto/schema"

// columnInt64 generated columns type for Int64
type columnInt64 struct {
	name   string
	values []int64
}

func (c *columnInt64) Name() string {
	return c.name
}

func (c *columnInt64) Type() FieldType {
	return FieldTypeInt64
}

func (c *columnInt64) FieldData() *schema.FieldData {
	fd := &schema.FieldData{
		FieldName: c.name,
	}
	fd.Field = &schema.FieldData_Scalars{
		Scalars: &schema.ScalarField{
			Data: &schema.ScalarField_IntData{
				IntData: &schema.IntArray{
					Data: []int32{},
				},
			},
		},
	}
	return fd
}

func NewColumnInt64(name string, values []int64) Column {
	return &columnInt64 {
		name: name,
		values: values,
	}
}