import QtQuick 2.7			
import QtQuick.Controls 1.4
import QtQuick.Controls.Styles 1.4
import QtGraphicalEffects 1.0

ApplicationWindow {
	id: window
	visible: true
	title: "Vertcoin - OneClick Miner"
	minimumWidth: 500
	minimumHeight: 400
    Rectangle {
        width: parent.width
        height: parent.height 
        color: "#202225"
        opacity: { 
            if(viewModel.activePage == "init") {
                return 1;
            } else {
                return 0;
            }
        }

        Column {
            anchors.centerIn: parent
            spacing: 15
            Text {
                id: initTitle
                text: "Vertcoin - OneClick Miner"
                font.pointSize: 16
                font.weight: Font.DemiBold
                color: "#555555"
                anchors.horizontalCenter: parent.horizontalCenter
            }

            Text {
                id: initStatus
                text: viewModel.initStatus
                font.pointSize: 14
                font.weight: Font.Light
                color: "#555555"
                anchors.horizontalCenter: parent.horizontalCenter
            }
        }
    }

    Rectangle {
        width: parent.width
        height: parent.height 
        color: "#202225"

        opacity: { 
            if(viewModel.activePage == "start") {
                return 1;
            } else {
                return 0;
            }
        }

        Column {
            anchors.centerIn: parent
            spacing: 25
            Text {
                bottomPadding: 10
                id: startTitle
                text: "Vertcoin - OneClick Miner"
                font.pointSize: 16
                font.weight: Font.DemiBold
                color: "#555555"
                anchors.horizontalCenter: parent.horizontalCenter
            }

            TextField {
                height: 56
                width: 400
                placeholderText: "Enter Your Wallet Address"
                anchors.horizontalCenter: parent.horizontalCenter
                style: TextFieldStyle {
                    textColor: "#555555"
                    placeholderTextColor: "#555555"
                    font.pointSize: 14
                    font.weight: Font.Light
                    background: Rectangle {
                        radius: 5
                        color: "#202225"
                        border.color: "#048657"
                        border.width: 1
                    }
                }
            }
            Item {
                anchors.horizontalCenter: parent.horizontalCenter
                height: 76
                width: 500
            
                Button {
                    anchors.horizontalCenter: parent.horizontalCenter
                    height: 56
                    width: 400
                    id: startButton
                    style: ButtonStyle {
                        background: Rectangle {
                            radius: 5
                            implicitWidth: 100
                            implicitHeight: 76
                            color: "#048657"
                        }

                        label: Text {
                            font.pointSize: 14
                            font.weight: Font.DemiBold
                            color: "#cdcdcd"
                            text: "Start mining"
                            wrapMode: Text.WordWrap
                            verticalAlignment: Text.AlignVCenter
                            horizontalAlignment: Text.AlignHCenter
                            anchors.fill: parent
                        }
                    }
                }

                Glow {
                    anchors.fill: startButton
                    radius: 12
                    samples: 17
                    color: "#213933"
                    source: startButton
                }
            }
        }
    }
}