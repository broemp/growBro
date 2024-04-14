package handler

import (
	"net/http"
	"strconv"

	db "github.com/broemp/cannaBro/db/postgres/sqlc"
	"github.com/broemp/cannaBro/view/diary"
)

func HandlePosts(w http.ResponseWriter, r *http.Request) error {
	URLLimit := r.URL.Query().Get("limit")
	URLOffset := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(URLLimit)
	if err != nil {
		return err
	}

	offset, err := strconv.Atoi(URLOffset)
	if err != nil {
		return err
	}

	args := db.ListPostsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	posts, err := db.Store.ListPosts(r.Context(), args)
	if err != nil {
		return err
	}
	return render(w, r, diary.Posts(posts))
}

func HandleNewDiaryEntrie(w http.ResponseWriter, r *http.Request) error {
	return render(w, r, diary.NewDiaryEntrie())
}
