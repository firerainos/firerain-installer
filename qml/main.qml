import QtQuick 2.7
import QtQuick.Controls 2.3
import QtQuick.Layouts 1.3
import "page"

ApplicationWindow {
  id: root
  visible: true
  title: "Hello World Example"
  minimumWidth: 900
  minimumHeight: 700
  maximumWidth: 900
  maximumHeight: 700
  // flags : Qt.window | Qt.FramelessWindowHint

  Button {
    id: backButton
    text : "back"
    anchors.top:parent.top
    anchors.left : parent.left

    onClicked:{
      stackLayout.currentIndex --
    }
    visible: stackLayout.currentIndex > 0
  }

  StackLayout {
    id : stackLayout
    // initialItem: view
    WelcomePage{}
    NetworkPage{}
    PartitionPage{}
    SelectDMPage{}
    AdditionalSoftwarePage{}
    InstallPage{}
  }


  Button {
    id:nextButton
    text : "next"
    anchors.bottom:parent.bottom
    anchors.horizontalCenter: parent.horizontalCenter;  
    onClicked:{
      stackLayout.currentIndex ++
    }
    visible: stackLayout.currentIndex < stackLayout.count-1
  }
}
