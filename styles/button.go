package styles

import _ "github.com/firerainos/firerain-installer/resources"

const BackButton string = `QPushButton {
	border:none;
	background-image: url(:/resources/backButton.png);
	background-position: center top; 
	background-repeat: no-repeat;
	padding-top: 60px;
}

QPushButton:hover:!pressed {
	background-image: url(:/resources/backButtonFocus.png);
}


QPushButton:hover:pressed {
	background-image: url(:/resources/backButtonFocusPressed.png);
}

QPushButton:pressed {
	background-image: url(:/resources/backButtonPressed.png);
}`

const NextButton string = `QPushButton {
	border:none;
	background-image: url(:/resources/nextButton.png);
	background-position: center top; 
	background-repeat: no-repeat;
	padding-top: 60px;
}

QPushButton:hover:!pressed {
	background-image: url(:/resources/nextButtonFocus.png);
}


QPushButton:hover:pressed {
	background-image: url(:/resources/nextButtonFocusPressed.png);
}

QPushButton:pressed {
	background-image: url(:/resources/nextButtonPressed.png);
}`
