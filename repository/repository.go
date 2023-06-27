package repository

import (
	"context"
	"fmt"
	"github.com/jun98427/linkedlist/ent"
	"github.com/jun98427/linkedlist/ent/node"
)

type Repository interface {
	AddNode(ctx context.Context, value int, prevID int) error
	DeleteNode(ctx context.Context, id int) error

	PrintList(ctx context.Context) (int, error)
}

type repository struct {
	client *ent.Client
}

func New(client *ent.Client) Repository {
	return &repository{
		client: client,
	}
}

func (r *repository) AddNode(ctx context.Context, value int, prevID int) error {
	var prev, next *ent.Node
	var err error
	if prevID != 0 {
		prev, err = r.client.Node.Query().Where(node.IDEQ(prevID)).WithNext().Only(ctx)
		if err != nil {
			return err
		}
		next = prev.Edges.Next
		if err := prev.Update().ClearNext().Exec(ctx); err != nil {
			return err
		}
	} else {
		next, err = r.client.Node.Query().Where(node.Not(node.HasPrev())).Only(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return err
		}
	}

	query := r.client.Node.Create().SetValue(value)
	if prev != nil {
		query = query.SetPrev(prev)
	}

	if next != nil {
		query = query.SetNext(next)
	}

	_, err = query.Save(ctx)
	return err
}

func (r *repository) DeleteNode(ctx context.Context, id int) error {
	target, err := r.client.Node.Query().Where(node.IDEQ(id)).WithNext().WithPrev().Only(ctx)
	if err != nil {
		return err
	}

	if err := r.client.Node.DeleteOne(target).Exec(ctx); err != nil {
		return err
	}

	if target.Edges.Prev != nil && target.Edges.Next != nil {
		if err := target.Edges.Prev.Update().SetNext(target.Edges.Next).Exec(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (r *repository) PrintList(ctx context.Context) (int, error) {
	head, err := r.client.Node.Query().Where(node.Not(node.HasPrev())).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return 0, err
	}

	if head == nil {
		fmt.Println("empty list")
		return 0, nil
	}

	var count int
	for curr := head; curr != nil; curr = curr.QueryNext().FirstX(ctx) {
		fmt.Printf("{id: %d val :%d} ", curr.ID, curr.Value)
		count++
	}

	fmt.Println()
	return count, nil
}
