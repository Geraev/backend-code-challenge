package mapbased

import (
	"bytes"
	"io"
	"log"
	"reflect"
	"sort"
	"sync"
	"testing"
)

var (
	storg = Storage{
		RWMutex: &sync.RWMutex{},
		users: map[uint64]User{
			1: {
				ID:           1,
				OnlineStatus: true,
				Friends: map[uint64]struct{}{
					2: {},
					3: {},
					4: {},
					5: {},
				},
				Connection: log.Writer(),
			},
			2: {
				ID:           2,
				OnlineStatus: true,
				Friends: map[uint64]struct{}{
					1: {},
					3: {},
					4: {},
					5: {},
				},
				Connection: log.Writer(),
			},
			3: {
				ID:           3,
				OnlineStatus: true,
				Friends: map[uint64]struct{}{
					1: {},
					2: {},
					4: {},
					5: {},
				},
				Connection: log.Writer(),
			},
			5: {
				ID:           5,
				OnlineStatus: false,
				Friends: map[uint64]struct{}{
					1: {},
					2: {},
				},
				Connection: log.Writer(),
			},
		},
	}
)

func TestStorage_AddUser(t *testing.T) {
	type args struct {
		userId  uint64
		friends []uint64
		conn    io.Writer
	}
	tests := []struct {
		name     string
		fields   Storage
		args     args
		wantFriends []uint64
		wantConn string
	}{
		{
			name:   "Testing AddUser",
			fields: storg,
			args: args{
				userId:  6,
				friends: []uint64{1, 2, 3},
				conn:    log.Writer(),
			},
			wantFriends: []uint64{1, 2, 3},
			wantConn: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				RWMutex: tt.fields.RWMutex,
				users:   tt.fields.users,
			}
			conn := &bytes.Buffer{}
			s.AddUser(tt.args.userId, tt.args.friends, conn)

			gotItem := s.users[tt.args.userId]
			gotFriends := mapKeysToSlice(gotItem.Friends)

			sortSlice(gotFriends)

			if len(s.users) == 0 {
				t.Errorf("AddUser(). Internl storage is empty")
			}

			if !reflect.DeepEqual(gotFriends, tt.wantFriends) {
				t.Errorf("AddUser(). Parametr friends: got %v, want %v", gotFriends, tt.wantFriends)
			}

			if gotConn := conn.String(); gotConn != tt.wantConn {
				t.Errorf("AddUser(). Parametr conn = %v, want %v", gotConn, tt.wantConn)
			}
		})
	}
}

func TestStorage_Followers(t *testing.T) {
	type args struct {
		userId uint64
	}
	tests := []struct {
		name          string
		fields        Storage
		args          args
		wantFollowers []uint64
	}{
		{
			name:          "Testing Followers",
			fields:        storg,
			args:          args{userId: 3},
			wantFollowers: []uint64{1, 2, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				RWMutex: tt.fields.RWMutex,
				users:   tt.fields.users,
			}

			gotFollowers := s.Followers(tt.args.userId);
			sortSlice(gotFollowers)

			if !reflect.DeepEqual(gotFollowers, tt.wantFollowers) {
				t.Errorf("Followers() = %v, want %v", gotFollowers, tt.wantFollowers)
			}
		})
	}
}

func TestStorage_GetUserStatus(t *testing.T) {
	type args struct {
		userId uint64
	}
	tests := []struct {
		name             string
		fields           Storage
		args             args
		wantOnlineStatus bool
	}{
		{
			name:             "Testing 1. GetUserStatus",
			fields:           storg,
			args:             args{userId: 3},
			wantOnlineStatus: true,
		},
		{
			name:             "Testing 2. GetUserStatus",
			fields:           storg,
			args:             args{userId: 5},
			wantOnlineStatus: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				RWMutex: tt.fields.RWMutex,
				users:   tt.fields.users,
			}
			if gotOnlineStatus := s.GetUserStatus(tt.args.userId); gotOnlineStatus != tt.wantOnlineStatus {
				t.Errorf("GetUserStatus() = %v, want %v", gotOnlineStatus, tt.wantOnlineStatus)
			}
		})
	}
}


func TestStorage_RemoveUser(t *testing.T) {
	type args struct {
		userId uint64
	}
	tests := []struct {
		name   string
		fields Storage
		args   args
	}{
		{
			name:   "Testing RemoveUser",
			fields: storg,
			args:   args{userId: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				RWMutex: tt.fields.RWMutex,
				users:   tt.fields.users,
			}

			s.RemoveUser(tt.args.userId)

			if _, ok := s.users[tt.args.userId]; ok {
				t.Errorf("GetUserStatus(). Internal storage have record with userId = %v although it should have been deleted", tt.args.userId)
			}

		})
	}
}


func TestStorage_SetUserOffline(t *testing.T) {
	type args struct {
		userId uint64
	}
	tests := []struct {
		name     string
		fields   Storage
		args     args
		wantIsOk bool
		wantOnlineStatus bool
	}{
		{
			name:     "Testing SetUserOffline",
			fields:   storg,
			args:     args{userId: 2},
			wantIsOk: true,
			wantOnlineStatus: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				RWMutex: tt.fields.RWMutex,
				users:   tt.fields.users,
			}
			if gotIsOk := s.SetUserOffline(tt.args.userId); gotIsOk != tt.wantIsOk {
				t.Errorf("SetUserOffline() = %v, want %v", gotIsOk, tt.wantIsOk)
			}

			if gotOnlineStatus := s.users[tt.args.userId].OnlineStatus; gotOnlineStatus != tt.wantOnlineStatus {
				t.Errorf("SetUserOffline() = %v, want %v", gotOnlineStatus, tt.wantOnlineStatus)
			}
		})
	}
}

// mapKeysToSlice Convert map to slice of keys.
func mapKeysToSlice(m map[uint64]struct{}) (keys []uint64) {
	for key, _ := range m {
		keys = append(keys, key)
	}
	return
}

func sortSlice(s []uint64) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}
