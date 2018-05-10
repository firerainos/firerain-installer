package styles

const PartitionList string = `QScrollArea {
	border: none;
	background:transparent
}`

const PartitionListItem string = `QPushButton[flat="true"] {
 	border: none;
	outline: none;
	color: #eaeaea;
}

QPushButton {
	background-color: transparent;
	border-radius: 0;
}

QPushButton::checked {
	background-color: rgba(255, 255, 255, 0.8);
	border-radius: 12px;
}`