package test

import (
	"reflect"
	"testing"

	"github.com/gzfs/Go~Leet/code"
)

func TestBinaryTreeInorderTraversal(t *testing.T) {
    var node_two code.TreeNode;
    var node_three code.TreeNode;

    node_one := code.TreeNode {
        Val: 2,
        Left: nil,
        Right: &node_two,
    }

    node_two = code.TreeNode {
        Val: 3,
        Left: &node_one,
        Right: &node_three,
    }

    node_three = code.TreeNode{
        Val: 5,
        Left: &node_two,
        Right: nil,
    }

    arr := code.InOrderTraversal(&node_one);

    if reflect.DeepEqual([]int{1 ,3, 2}, arr) {
        t.Fatal("Error");
    }
 }
