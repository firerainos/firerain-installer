package styles

const ItemList = `QListWidget {
	border: none;
	background: transparent;
}
QListWidget::item {
	height: 52px;
}

QListWidget::item:selected {
	color: white;
	background-color: rgb(41, 189, 139);
}

QListWidget::item:hover {
	color: white;
	background-color: rgba(41, 189, 139,0.5);
}`

const PackageList = `QListView {
	border: none;
	alternate-background-color: rgb(245, 245, 245);
}`

const InstallList = `QListView {
	border: none;
	alternate-background-color: rgb(245, 245, 245);
}`