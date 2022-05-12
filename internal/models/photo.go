// Copyright 2022 Alexander Samorodov <blitzerich@gmail.com>

package models

type Photo struct {
	ID          int64  `json:"id"`
	UrlOriginal string `json:"url_original"`
	UrlW2048    string `json:"url_w2048"`
	UrlW1024    string `json:"url_w1024"`
	UrlW512     string `json:"url_w512"`
	UrlW256     string `json:"url_w256"`
	UrlW128     string `json:"url_w128"`
}
