# 组合模式 Composite Pattern

将对象组织成树状结构，并以统一的方式处理单个对象和对象组合。

组合模式通过将对象组合成树状结构，使得客户端可以以相同的方式处理单个对象和对象组合，而无需区分它们的差异。 

组合模式常用于树状结构，用于统一叶子节点和树节点的访问，并且可以用于应用某一操作到所有子节点。

组合模式可以提供灵活性和可扩展性，因为可以动态地添加、删除和修改对象的结构。它常用于处理树形结构的数据，例如文件系统、菜单导航等。

## 主要类型的对象：叶节点和组合节点

叶节点表示树中的最底层对象，它们没有子节点。而组合节点表示树中的分支节点，它们可以包含其他子节点，可以是叶节点，也可以是其他组合节点。 

