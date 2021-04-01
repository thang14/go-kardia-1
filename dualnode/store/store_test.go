package store

import (
	"testing"

	"github.com/kardiachain/go-kardia/kai/kaidb"
	"github.com/kardiachain/go-kardia/kai/kaidb/memorydb"
)

func TestStore_GetSetCheckpoint(t *testing.T) {
	type fields struct {
		db kaidb.Database
	}
	type args struct {
		checkpoint uint64
		chainId    int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint64
	}{
		{
			name:   "",
			fields: fields{db: memorydb.New()},
			args: args{
				checkpoint: 2,
				chainId:    1,
			},
			want: 2,
		},
		{
			name:   "",
			fields: fields{db: memorydb.New()},
			args: args{
				checkpoint: 96,
				chainId:    69,
			},
			want: 96,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{
				db: tt.fields.db,
			}
			err := s.SetCheckpoint(tt.args.checkpoint, tt.args.chainId)
			if err != nil {
				t.Fatalf("cannot set checkpoint, err: %v", err)
			}
			checkpoint, err := s.GetCheckpoint(tt.args.chainId)
			if err != nil {
				t.Logf("GetCheckpoint() err %+v", err)
			}
			if checkpoint != tt.want {
				t.Errorf("GetCheckpoint() mismatch = %v, want %v", checkpoint, tt.want)
			}
		})
	}
}
