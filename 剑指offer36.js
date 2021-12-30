/**
 * // Definition for a Node.
 * function Node(val,left,right) {
 *    this.val = val;
 *    this.left = left;
 *    this.right = right;
 * };
 */
/**
 * @param {Node} root
 * @return {Node}
 */
 var treeToDoublyList = function(root) {
    if (root == null) {
        return null;
    }
    let prev, head;
    var dfs = function(root) {
        if (root == null) {
            return ;
        }
        dfs(root.left);
        if (prev == null) {
            head = root;
        } else {
            prev.right = root;
        }
        root.left = prev;
        prev = root;
        dfs(root.right);
    }
    dfs(root);
    head.left = prev;
    prev.right = head;
    return head
};