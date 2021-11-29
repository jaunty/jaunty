package modelsx

import (
	"context"

	"github.com/jaunty/jaunty/internal/database/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Requests fetches every request for a user.
func Requests(ctx context.Context, exec boil.ContextTransactor, sf string) (models.WhitelistSlice, error) {
	w, err := models.Whitelists(
		qm.Where("sf = ?", sf),
	).All(ctx, exec)
	if err != nil {
		return nil, err
	}

	return w, nil
}

// ApproveRequest approves a request.
func ApproveRequest(ctx context.Context, exec boil.ContextTransactor, sf, uuid string) error {
	w, err := models.Whitelists(
		qm.Where("sf = ?", sf),
		qm.Where("uuid = ?", uuid),
	).One(ctx, exec)
	if err != nil {
		return err
	}

	w.WhitelistStatus = models.WhitelistStatusApproved
	if err := w.Update(ctx, exec, boil.Infer()); err != nil {
		return nil
	}

	return nil
}

// RejectRequest rejects a request.
func RejectRequest(ctx context.Context, exec boil.ContextTransactor, sf, uuid string) error {
	w, err := models.Whitelists(
		qm.Where("sf = ?", sf),
		qm.Where("uuid = ?", uuid),
	).One(ctx, exec)
	if err != nil {
		return err
	}

	w.WhitelistStatus = models.WhitelistStatusRejected
	if err := w.Update(ctx, exec, boil.Infer()); err != nil {
		return err
	}

	return nil
}

// CancelRequest cancels a request.
func CancelRequest(ctx context.Context, exec boil.ContextTransactor, sf, uuid string) error {
	w, err := models.Whitelists(
		qm.Where("sf = ?", sf),
		qm.Where("uuid = ?", uuid),
	).One(ctx, exec)
	if err != nil {
		return err
	}

	w.WhitelistStatus = models.WhitelistStatusCancelled
	return w.Update(ctx, exec, boil.Infer())
}

// DeleteRequests deletes all of a User's requests.
func DeleteRequests(ctx context.Context, exec boil.ContextTransactor, sf string) error {
	reqs, err := models.Whitelists(qm.Where("sf = ?", sf)).All(ctx, exec)
	if err != nil {
		return err
	}

	return reqs.DeleteAll(ctx, exec)
}
