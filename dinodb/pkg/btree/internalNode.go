package btree

import (
	"dinodb/pkg/pager"
	"encoding/binary"
	"fmt"
	"io"
	"sort"
	"strconv"
)

// InternalNode represents a non-leaf node in our B+Tree that stores search keys and pointers to child nodes to aid traversal.
type InternalNode struct {
	NodeHeader // Embeds all NodeHeader fields.
}

// insert finds the appropriate place in a leaf node to insert a new tuple.
func (node *InternalNode) insert(key int64, value int64, update bool) (Split, error) {
	// Insert the entry into the appropriate child node.
	childIdx := node.search(key)
	child, childErr := node.getChildAt(childIdx)
	if childErr != nil {
		return Split{}, childErr
	}

	pager := child.getPage().GetPager()
	defer pager.PutPage(child.getPage())
	// Insert value into the child.
	result, childErr := child.insert(key, value, update)
	if childErr != nil {
		return Split{}, childErr
	}
	// Insert a new key into our node if necessary.
	if result.isSplit {
		split, insertSplitErr := node.insertSplit(result)
		return split, insertSplitErr
	}
	// This is the case when there was no split and no child err
	return Split{}, nil
}

// insertSplit inserts a split result into an internal node.
// If this insertion results in another split, the split is cascaded upwards.
func (node *InternalNode) insertSplit(split Split) (Split, error) { // AS here
	// insert split key
	insertPos := node.search(split.key)
	for i := node.numKeys; i > insertPos; i-- {
		// Move exsiting ones backward
		node.updateKeyAt(i, node.getKeyAt(i))
	}
	node.updateKeyAt(insertPos, split.key)
	pnIndex := insertPos + 1
	for i := node.numKeys + 1; i > pnIndex; i-- {
		// Move exsiting ones backward
		node.updatePNAt(i, node.getPNAt(i))
	}
	node.updatePNAt(pnIndex, split.rightPN)

	node.updateNumKeys(node.numKeys + 1)
	if node.numKeys >= KEYS_PER_INTERNAL_NODE {
		return node.split()
	}
	return Split{}, nil
}

// split is a helper function that splits an internal node, then propagates the split upwards.
func (node *InternalNode) split() (Split, error) { // AS here
	// Create a new internal node
	newNode, err := createInternalNode(node.page.GetPager())
	if err != nil {
		return Split{}, err
	}

	// Get the midpoint for splitting
	midIndex := (KEYS_PER_INTERNAL_NODE + 1) / 2

	// Move keys and page numbers to the new node
	for i := midIndex; i < KEYS_PER_INTERNAL_NODE; i++ {
		newNode.updateKeyAt(i-midIndex, node.getKeyAt(i))
		newNode.updatePNAt(i-midIndex, node.getPNAt(i+1))
	}
	newNode.updateNumKeys(KEYS_PER_INTERNAL_NODE - midIndex)
	node.updateNumKeys(midIndex) // Midpoint key gets pushed up

	// node.page.GetPager().PutPage(newNode.page)

	return Split{
		isSplit: true,
		key:     node.getKeyAt(midIndex - 1),
		leftPN:  node.page.GetPageNum(),
		rightPN: newNode.page.GetPageNum(),
	}, nil
}

// delete removes a given tuple from the leaf node, if the given key exists.
func (node *InternalNode) delete(key int64) {
	// Get the next child node where the key would be located under
	childIdx := node.search(key)
	child, err := node.getChildAt(childIdx)
	if err != nil {
		return
	}
	pager := child.getPage().GetPager()
	defer pager.PutPage(child.getPage())
	// Delete from child
	child.delete(key)
}

// get returns the value associated with a given key from the leaf node.
func (node *InternalNode) get(key int64) (value int64, found bool) {
	// Find the child.
	childIdx := node.search(key)
	child, err := node.getChildAt(childIdx)
	if err != nil {
		return 0, false
	}
	pager := child.getPage().GetPager()
	defer pager.PutPage(child.getPage())
	return child.get(key)
}

/////////////////////////////////////////////////////////////////////////////
///////////////////// Internal Node  Helper Functions ///////////////////////
/////////////////////////////////////////////////////////////////////////////

// search returns the first index where key > given key.
// If no such index exists, it returns numKeys.
func (node *InternalNode) search(key int64) int64 {
	// Binary search for the key.
	minIndex := sort.Search(
		int(node.numKeys),
		func(idx int) bool {
			return node.getKeyAt(int64(idx)) > key
		},
	)
	return int64(minIndex)
}

// keyToNodeEntry is a helper function used to create cursors that point to the entry with the given key.
// Returns the node and index within that node where the entry is found.
func (node *InternalNode) keyToNodeEntry(key int64) (*LeafNode, int64, error) {
	index := node.search(key)
	child, err := node.getChildAt(index)
	if err != nil {
		return &LeafNode{}, 0, err
	}
	pager := child.getPage().GetPager()
	defer pager.PutPage(child.getPage())
	return child.keyToNodeEntry(key)
}

// printNode pretty prints our internal node.
func (node *InternalNode) printNode(w io.Writer, firstPrefix string, prefix string) {
	// Format header data.
	var nodeType string = "Internal"
	var isRoot string
	if node.isRoot() {
		isRoot = " (root)"
	}
	numKeys := strconv.Itoa(int(node.numKeys + 1))
	// Print header data.
	io.WriteString(w, fmt.Sprintf("%v[%v] %v%v size: %v\n",
		firstPrefix, node.page.GetPageNum(), nodeType, isRoot, numKeys))
	// Print entries.
	nextFirstPrefix := prefix + " |--> "
	nextPrefix := prefix + " |    "
	for idx := int64(0); idx <= node.numKeys; idx++ {
		io.WriteString(w, fmt.Sprintf("%v\n", nextPrefix))
		child, err := node.getChildAt(idx)
		if err != nil {
			return
		}
		pager := child.getPage().GetPager()
		defer pager.PutPage(child.getPage())
		child.printNode(w, nextFirstPrefix, nextPrefix)
		if idx != node.numKeys {
			io.WriteString(w, fmt.Sprintf("\n%v[KEY] %v\n", nextPrefix, node.getKeyAt(idx)))
		}
	}
}

// pageToInternalNode returns the internal node corresponding to the given page.
func pageToInternalNode(page *pager.Page) *InternalNode {
	nodeHeader := pageToNodeHeader(page)
	return &InternalNode{nodeHeader}
}

// createInternalNode creates and returns a new internal node.
// Nodes created with this function must use `PutPage()` accordingly after use.
func createInternalNode(pager *pager.Pager) (*InternalNode, error) {
	newPage, err := pager.GetNewPage()
	if err != nil {
		return &InternalNode{}, err
	}
	initPage(newPage, INTERNAL_NODE)
	return pageToInternalNode(newPage), nil
}

// getPage returns the internal node's page.
func (node *InternalNode) getPage() *pager.Page {
	return node.page
}

// getNodeType returns internalNode.
func (node *InternalNode) getNodeType() NodeType {
	return node.nodeType
}

// copy copies the metadata and data of the passed in InternalNode to this InternalNode.
func (node *InternalNode) copy(toCopy *InternalNode) {
	copy(node.page.GetData(), toCopy.page.GetData())
	node.updateNumKeys(toCopy.numKeys)
}

// isRoot returns true if the current node is the root node.
func (node *InternalNode) isRoot() bool {
	return node.page.GetPageNum() == ROOT_PN
}

// keyPos returns the offset in the page to the internal node's ith key.
func keyPos(index int64) int64 {
	return KEYS_OFFSET + index*KEY_SIZE
}

// pnPos returns the page offset to the internal node's ith child's pagenumber
func pnPos(index int64) int64 {
	return PNS_OFFSET + index*PN_SIZE
}

// getKeyAt returns the key stored at the given index of the internal node.
func (node *InternalNode) getKeyAt(index int64) int64 {
	startPos := keyPos(index)
	key, _ := binary.Varint(node.page.GetData()[startPos : startPos+KEY_SIZE])
	return key
}

// updateKeyAt updates the key at the given index of the internal node.
func (node *InternalNode) updateKeyAt(index int64, newKey int64) {
	// Serialize the key data
	data := make([]byte, KEY_SIZE)
	binary.PutVarint(data, newKey)
	startPos := keyPos(index)
	node.page.Update(data, startPos, KEY_SIZE)
}

// getPNAt returns the pagenumber stored at the given index of the internal node.
func (node *InternalNode) getPNAt(index int64) int64 {
	startPos := pnPos(index)
	pagenum, _ := binary.Varint(node.page.GetData()[startPos : startPos+PN_SIZE])
	return pagenum
}

// updatePNAt updates the pagenumber at the given index of the internal node.
func (node *InternalNode) updatePNAt(index int64, newPagenum int64) {
	// Serialize the pagenum data
	data := make([]byte, PN_SIZE)
	binary.PutVarint(data, newPagenum)
	startPos := pnPos(index)
	node.page.Update(data, startPos, PN_SIZE)
}

// getChildAt returns the internal node's ith child.
// Child nodes retrieved via this function must call `PutPage()` accordingly after use.
func (node *InternalNode) getChildAt(index int64) (Node, error) {
	// Get the child's page
	pagenum := node.getPNAt(index)
	page, err := node.page.GetPager().GetPage(pagenum)
	if err != nil {
		return &InternalNode{}, err
	}
	return pageToNode(page), nil
}

// updateNumKeys updates the numKeys field in the node struct and the underlying page.
func (node *InternalNode) updateNumKeys(newNumKeys int64) {
	node.numKeys = newNumKeys
	// Write the new data to the page
	nKeysData := make([]byte, NUM_KEYS_SIZE)
	binary.PutVarint(nKeysData, newNumKeys)
	node.page.Update(nKeysData, NUM_KEYS_OFFSET, NUM_KEYS_SIZE)
}
