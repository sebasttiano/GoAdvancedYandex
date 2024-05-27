package main

import (
	"context"
	"errors"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
	"golang.org/x/sync/singleflight"
)

type PostID2 string

type Post2 struct {
	ID      PostID2
	Content string
}

type PostProvider2 interface {
	GetPost(ctx context.Context, postID PostID2) (*Post2, error)
}

type BatchPostProvider2 interface {
	GetPosts(ctx context.Context, postIDs []PostID2) (map[PostID2]*Post2, error)
}

type batchPostProvider2 struct {
	provider PostProvider2
	sflgrp   *singleflight.Group // для переиспользования результатов
	sem      *semaphore.Weighted // для поддержки очереди запросов к PostProvider2
}

// конструктор
func NewbatchPostProvider2(p PostProvider2, n int64) *batchPostProvider2 {
	b := new(batchPostProvider2)
	b.provider = p
	b.sflgrp = new(singleflight.Group)
	b.sem = semaphore.NewWeighted(n) // лимит очереди запросов n
	return b
}

func (p *batchPostProvider2) GetPosts(ctx context.Context,
	postIDs []PostID2) (map[PostID2]*Post2, error) {
	grp, ctx := errgroup.WithContext(ctx)
	posts := make([]*Post2, len(postIDs))
	for i, id := range postIDs {
		index, idx := i, id
		grp.Go(func() error {
			// используем новый метод запроса с кешированием результатов
			post, err := p.SingleGetPost(ctx, idx)
			if err != nil {
				return err
			}
			posts[index] = post
			return nil
		})
	}
	if err := grp.Wait(); err != nil {
		return nil, err
	}
	return p.collectPostsByIDs(postIDs, posts), nil
}

func (p *batchPostProvider2) collectPostsByIDs(ids []PostID2,
	posts []*Post2) map[PostID2]*Post2 {
	result := make(map[PostID2]*Post2, len(ids))
	for i, id := range ids {
		result[id] = posts[i]
	}
	return result
}

// метод с той же сигнатурой, что и GetPost, но переиспользующий результаты запросов
// и поддерживающий очередь семафора
func (p *batchPostProvider2) SingleGetPost(ctx context.Context, postID PostID2) (*Post2, error) {
	// лямбда-функция в аргументе будет вызвана только один раз для каждого ключа postID
	// для передачи параметров используем замыкание — closure
	post, err, _ := p.sflgrp.Do(string(postID), func() (interface{}, error) {
		// если возможно, занимаем место в очереди семафора
		if !p.sem.TryAcquire(1) {
			// если очередь переполнена, возвращаем ошибку
			return nil, errors.New("too many requests")
		}
		pst, er := p.provider.GetPost(ctx, postID)
		// освобождаем место в очереди семафора
		p.sem.Release(1)
		return pst, er
	})
	if err != nil {
		return nil, err
	}
	// неудобство функций, возвращающих интерфейсы,
	// в том, что необходим type assertion
	return post.(*Post2), nil
}
