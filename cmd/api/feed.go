package main

import (
	"net/http"

	"github.com/karthikbhandary2/Social/internal/store"
)


// GetUserFeed godoc
//
//	@Summary		Fetches the user feed
//	@Description	Fetches the user feed with optional filters and pagination
//	@Tags			feed
//	@Accept			json
//	@Produce		json
//	@Param			since	query		string						false	"Fetch posts since this timestamp (ISO 8601 format)"
//	@Param			until	query		string						false	"Fetch posts until this timestamp (ISO 8601 format)"
//	@Param			limit	query		int							false	"Maximum number of posts to fetch (default: 20)"
//	@Param			offset	query		int							false	"Number of posts to skip before starting to fetch"
//	@Param			sort	query		string						false	"Sort order (e.g., 'asc' or 'desc')"
//	@Param			tags	query		string						false	"Comma-separated list of tags to filter posts"
//	@Param			search	query		string						false	"Search term to filter posts by content"
//	@Success		200		{object}	[]store.PostWithMetadata	"List of posts with metadata"
//	@Failure		400		{object}	error						"Bad request"
//	@Failure		500		{object}	error						"Internal server error"
//	@Security		ApiKeyAuth
//	@Router			/users/feed [get]
func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	fq := store.PaginatedFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
		Tags:   []string{},
		Search: "",
	}

	fq, err := fq.Parse(r)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	if err := Validate.Struct(fq); err != nil {
		app.badRequest(w, r, err)
		return
	}

	ctx := r.Context()
	user := getUserFromContext(r)

	feed, err := app.store.Posts.GetUserFeed(ctx, user.ID, fq)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, feed); err != nil {
		app.internalServerError(w, r, err)
	}
}
