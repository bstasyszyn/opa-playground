/*
   Copyright Gen Digital Inc.

   This file contains software code that is the intellectual property of Gen Digital.
   Gen Digital reserves all rights in the code and you may not use it without
	 written permission from Gen Digital.
*/

package api

type Policy struct {
	ID       string   `json:"id"`
	Name     string   `json:"name,omitempty"`
	Language string   `json:"language"`
	Source   []string `json:"source"`
}

type Policies struct {
	Policies []Policy `json:"policies"`
}
