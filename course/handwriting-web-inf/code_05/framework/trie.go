/**
 * @Author: vincent
 * @Description:
 * @File:  trie
 * @Version: 1.0.0
 * @Date: 2021/9/21 23:07
 */

package framework

import (
	"errors"
	"strings"
)

type Tree struct {
	root *node // 根节点
}

// 节点
type node struct {
	isLast   bool                // 该节点是否能成为一个独立的uri，是否自身就是一个终极节点
	segment  string              // uri中的字符串
	handlers []ControllerHandler // 中间件+控制器
	childs   []*node             // 子节点
	parent   *node               // 父节点，双向指针
}

func newNode() *node {
	return &node{
		isLast:  false,
		segment: "",
		childs:  []*node{},
		parent:  nil,
	}
}

func NewTree() *Tree {
	root := newNode()
	return &Tree{root: root}
}

// 判断一个segment是否是通用segment,即以:开头
func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

// 过滤下一层满足segment规则的子节点
func (n *node) filterChildNodes(segment string) []*node {
	if len(n.childs) == 0 {
		return nil
	}

	// 如果segment是通配符,则所有下一层子节点都满足需求
	if isWildSegment(segment) {
		return n.childs
	}

	nodes := make([]*node, 0, len(n.childs))
	for _, child := range n.childs {
		if isWildSegment(child.segment) {
			// 如果下一层子节点有通配符，则满足需求
			nodes = append(nodes, child)
		} else if child.segment == segment {
			// 如果下一层子节点没有通配符,但是文本完全匹配，则满足需求
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 判断路由是否已经在节点的所有子节点树中存在了
func (n *node) matchNode(uri string) *node {
	// 使用分隔符将uri切割为两个部分
	segments := strings.SplitN(uri, "/", 2)
	// 取第一个部分用于匹配下一层子节点
	segment := segments[0]
	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}
	// 匹配符合的下一层子节点
	childNodes := n.filterChildNodes(segment)
	// 如果当前子节点没有一个符合,那么说明这个uri一定是之前不存在,直接返回nil
	if childNodes == nil || len(childNodes) == 0 {
		return nil
	}

	// 如果只有一个segment, 则是最后一个标记
	if len(segments) == 1 {
		// 如果segment已经是最后一个节点， 判断这些childNode是否有isLast标志
		for _, tailNode := range childNodes {
			if tailNode.isLast {
				return tailNode
			}
		}

		// 都不是最后一个节点
		return nil
	}

	// 如果有2个segment，递归每个子节点继续进行查找
	for _, tailNode := range childNodes {
		tailNodeMatch := tailNode.matchNode(segments[1])
		if tailNodeMatch != nil {
			return tailNodeMatch
		}
	}
	return nil
}

// 增加路由节点,路由节点有先后顺序
/*
/book/list
/book/:id (冲突)
/book/:id/name
/book/:student/age
/:user/name(冲突)
/:user/name/:age
*/
func (tree *Tree) AddRouter(uri string, handlers []ControllerHandler) error {
	n := tree.root
	if n.matchNode(uri) != nil {
		return errors.New("route exist: " + uri)
	}

	segments := strings.Split(uri, "/")

	// 对每个segment
	for index, segment := range segments {
		// 最终进入Node segment字段
		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}
		isLast := index == len(segments)-1

		var objNode *node // 标记是否有合适的子节点
		childNodes := n.filterChildNodes(segment)
		// 如果有匹配的子节点
		if len(childNodes) > 0 {
			// 如果有segment相同的节点，则选择这个子节点
			for _, cnode := range childNodes {
				if cnode.segment == segment {
					objNode = cnode
					break
				}
			}
		}

		if objNode == nil {
			// 创建一个当前node的节点
			cnode := newNode()
			cnode.segment = segment
			if isLast {
				cnode.isLast = true
				cnode.handlers = handlers
			}
			// 父节点指针修改
			cnode.parent = n
			n.childs = append(n.childs, cnode)
			objNode = cnode
		}

		n = objNode
	}

	return nil
}

// 匹配uri
func (tree *Tree) FindHandler(uri string) []ControllerHandler {
	matchNode := tree.root.matchNode(uri)
	if matchNode == nil {
		return nil
	}
	return matchNode.handlers
}

// 将uri解析为params
func (n *node) parseParamsFromEndNode(uri string) map[string]string {
	ret := map[string]string{}
	segments := strings.Split(uri, "/")
	cnt := len(segments)
	cur := n
	for i := cnt - 1; i >= 0; i-- {
		if cur.segment == "" {
			break
		}

		// 如果是通配符节点
		if isWildSegment(cur.segment) {
			// 设置params,从1赋值，是因为要把":id"里的":"跳过
			ret[cur.segment[1:]] = segments[i]
		}
		cur = cur.parent
	}
	return ret
}