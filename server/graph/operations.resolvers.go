package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"

	"github.com/minskylab/calab/experiments"
	"github.com/minskylab/calab/server/graph/generated"
	"github.com/minskylab/calab/server/graph/model"
	"github.com/pkg/errors"
)

func (r *mutationResolver) TakeSnapshot(ctx context.Context, id string) (*model.PetriDishSnapshot, error) {
	filepath, ticks := r.Exp.Snapshot(id)

	return &model.PetriDishSnapshot{
		PetriDishID: id,
		TickStamp:   int(ticks),
		Image:       filepath,
	}, nil
}

func (r *mutationResolver) Run(ctx context.Context, id string, time *string, ticks *int) (*model.PetriDish, error) {
	if time != nil {
		r.Exp.Run(id, experiments.WithTime(*time))
	} else if ticks != nil {
		r.Exp.Run(id, experiments.WithTicks(uint64(*ticks)))
	} else {
		return nil, errors.New("please choose one mode (ticks or time) to run.")
	}

	return &model.PetriDish{
		ID:        id,
		Ticks:     int(r.Exp.Ticks(id)),
		System:    &model.DynamicalSystem{ID: ""},
		Snapshots: []*model.PetriDishSnapshot{},
	}, nil
}

func (r *queryResolver) PetriDish(ctx context.Context, id string) (*model.PetriDish, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) Observer(ctx context.Context, id string) (<-chan *model.PetriDishFrame, error) {
	frameBroadcast, err := r.Exp.Observe(id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	broadcast := make(chan *model.PetriDishFrame, 1)

	go func(id string, bc chan image.Image) {
		buffer := bytes.NewBuffer([]byte{})
		for frame := range bc {
			if err := jpeg.Encode(buffer, frame, &jpeg.Options{Quality: 60}); err != nil {
				panic(err)
			}

			broadcast <- &model.PetriDishFrame{
				PetriDishID: id,
				TickStamp:   0,
				Data:        base64.StdEncoding.EncodeToString(buffer.Bytes()),
			}
		}
	}(id, frameBroadcast)

	return broadcast, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type (
	mutationResolver     struct{ *Resolver }
	queryResolver        struct{ *Resolver }
	subscriptionResolver struct{ *Resolver }
)
