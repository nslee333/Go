package blogposts

import (
	// blogposts "github.com/quii/learn-go-with-tests/reading-files"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello world2.md": {Data: []byte("hola")},
	}

	posts, err := NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d, wanted %d posts", len(posts), len(fs))
	}

	assertPost(t, posts[0], Post{Title: "Post 1", Description: "Description 1"})
}

func assertPost(t *testing.T, got Post, want Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
