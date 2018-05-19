package styles

const DEList string = `QListWidget {
	border: none;
	background:transparent
}
QListWidget::item {
	border-radius: 12px;
	padding-top: 15px;
	padding-bottom: 10px;
	color:black;
}

QListWidget::item:selected {
     background-color: rgba(255, 255, 255, 0.8);
}

QListWidget::item:hover {
     background-color: rgba(255, 255, 255, 0.5);
}`
