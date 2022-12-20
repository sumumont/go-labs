package main

import (
	"errors"
	"reflect"
)

type Iterator struct {
	members interface{}
	page    *page
}
type page struct {
	pageNum     int //总共多少页
	currentPage int //当前页码
	total       int //总数
	pageSize    int //分页数
}

func (iterator *Iterator) Next() interface{} {
	if iterator.page.currentPage > iterator.page.pageNum-1 {
		return nil
	}
	start := iterator.page.currentPage * iterator.page.pageSize
	end := (iterator.page.currentPage + 1) * iterator.page.pageSize
	if iterator.page.currentPage == iterator.page.pageNum-1 {
		end = iterator.page.total
		l := end - start
		list := reflect.MakeSlice(reflect.TypeOf(iterator.members), l, l)
		s := reflect.ValueOf(iterator.members)
		for i := start; i < end; i++ {
			ele := s.Index(i)
			list.Index(i - start).Set(ele)
		}
		iterator.page.currentPage++
		return list.Interface()
	} else {
		l := end - start
		list := reflect.MakeSlice(reflect.TypeOf(iterator.members), l, l)
		s := reflect.ValueOf(iterator.members)
		for i := start; i < end; i++ {
			ele := s.Index(i)
			list.Index(i - start).Set(ele)
		}
		iterator.page.currentPage++
		return list.Interface()
	}
}

func IteratorNew(pageSize int, data interface{}) (*Iterator, error) {
	if pageSize < 2 {
		pageSize = 2
	}
	iterator := &Iterator{
		page: &page{
			pageNum:     0,
			currentPage: 0,
			total:       0,
			pageSize:    pageSize,
		},
	}
	kind := reflect.TypeOf(data).Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil, errors.New("this data is not array")
	}

	s := reflect.ValueOf(data)
	iterator.page.total = s.Len()
	iterator.members = data
	lastPage := iterator.page.total % iterator.page.pageSize
	iterator.page.pageNum = iterator.page.total / iterator.page.pageSize
	if lastPage > 0 {
		iterator.page.pageNum++
	}
	return iterator, nil
}
