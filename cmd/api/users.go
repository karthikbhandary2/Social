package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/karthikbhandary2/Social/internal/store"
)

type userKey string

const userCtx postKey = "user"

// ShowAccount godoc
//	@Summary		Fetches a user profile
//	@Description	Fetches a user profile by ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	store.User
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Security		ApiKeyAuth
//	@Router			/accounts/{id} [get]

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {

	user := app.getUserFromContext(r)

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) usersContextMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        idParam := chi.URLParam(r, "userID")
        app.logger.Info("userID param: %s", idParam) // Log the userID parameter

        userID, err := strconv.ParseInt(idParam, 10, 64)
        if err != nil {
            app.badRequest(w, r, err)
            return
        }
        ctx := r.Context()
        user, err := app.store.Users.GetByID(ctx, userID)
        if err != nil {
            switch {
            case errors.Is(err, store.ErrNotFound):
                app.badRequest(w, r, err)
                return
            default:
                app.internalServerError(w, r, err)
                return
            }
        }
        if user == nil {
            app.notFound(w, r, errors.New("user not found"))
            return
        }

        // Log the user details for debugging
        app.logger.Info("User found: %+v", user)

        ctx = context.WithValue(ctx, userCtx, user)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func (app *application) getUserFromContext(r *http.Request) *store.User {
	user, ok := r.Context().Value(userCtx).(*store.User)
	if !ok {
		app.logger.Info("Post not found in context")
	} else {
		app.logger.Info("Post retrieved from context: %+v", user)
	}
	return user
}

type FollowUser struct {
	UserID int64 `json:"user_id"`
}

func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	followerUser := app.getUserFromContext(r)

	var payload FollowUser
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequest(w, r, err)
		return
	}

	ctx := r.Context()
	if err := app.store.Followers.Follow(ctx, followerUser.ID, payload.UserID); err != nil {
		switch err {
		case store.ErrConflict:
			app.conflict(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}

	}
	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	unFolloweduser := app.getUserFromContext(r)

	var payload FollowUser
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequest(w, r, err)
		return
	}
	ctx := r.Context()
	if err := app.store.Followers.Unfollow(ctx, unFolloweduser.ID, payload.UserID); err != nil {
		app.internalServerError(w, r, err)
		return
	}
	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
	}
}
