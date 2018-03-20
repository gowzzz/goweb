package main

import (
	"testing"
	"time"
)

func TestDecode(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)
	post, err := decode("post.json")
	if err != nil {
		t.Error(err)
	}

	if post.Id != 1 {
		t.Error("wrong id,was expecting 1 but got :", post.Id)
	}

	if post.Content != "Helloworld" {
		t.Error("wrong id,was expecting 'Helloworld' but got :", post.Content)
	}
}

func TestEncode(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
	t.Log("TestEncode  is running")
	// t.Skip("skiping encoding for now")
}

func TestLongRunningFunc(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
	if testing.Short() {
		t.Skip("skip long time func:TestLongRunningFunc")
	}
	t.Log("TestLongRunningFunc  is running")
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decode("post.json")
	}
}
func BenchmarkEncoder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encoder("post.json")
	}
}
