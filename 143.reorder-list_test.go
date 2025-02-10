package main

import (
	"reflect"
	"testing"
)

func Test_reorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	type test struct {
		name    string
		args    args
		want    *ListNode
		wantErr bool
	}
	tests := []test{
		{
			name: "4要素のリスト",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  4,
								Next: nil,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "5要素のリスト",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  3,
								Next: nil,
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "2要素のリスト",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		reorderList(tt.args.head)
		if !reflect.DeepEqual(tt.args.head, tt.want) {
			t.Errorf("%q. reorderList() = %v, want %v", tt.name, tt.args.head, tt.want)
		}
	}
}
