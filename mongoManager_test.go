package main

import (
	"reflect"
	"testing"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func Test_startSession(t *testing.T) {
	tests := []struct {
		name string
		want *mongo.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := startSession(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("startSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDatabaseInformations(t *testing.T) {
	type args struct {
		c *mongo.Client
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getDatabaseInformations(tt.args.c)
		})
	}
}
